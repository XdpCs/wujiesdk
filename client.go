package wujiesdk

// @Title        client.go
// @Description  request wujie's api
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-21 19:51

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// HttpHook uses BeforeRequest and AfterRequest
type HttpHook interface {
	BeforeRequest(req *http.Request) error
	AfterRequest(response *http.Response, err error)
}

// HttpHooks is a slice of HttpHook
type HttpHooks []HttpHook

// Client is the client for wujie's api
type Client struct {
	httpClient    *http.Client // http httpClient
	MaxRetryTimes int          // max retry times
	HttpHooks     HttpHooks    // hook before and after request
	Credentials   *Credentials
	Logger        *Logger
}

// Logger is the logger for wujie's api
type Logger struct {
	*log.Logger
	LogLevel int
}

// NewLogger new logger
func NewLogger(logLevel int, log *log.Logger) *Logger {
	return newLogger(logLevel, log)
}

// NewDefaultLogger default logger
func NewDefaultLogger() *Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return newLogger(LogInfo, logger)
}

// NewDebugLogger debug logger
func NewDebugLogger() *Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return newLogger(LogDebug, logger)
}

func newLogger(logLevel int, log *log.Logger) *Logger {
	return &Logger{
		Logger:   log,
		LogLevel: logLevel,
	}
}

// NewDefaultClient all api need auth
func NewDefaultClient(c *Credentials) *Client {
	return newClient(newDefaultHttpClient(), 3, c, NewDefaultLogger())
}

// NewDebugClient log http request and response
func NewDebugClient(c *Credentials) *Client {
	client := newClient(newDefaultHttpClient(), 3, c, NewDebugLogger())
	return client
}

// NewClient new client
func NewClient(httpClient *http.Client, maxRetryTimes int, c *Credentials, logger *Logger) *Client {
	return newClient(httpClient, maxRetryTimes, c, logger)
}

func newDefaultHttpClient() *http.Client {
	return &http.Client{
		Timeout: 200 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func newClient(httpClient *http.Client, maxRetryTimes int, c *Credentials, logger *Logger) *Client {
	client := &Client{
		httpClient:    httpClient,
		MaxRetryTimes: maxRetryTimes,
		Credentials:   c,
		Logger:        logger,
	}
	client.AddHttpHooks(c)
	return client
}

// AddHttpHooks add hooks
func (c *Client) AddHttpHooks(hooks ...HttpHook) {
	c.HttpHooks = append(c.HttpHooks, hooks...)
}

func (c *Client) do(req *http.Request, rawBody []byte) (*http.Response, error) {
	for _, hook := range c.HttpHooks {
		if err := hook.BeforeRequest(req); err != nil {
			return nil, fmt.Errorf("hook.BeforeRequest: %w", err)
		}
	}
	// make sure request one time at least
	if c.MaxRetryTimes <= 0 {
		c.MaxRetryTimes = 1
	}
	var (
		resp        *http.Response
		err         error
		requestBody *bytes.Reader
	)
	requestBody = bytes.NewReader(rawBody)
	for i := 0; i < c.MaxRetryTimes; i++ {
		_, err = requestBody.Seek(0, io.SeekStart)
		if err != nil {
			return nil, fmt.Errorf("requestBody.Seek: error: %w", err)
		}
		req.Body = io.NopCloser(requestBody)
		resp, err = c.httpClient.Do(req)
		if c.Logger.LogLevel == LogDebug {
			c.LoggerHTTPReq(req)
			// if this logger does not log response, response will be nil
			c.LoggerHTTPResp(resp)
		}
		if err != nil {
			err = fmt.Errorf("c.httpClient.Do error: %v", err)
			continue
		}
		if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
			err = fmt.Errorf("http status code: %d, %s, trace_id: %v", resp.StatusCode,
				http.StatusText(resp.StatusCode), getTraceID(resp))
			continue
		}
		break
	}

	for _, hook := range c.HttpHooks {
		hook.AfterRequest(resp, err)
	}

	return resp, err
}

func (c *Client) ctxJson(ctx context.Context, httpMethod string, api string, params url.Values, body interface{}) (*http.Response, error) {
	apiParams := params.Encode()
	var rawBody []byte
	if apiParams != "" {
		api = api + "?" + apiParams
	}
	req, err := http.NewRequestWithContext(ctx, httpMethod, api, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", api, err)
	}
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal: marshal body error: %w", err)
		}
		rawBody = data
		req.Body = io.NopCloser(bytes.NewReader(data))
	}

	// body must be set some data in Post
	if httpMethod == http.MethodPost && body == nil {
		rawBody = []byte("{}")
		req.Body = io.NopCloser(bytes.NewReader(rawBody))
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req, rawBody)
}

