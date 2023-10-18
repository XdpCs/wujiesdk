package wujiesdk

// @Title        caller.go
// @Description  handle wujie sdk's response
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-17 10:22

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Caller is the caller for wujie sdk
type Caller struct {
	Client *Client
}

// NewCaller create a new caller
func NewCaller(c *Client) *Caller {
	return &Caller{Client: c}
}

// AvailableIntegralBalance get available integral balance
func (c *Caller) AvailableIntegralBalance(ctx context.Context) (WujieCode, int, error) {
	resp, err := c.Client.AvailableIntegralBalance(ctx)
	if err != nil {
		return ErrorWujieCode, 0, fmt.Errorf("c.Client.AvailableIntegralBalance: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var bResp AvailableIntegralBalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&bResp); err != nil {
		return ErrorWujieCode, 0, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(bResp.Code)
	if err := code.Err(); err != nil {
		return code, 0, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, bResp.Message)
	}

	return code, bResp.Data.Balance, nil
}

// ExchangePoint exchange points with people
func (c *Caller) ExchangePoint(ctx context.Context, eReq *ExchangePointRequest) (WujieCode, bool, error) {
	resp, err := c.Client.ExchangePoint(ctx, eReq)
	if err != nil {
		return ErrorWujieCode, false, fmt.Errorf("c.Client.ExchangePoint: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var eResp BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&eResp); err != nil {
		return ErrorWujieCode, false, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(eResp.Code)
	if err := code.Err(); err != nil {
		return code, false, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, ExchangePointRequest: %s",
			getTraceID(resp), err, eResp.Message, eReq.String())
	}
	return code, true, nil
}

// ModelBaseInfos get model base infos
func (c *Caller) ModelBaseInfos(ctx context.Context) (WujieCode, []ModelBaseInfoData, error) {
	resp, err := c.Client.ModelBaseInfos(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ModelBaseInfos: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp ModelBaseInfosResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, mResp.Message)
	}

	return code, mResp.Data, nil
}

// DefaultResourceStyleModel get default resource style model
func (c *Caller) DefaultResourceStyleModel(ctx context.Context) (WujieCode, []StyleModel, error) {
	resp, err := c.Client.DefaultResourceStyleModel(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.DefaultResourceStyleModel: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var dResp DefaultResourceStyleModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&dResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(dResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, dResp.Message)
	}
	return code, dResp.Data.StyleModels, nil
}

// DefaultResourceModel get model's default resource
func (c *Caller) DefaultResourceModel(ctx context.Context, model int32) (WujieCode, *DefaultResourceModelData, error) {
	resp, err := c.Client.DefaultResourceModel(ctx, model)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.DefaultResourceModel: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp DefaultResourceModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, mResp.Message)
	}
	return code, &mResp.Data, nil
}

// CreateImage create image
func (c *Caller) CreateImage(ctx context.Context, cReq *CreateImageRequest) (WujieCode, *CreateImageData, error) {
	resp, err := c.Client.CreateImage(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateImage: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateImageRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data, nil
}

// GeneratingInfo get image generating info
func (c *Caller) GeneratingInfo(ctx context.Context, gReq *GeneratingInfoRequest) (WujieCode, []ImageGeneratingInfo, error) {
	resp, err := c.Client.GeneratingInfo(ctx, gReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.GeneratingInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GeneratingInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, GeneratingInfoRequest: %s",
			getTraceID(resp), err, gResp.Message, gReq.String())
	}
	return code, gResp.Data.List, nil
}

// ImageInfo get image detail
func (c *Caller) ImageInfo(ctx context.Context, key string) (WujieCode, *ImageInfoData, error) {
	resp, err := c.Client.ImageInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ImageInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImageInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, iResp.Message, key)
	}
	return code, &iResp.Data, nil
}

// ImagePriceInfo get image price info
func (c *Caller) ImagePriceInfo(ctx context.Context, iReq *ImagePriceInfoRequest) (WujieCode, *ImagePriceInfoData, error) {
	resp, err := c.Client.ImagePriceInfo(ctx, iReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ImagePriceInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImagePriceInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, ImagePriceInfoRequest: %s",
			getTraceID(resp), err, iResp.Message, iReq.String())
	}
	return code, &iResp.Data, nil
}

// PostSuperSize create super size
func (c *Caller) PostSuperSize(ctx context.Context, pReq *PostSuperSizeRequest) (WujieCode, string, error) {
	resp, err := c.Client.PostSuperSize(ctx, pReq)
	if err != nil {
		return ErrorWujieCode, "", fmt.Errorf("c.Client.PostSuperSize: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var pResp PostSuperSizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return ErrorWujieCode, "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(pResp.Code)
	if err := code.Err(); err != nil {
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, PostSuperSizeRequest: %s",
			getTraceID(resp), err, pResp.Message, pReq.String())
	}
	return code, pResp.Data.Key, nil
}

// GetSuperSize get super size result
func (c *Caller) GetSuperSize(ctx context.Context, keys []string) (WujieCode, []SuperSizeInfo, error) {
	resp, err := c.Client.GetSuperSize(ctx, keys)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.GetSuperSize: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GetSuperSizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, keys: %v",
			getTraceID(resp), err, gResp.Message, keys)
	}
	return code, gResp.Data, nil
}

// CreateParams get create params
func (c *Caller) CreateParams(ctx context.Context, cReq *CreateParamsRequest) (WujieCode, []CreateParams, error) {
	resp, err := c.Client.CreateParams(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateParams: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateParamsResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateParamsRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data, nil
}

// ImageModelQueueInfo get image model queue info
func (c *Caller) ImageModelQueueInfo(ctx context.Context, model int32) (WujieCode, *ImageModelQueueInfoData, error) {
	resp, err := c.Client.ImageModelQueueInfo(ctx, model)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ImageModelQueueInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImageModelQueueInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, ImageModelQueueInfoRequest: %v",
			getTraceID(resp), err, iResp.Message, model)
	}
	return code, &iResp.Data, nil
}

