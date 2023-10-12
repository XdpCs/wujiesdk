package wujiesdk

// @Title        client.go
// @Description  request wujie's api
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-10 20:47

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
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

type HttpHooks []HttpHook

type Client struct {
	httpClient    *http.Client // http httpClient
	MaxRetryTimes int          // max retry times
	HttpHooks     HttpHooks    // hook before and after request
	Credentials   *Credentials
	Logger        *log.Logger
	IsDebug       bool
}

// NewDefaultClient all api need auth
func NewDefaultClient(c *Credentials) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 200 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives: true,
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			},
		},
		MaxRetryTimes: 3,
		Credentials:   c,
		Logger:        log.New(os.Stdout, "", log.LstdFlags),
	}
	client.AddHttpHooks(c)
	return client
}

func (c *Client) AddHttpHooks(hooks ...HttpHook) {
	c.HttpHooks = append(c.HttpHooks, hooks...)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.do(req)
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	for _, hook := range c.HttpHooks {
		if err := hook.BeforeRequest(req); err != nil {
			return nil, fmt.Errorf("hook.BeforeRequest: %w", err)
		}
	}
	var (
		resp *http.Response
		err  error
	)

	for i := 0; i < c.MaxRetryTimes; i++ {
		resp, err = c.httpClient.Do(req)
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
	if err != nil {
		c.LoggerHTTPReq(req)
		c.LoggerHTTPResp(req, resp)
	}

	for _, hook := range c.HttpHooks {
		hook.AfterRequest(resp, err)
	}

	return resp, err
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

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
	c.WriteLog(LogError, "%s\n", logBuffer.String())
}

// LoggerHTTPResp Print Response to http request
func (c *Client) LoggerHTTPResp(req *http.Request, resp *http.Response) {
	var logBuffer bytes.Buffer
	logBuffer.WriteString(fmt.Sprintf("[Resp:%p]StatusCode:%d\t", req, resp.StatusCode))
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
	c.WriteLog(LogError, "%s\n", logBuffer.String())
}

func (c *Client) AvailableIntegralBalance(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(AvailableIntegralBalanceWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(AvailableIntegralBalanceWujieRouter), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationFormUrlencodedUTF8)
	return c.do(req)
}

func (c *Client) ExchangePoint(ctx context.Context, eReq *ExchangePointRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ExchangePointWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ExchangePointWujieRouter), err)
	}
	data, err := json.Marshal(eReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", eReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) ModelBaseInfos(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ModelBaseInfosWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ModelBaseInfosWujieRouter), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationFormUrlencodedUTF8)
	return c.do(req)
}

func (c *Client) DefaultResourceStyleModel(ctx context.Context) (*http.Response, error) {
	path, err := url.Parse(Domain + string(DefaultResourceStyleModelWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(DefaultResourceStyleModelWujieRouter), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationFormUrlencodedUTF8)
	return c.do(req)
}

func (c *Client) DefaultResourceModel(ctx context.Context, model int32) (*http.Response, error) {
	values := url.Values{
		"model": []string{fmt.Sprintf("%d", model)},
	}
	rawURL := Domain + string(DefaultResourceModelWujieRouter) + "?" + values.Encode()
	path, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", rawURL, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationFormUrlencodedUTF8)
	return c.do(req)
}

func (c *Client) CreateImage(ctx context.Context, cReq *CreateImageRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateImageWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateImageWujieRouter), err)
	}
	data, err := json.Marshal(cReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", cReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) GeneratingInfo(ctx context.Context, gReq *GeneratingInfoRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(GeneratingInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(GeneratingInfoWujieRouter), err)
	}
	data, err := json.Marshal(gReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", gReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) ImageInfo(ctx context.Context, key string) (*http.Response, error) {
	values := url.Values{
		"key": []string{key},
	}
	rawURL := Domain + string(ImageInfoWujieRouter) + "?" + values.Encode()
	path, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", rawURL, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) ImagePriceInfo(ctx context.Context, iReq *ImagePriceInfoRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(ImagePriceInfoWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(ImagePriceInfoWujieRouter), err)
	}
	data, err := json.Marshal(iReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", iReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) PostSuperSize(ctx context.Context, sReq *PostSuperSizeRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(SuperSizeWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(SuperSizeWujieRouter), err)
	}
	data, err := json.Marshal(sReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", sReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) GetSuperSize(ctx context.Context, keys []string) (*http.Response, error) {
	values := url.Values{
		"key": keys,
	}
	rawURL := Domain + string(SuperSizeWujieRouter) + "?" + values.Encode()
	path, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", rawURL, err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) CancelImage(ctx context.Context, cReq *CancelImageRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CancelImageWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CancelImageWujieRouter), err)
	}
	data, err := json.Marshal(cReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", cReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) CreateImagePro(ctx context.Context, cReq *CreateImageProRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(CreateImageProWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(CreateImageProWujieRouter), err)
	}
	data, err := json.Marshal(cReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", cReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}

func (c *Client) GeneratingInfoPro(ctx context.Context, gReq *GeneratingInfoProRequest) (*http.Response, error) {
	path, err := url.Parse(Domain + string(GeneratingInfoProWujieRouter))
	if err != nil {
		return nil, fmt.Errorf("url.Parse: url: %v, parse url error: %w", Domain+string(GeneratingInfoProWujieRouter), err)
	}
	data, err := json.Marshal(gReq)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: req: %v, marshal req error: %w", gReq.String(), err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: url: %v, new request error: %w", path.String(), err)
	}
	req.Header.Set(ContentType, ApplicationJson)
	return c.do(req)
}