// CtxPostJson http post json
func (c *Client) CtxPostJson(ctx context.Context, api string, params url.Values, body interface{}) (*http.Response, error) {
	return c.ctxJson(ctx, http.MethodPost, api, params, body)
}

// CtxGetJson http get json
func (c *Client) CtxGetJson(ctx context.Context, api string, params url.Values) (*http.Response, error) {
	return c.ctxJson(ctx, http.MethodGet, api, params, nil)
}

// HTTPClient return http client
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// SetHTTPClient set http client
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// WriteLog output log function
func (c *Client) WriteLog(LogLevel int, format string, a ...interface{}) {
	if c.Logger == nil {
		return
	}

	var logBuffer bytes.Buffer
	logBuffer.WriteString(LogTag[LogLevel-1])
	logBuffer.WriteString(fmt.Sprintf(format, a...))
	c.Logger.Printf("%s", logBuffer.String())
}

// LoggerHTTPReq Print the header information of the http request
func (c *Client) LoggerHTTPReq(req *http.Request) {
	var logBuffer bytes.Buffer
	logBuffer.WriteString(fmt.Sprintf("[Req:%p]Method:%s\t", req, req.Method))
	logBuffer.WriteString(fmt.Sprintf("Host:%s\t", req.URL.Host))
	logBuffer.WriteString(fmt.Sprintf("Path:%s\t", req.URL.Path))
	logBuffer.WriteString(fmt.Sprintf("Query:%s\t", req.URL.RawQuery))
	logBuffer.WriteString("Header info:")

	for k, v := range req.Header {
		var valueBuffer bytes.Buffer
		for j := 0; j < len(v); j++ {
			if j > 0 {
				valueBuffer.WriteString(" ")
			}
			valueBuffer.WriteString(v[j])
		}
		logBuffer.WriteString(fmt.Sprintf("\t%s:%s", k, valueBuffer.String()))
	}
	c.WriteLog(LogDebug, "%s\n", logBuffer.String())
}

// LoggerHTTPResp Print Response to http request
func (c *Client) LoggerHTTPResp(resp *http.Response) {
	if resp == nil {
		return
	}
	var logBuffer bytes.Buffer
	logBuffer.WriteString(fmt.Sprintf("[Resp:%p]StatusCode:%d\t", resp, resp.StatusCode))
	logBuffer.WriteString("Header info:")
	for k, v := range resp.Header {
		var valueBuffer bytes.Buffer
		for j := 0; j < len(v); j++ {
			if j > 0 {
				valueBuffer.WriteString(" ")
			}
			valueBuffer.WriteString(v[j])
		}
		logBuffer.WriteString(fmt.Sprintf("\t%s:%s", k, valueBuffer.String()))
	}
	c.WriteLog(LogDebug, "%s\n", logBuffer.String())
}