// CancelImage cancel image
func (c *Caller) CancelImage(ctx context.Context, cReq *CancelImageRequest) (WujieCode, string, error) {
	resp, err := c.Client.CancelImage(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, "", fmt.Errorf("c.Client.CancelImage: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CancelImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CancelImageRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data, nil
}

// AccelerateImage accelerate image
func (c *Caller) AccelerateImage(ctx context.Context, aReq *AccelerateImageRequest) (WujieCode, bool, error) {
	resp, err := c.Client.AccelerateImage(ctx, aReq)
	if err != nil {
		return ErrorWujieCode, false, fmt.Errorf("c.Client.AccelerateImage: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var aResp AccelerateImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&aResp); err != nil {
		return ErrorWujieCode, false, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(aResp.Code)
	if err := code.Err(); err != nil {
		return code, false, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, AccelerateImageRequest: %s",
			getTraceID(resp), err, aResp.Message, aReq.String())
	}
	return code, true, nil
}

// PromptOptimizeSubmit submit prompt optimize
func (c *Caller) PromptOptimizeSubmit(ctx context.Context, pReq *PromptOptimizeSubmitRequest) (WujieCode, bool, error) {
	resp, err := c.Client.PromptOptimizeSubmit(ctx, pReq)
	if err != nil {
		return ErrorWujieCode, false, fmt.Errorf("c.Client.PromptOptimizeSubmit: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var pResp PromptOptimizeSubmitResponse
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return ErrorWujieCode, false, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(pResp.Code)
	if err := code.Err(); err != nil {
		return code, false, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, PromptOptimizeSubmitRequest: %s",
			getTraceID(resp), err, pResp.Message, pReq.String())
	}
	return code, true, nil
}

// PromptOptimizeResult get prompt optimize result
func (c *Caller) PromptOptimizeResult(ctx context.Context, taskID string) (WujieCode, *PromptOptimizeResultData, error) {
	resp, err := c.Client.PromptOptimizeResult(ctx, taskID)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.PromptOptimizeResult: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var pResp PromptOptimizeResultResponse
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(pResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, taskID: %s",
			getTraceID(resp), err, pResp.Message, taskID)
	}
	return code, &pResp.Data, nil
}

// Youthify youthify image
func (c *Caller) Youthify(ctx context.Context, yReq *YouthifyRequest) (WujieCode, *YouthifyData, error) {
	resp, err := c.Client.Youthify(ctx, yReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.Youthify: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var yResp YouthifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&yResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(yResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, YouthifyRequest: %s",
			getTraceID(resp), err, yResp.Message, yReq.String())
	}
	return code, &yResp.Data, nil
}

// QuerySpell query spell
func (c *Caller) QuerySpell(ctx context.Context) (WujieCode, []QuerySpellData, error) {
	resp, err := c.Client.QuerySpell(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.QuerySpell: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var qResp QuerySpellResponse
	if err := json.NewDecoder(resp.Body).Decode(&qResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(qResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, qResp.Message)
	}
	return code, qResp.Data, nil
}

// CreateImagePro create pro image
func (c *Caller) CreateImagePro(ctx context.Context, cReq *CreateImageProRequest) (WujieCode, []CreateImageProResult, error) {
	resp, err := c.Client.CreateImagePro(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateImagePro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateImageProResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateImageProRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data.Results, nil
}

// GeneratingInfoPro get pro image generating info
func (c *Caller) GeneratingInfoPro(ctx context.Context, gReq *GeneratingInfoProRequest) (WujieCode, []GeneratingInfoPro, error) {
	resp, err := c.Client.GeneratingInfoPro(ctx, gReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.GeneratingInfoPro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GeneratingInfoProResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, GeneratingInfoProRequest: %s",
			getTraceID(resp), err, gResp.Message, gReq.String())
	}
	return code, gResp.Data.Infos, nil
}

// CreateAvatarArtwork create avatar artwork
func (c *Caller) CreateAvatarArtwork(ctx context.Context, cReq *CreateAvatarArtworkRequest) (WujieCode, *CreateAvatarArtworkData, error) {
	resp, err := c.Client.CreateAvatarArtwork(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateAvatarArtwork: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateAvatarArtworkResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateAvatarArtworkRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data, nil
}

// AvatarDefaultResource get avatar default resource
func (c *Caller) AvatarDefaultResource(ctx context.Context) (WujieCode, *AvatarDefaultResource, error) {
	resp, err := c.Client.AvatarDefaultResource(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.AvatarDefaultResource: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var aResp AvatarDefaultResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&aResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(aResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, aResp.Message)
	}
	return code, &aResp.Data, nil
}

// CreateSpellAnalysis create spell analysis
func (c *Caller) CreateSpellAnalysis(ctx context.Context, cReq *CreateSpellAnalysisRequest) (WujieCode, string, error) {
	resp, err := c.Client.CreateSpellAnalysis(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, "", fmt.Errorf("c.Client.CreateSpellAnalysis: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateSpellAnalysisResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateSpellAnalysisRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data.Key, nil
}

// SpellAnalysisInfo get spell analysis info
func (c *Caller) SpellAnalysisInfo(ctx context.Context, key string) (WujieCode, *SpellAnalysisInfo, error) {
	resp, err := c.Client.SpellAnalysisInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.SpellAnalysisInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var sResp SpellAnalysisInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&sResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(sResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, sResp.Message, key)
	}
	return code, &sResp.Data, nil
}

func getTraceID(resp *http.Response) string {
	return resp.Header.Get(TraceID)
}