// AvailableIntegralBalance get available integral balance
func (c *Client) AvailableIntegralBalance(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(AvailableIntegralBalanceWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(AvailableIntegralBalanceWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// ExchangePoint exchange points with people
func (c *Client) ExchangePoint(ctx context.Context, eReq *ExchangePointRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ExchangePointWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ExchangePointWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, eReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", eReq.String(), err)
	}
	return resp, nil
}

// ModelBaseInfos get model base infos
func (c *Client) ModelBaseInfos(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ModelBaseInfosWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ModelBaseInfosWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// DefaultResourceStyleModel get default resource style model
func (c *Client) DefaultResourceStyleModel(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(DefaultResourceStyleModelWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(DefaultResourceStyleModelWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// DefaultResourceModel get model's default resource
func (c *Client) DefaultResourceModel(ctx context.Context, model int32) (*http.Response, error) {
	values := url.Values{
		"model": []string{fmt.Sprintf("%d", model)},
	}
	path, err := url.Parse(Domain + string(DefaultResourceModelWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(DefaultResourceModelWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", model, err)
	}
	return resp, nil
}

// CreateImage create image
func (c *Client) CreateImage(ctx context.Context, cReq *CreateImageRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateImageWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateImageWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// GeneratingInfo get image generating info
func (c *Client) GeneratingInfo(ctx context.Context, keys []string) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ImageGeneratingInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImageGeneratingInfoWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, keys)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: keys: %v, error: %w", keys, err)
	}
	return resp, nil
}

// ImageInfo get image detail
func (c *Client) ImageInfo(ctx context.Context, key string) (*http.Response, error) {
	values := url.Values{
		"key": []string{key},
	}
	path, err := url.Parse(Domain + string(ImageInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImageInfoWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", key, err)
	}
	return resp, nil
}

// ImagePriceInfo get image price info
func (c *Client) ImagePriceInfo(ctx context.Context, iReq *ImagePriceInfoRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ImagePriceInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImagePriceInfoWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, iReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", iReq.String(), err)
	}
	return resp, nil
}

// PostSuperSize create super size
func (c *Client) PostSuperSize(ctx context.Context, sReq *PostSuperSizeRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(SuperSizeWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(SuperSizeWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, sReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", sReq.String(), err)
	}
	return resp, nil
}

// GetSuperSize get super size result
func (c *Client) GetSuperSize(ctx context.Context, keys []string) (*http.Response, error) {
	values := url.Values{
		"key": keys,
	}
	path, err := url.Parse(Domain + string(SuperSizeWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(SuperSizeWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", keys, err)
	}
	return resp, nil
}

// CreateParams get create params
func (c *Client) CreateParams(ctx context.Context, keys []string) (*http.Response, error) {
	body := &struct {
		Key []string `json:"key"`
	}{
		Key: keys,
	}
	path, err := url.Parse(Domain + string(CreateParamsWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateParamsWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, body)
	if err != nil {
		return nil, fmt.Errorf("c.CtxJson: keys: %v, error: %w", keys, err)
	}
	return resp, nil
}

// ImageModelQueueInfo get image model queue info
func (c *Client) ImageModelQueueInfo(ctx context.Context, model int32) (*http.Response, error) {
	values := url.Values{
		"model": []string{fmt.Sprintf("%d", model)},
	}
	path, err := url.Parse(Domain + string(ImageModelQueueInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImageModelQueueInfoWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", model, err)
	}
	return resp, nil
}

// CancelImage cancel image
func (c *Client) CancelImage(ctx context.Context, key string) (*http.Response, error) {
	body := &struct {
		Key string `json:"key"`
	}{
		Key: key,
	}
	path, err := url.Parse(Domain + string(CancelImageWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CancelImageWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, body)
	if err != nil {
		return nil, fmt.Errorf("c.CtxJson: key: %v, error: %w", key, err)
	}
	return resp, nil
}

// AccelerateImage accelerate image
func (c *Client) AccelerateImage(ctx context.Context, aReq *AccelerateImageRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(AccelerateImageWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(AccelerateImageWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, aReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxJson: req: %v, error: %w", aReq.String(), err)
	}
	return resp, nil
}

// PromptOptimizeSubmit submit prompt optimize
func (c *Client) PromptOptimizeSubmit(ctx context.Context, pReq *PromptOptimizeSubmitRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(PromptOptimizeSubmitWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(PromptOptimizeSubmitWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, pReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", pReq.String(), err)
	}
	return resp, nil
}

// PromptOptimizeResult get prompt optimize result
func (c *Client) PromptOptimizeResult(ctx context.Context, taskID string) (*http.Response, error) {
	values := url.Values{
		"taskId": []string{taskID},
	}
	path, err := url.Parse(Domain + string(PromptOptimizeResultWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(PromptOptimizeResultWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", taskID, err)
	}
	return resp, nil
}

// Youthify youthify image
func (c *Client) Youthify(ctx context.Context, yReq *YouthifyRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(YouthifyWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(YouthifyWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, yReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", yReq.String(), err)
	}
	return resp, nil
}

// QuerySpell query spell
func (c *Client) QuerySpell(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(QuerySpellWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(QuerySpellWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// CreateImagePro create pro image
func (c *Client) CreateImagePro(ctx context.Context, cReq *CreateImageProRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateImageProWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateImageProWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// GeneratingInfoPro get pro image generating info
func (c *Client) GeneratingInfoPro(ctx context.Context, keys []string) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ImageGeneratingInfoProWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImageGeneratingInfoProWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, keys)
	if err != nil {
		return nil, fmt.Errorf("c.CtxJson: req: %v, error: %w", keys, err)
	}
	return resp, nil
}

// CreateAvatar create avatar
func (c *Client) CreateAvatar(ctx context.Context, cReq *CreateAvatarRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateAvatarWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateAvatarWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// DeleteAvatar delete avatar
func (c *Client) DeleteAvatar(ctx context.Context, key string) (*http.Response, error) {
	body := &struct {
		AvatarKey string `json:"avatar_key"`
	}{
		AvatarKey: key,
	}
	path, err := url.Parse(Domain + string(DeleteAvatarWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(DeleteAvatarWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, body)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: key: %v, error: %w", key, err)
	}
	return resp, nil
}

// AvatarInfo get avatar info
func (c *Client) AvatarInfo(ctx context.Context, key string) (*http.Response, error) {
	values := url.Values{
		"key": []string{key},
	}
	path, err := url.Parse(Domain + string(AvatarInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(AvatarInfoWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: key: %v, error: %w", key, err)
	}
	return resp, nil
}

// ImageBatchCheck image batch check
func (c *Client) ImageBatchCheck(ctx context.Context, imageURLList []string) (*http.Response, error) {
	body := &struct {
		ImageURLList []string `json:"image_url_list"`
	}{
		ImageURLList: imageURLList,
	}
	path, err := url.Parse(Domain + string(ImageBatchCheckWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImageBatchCheckWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, body)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: imageURLList: %v, error: %w", imageURLList, err)
	}
	return resp, nil
}

// CreateAvatarArtwork create avatar artwork
func (c *Client) CreateAvatarArtwork(ctx context.Context, cReq *CreateAvatarArtworkRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateAvatarArtworkWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateAvatarArtworkWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// AvatarDefaultResource get avatar default resource
func (c *Client) AvatarDefaultResource(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(AvatarDefaultResourceWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(AvatarDefaultResourceWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// CreateSpellAnalysis create spell analysis
func (c *Client) CreateSpellAnalysis(ctx context.Context, cReq *CreateSpellAnalysisRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateSpellAnalysisWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateSpellAnalysisWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// SpellAnalysisInfo get spell analysis info
func (c *Client) SpellAnalysisInfo(ctx context.Context, key string) (*http.Response, error) {
	values := url.Values{
		"spellAnalysisKey": []string{key},
	}
	path, err := url.Parse(Domain + string(SpellAnalysisInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(SpellAnalysisInfoWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), values)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: req: %v, error: %w", key, err)
	}
	return resp, nil
}

// MagicDiceTheme get magic dice theme
func (c *Client) MagicDiceTheme(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(MagicDiceThemeWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(MagicDiceThemeWujieRouter), err)
	}
	resp, err := c.CtxGetJson(ctx, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("c.CtxGetJson: error: %w", err)
	}
	return resp, nil
}

// CreateMagicDice create magic dice
func (c *Client) CreateMagicDice(ctx context.Context, cReq *CreateMagicDiceRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateMagicDiceWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateMagicDiceWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}

// CreateVideo create video
func (c *Client) CreateVideo(ctx context.Context, cReq *CreateVideoRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateVideoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateVideoWujieRouter), err)
	}
	resp, err := c.CtxPostJson(ctx, path.String(), nil, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.CtxPostJson: req: %v, error: %w", cReq.String(), err)
	}
	return resp, nil
}
