package wujiesdk

// @Title        caller.go
// @Description  handle wujie sdk's response
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2024-01-29 14:18

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
func (c *Caller) ModelBaseInfos(ctx context.Context) (WujieCode, []ModelBaseInfo, error) {
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
func (c *Caller) GeneratingInfo(ctx context.Context, keys []string) (WujieCode, []ImageGeneratingInfo, error) {
	resp, err := c.Client.GeneratingInfo(ctx, keys)
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
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, keys: %v",
			getTraceID(resp), err, gResp.Message, keys)
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
func (c *Caller) CreateParams(ctx context.Context, keys []string) (WujieCode, []CreateParams, error) {
	resp, err := c.Client.CreateParams(ctx, keys)
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
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, keys: %v",
			getTraceID(resp), err, cResp.Message, keys)
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
func (c *Caller) CancelImage(ctx context.Context, key string) (WujieCode, string, error) {
	resp, err := c.Client.CancelImage(ctx, key)
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
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, cResp.Message, key)
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
func (c *Caller) GeneratingInfoPro(ctx context.Context, keys []string) (WujieCode, []GeneratingInfoPro, error) {
	resp, err := c.Client.GeneratingInfoPro(ctx, keys)
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
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, keys: %v",
			getTraceID(resp), err, gResp.Message, keys)
	}
	return code, gResp.Data.Infos, nil
}

// AccountBalancePro get account balance pro
func (c *Caller) AccountBalancePro(ctx context.Context) (WujieCode, int, error) {
	resp, err := c.Client.AccountBalancePro(ctx)
	if err != nil {
		return ErrorWujieCode, 0, fmt.Errorf("c.Client.AccountBalancePro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var aResp AccountBalanceProResponse
	if err := json.NewDecoder(resp.Body).Decode(&aResp); err != nil {
		return ErrorWujieCode, 0, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(aResp.Code)
	if err := code.Err(); err != nil {
		return code, 0, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, aResp.Message)
	}
	return code, aResp.Data.ResourceBalance, nil
}

// ModelBaseInfosPro get model base infos pro
func (c *Caller) ModelBaseInfosPro(ctx context.Context) (WujieCode, []ModelBaseInfoPro, error) {
	resp, err := c.Client.ModelBaseInfosPro(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ModelBaseInfosPro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp ModelBaseInfosProResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, mResp.Message)
	}
	return code, mResp.Data, nil
}

// ControlNetOptionPro control net option pro
func (c *Caller) ControlNetOptionPro(ctx context.Context) (WujieCode, []ControlNetOptionPro, error) {
	resp, err := c.Client.ControlNetOptionPro(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ControlNetOptionPro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp ControlNetOptionProResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, cResp.Message)
	}
	return code, cResp.Data, nil
}

// ImageInfoPro get image info pro
func (c *Caller) ImageInfoPro(ctx context.Context, key string) (WujieCode, *ImageInfoPro, error) {
	resp, err := c.Client.ImageInfoPro(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ImageInfoPro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImageInfoProResponse
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

// CreateAvatar create avatar
func (c *Caller) CreateAvatar(ctx context.Context, cReq *CreateAvatarRequest) (WujieCode, *CreateAvatarData, error) {
	resp, err := c.Client.CreateAvatar(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateAvatar: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateAvatarResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateAvatarRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data, nil
}

// DeleteAvatar delete avatar
func (c *Caller) DeleteAvatar(ctx context.Context, key string) (WujieCode, bool, error) {
	resp, err := c.Client.DeleteAvatar(ctx, key)
	if err != nil {
		return ErrorWujieCode, false, fmt.Errorf("c.Client.DeleteAvatar: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var dResp BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&dResp); err != nil {
		return ErrorWujieCode, false, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(dResp.Code)
	if err := code.Err(); err != nil {
		return code, false, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, dResp.Message, key)
	}
	return code, true, nil
}

// AvatarInfo get avatar info
func (c *Caller) AvatarInfo(ctx context.Context, key string) (WujieCode, *AvatarInfoData, error) {
	resp, err := c.Client.AvatarInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.AvatarInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var aResp AvatarInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&aResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(aResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, aResp.Message, key)
	}
	return code, &aResp.Data, nil
}

// ImageBatchCheck image batch check
func (c *Caller) ImageBatchCheck(ctx context.Context, imageURLList []string) (WujieCode, []ImageCheckInfo, error) {
	resp, err := c.Client.ImageBatchCheck(ctx, imageURLList)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.ImageBatchCheck: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImageBatchCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, iResp.Message)
	}
	return code, iResp.Data.ImageCheckInfoList, nil
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

// MagicDiceTheme get magic dice theme
func (c *Caller) MagicDiceTheme(ctx context.Context) (WujieCode, []MagicDiceTheme, error) {
	resp, err := c.Client.MagicDiceTheme(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.MagicDiceTheme: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp MagicDiceThemeResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s", getTraceID(resp), err, mResp.Message)
	}
	return code, mResp.Data, nil
}

// CreateMagicDice create magic dice
func (c *Caller) CreateMagicDice(ctx context.Context, cReq *CreateMagicDiceRequest) (WujieCode, *CreateMagicDiceResult, error) {
	resp, err := c.Client.CreateMagicDice(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateMagicDice: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateMagicDiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateMagicDiceRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data, nil
}

// CreateVideo create video
func (c *Caller) CreateVideo(ctx context.Context, cReq *CreateVideoRequest) (WujieCode, string, error) {
	resp, err := c.Client.CreateVideo(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, "", fmt.Errorf("c.Client.CreateVideo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateVideoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateVideoRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data.Key, nil
}

// VideoInfo get video info
func (c *Caller) VideoInfo(ctx context.Context, key string) (WujieCode, *VideoInfo, error) {
	resp, err := c.Client.VideoInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.VideoInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var vResp VideoInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(vResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, vResp.Message, key)
	}
	return code, &vResp.Data, nil
}

// VideoOptionMenuAndPriceTable get video option menu and price table
func (c *Caller) VideoOptionMenuAndPriceTable(ctx context.Context) (WujieCode, *VideoOptionMenuAndPriceTable, error) {
	resp, err := c.Client.VideoOptionMenuAndPriceTable(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.VideoOptionMenuAndPriceTable: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var vResp VideoOptionMenuAndPriceTableResponse
	if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(vResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, vResp.Message)
	}
	return code, &vResp.Data, nil
}

// VideoModelQueueInfo get video queue info
func (c *Caller) VideoModelQueueInfo(ctx context.Context, model int32) (WujieCode, *VideoModelQueueInfo, error) {
	resp, err := c.Client.VideoModelQueueInfo(ctx, model)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.VideoModelQueueInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var vResp VideoModelQueueInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(vResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, vResp.Message)
	}
	return code, &vResp.Data, nil
}

// VideoGeneratingInfo get video generating info
func (c *Caller) VideoGeneratingInfo(ctx context.Context, keys []string) (WujieCode, *VideoGeneratingInfo, error) {
	resp, err := c.Client.VideoGeneratingInfo(ctx, keys)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.VideoGeneratingInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var vResp VideoGeneratingInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(vResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, vResp.Message)
	}
	return code, &vResp.Data, nil
}

// CameraTemplateOptions get camera template options
func (c *Caller) CameraTemplateOptions(ctx context.Context) (WujieCode, []CameraTemplateOption, error) {
	resp, err := c.Client.CameraTemplateOptions(ctx)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CameraTemplateOptions: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CameraTemplateOptionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, cResp.Message)
	}
	return code, cResp.Data, nil
}

// CreateCamera create camera
func (c *Caller) CreateCamera(ctx context.Context, cReq *CreateCameraRequest) (WujieCode, *CreateCameraResult, error) {
	resp, err := c.Client.CreateCamera(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateCamera: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateCameraResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateCameraRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data, nil
}

// CameraGeneratingInfo get camera generating info
func (c *Caller) CameraGeneratingInfo(ctx context.Context, keys []string) (WujieCode, []CameraGeneratingInfo, error) {
	resp, err := c.Client.CameraGeneratingInfo(ctx, keys)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CameraGeneratingInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CameraGeneratingInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, cResp.Message)
	}
	return code, cResp.Data.Infos, nil
}

// CameraInfo get camera info
func (c *Caller) CameraInfo(ctx context.Context, key string) (WujieCode, *CameraInfo, error) {
	resp, err := c.Client.CameraInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CameraInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CameraInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, key: %s",
			getTraceID(resp), err, cResp.Message, key)
	}
	return code, &cResp.Data, nil
}

// LabOptions get lab options
func (c *Caller) LabOptions(ctx context.Context, lReq *LabOptionsRequest) (WujieCode, []LabOption, error) {
	resp, err := c.Client.LabOptions(ctx, lReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.LabOptions: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var lResp LabOptionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&lResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(lResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, lResp.Message)
	}
	return code, lResp.Data.AiLabQuery.Options, nil
}

// LabInfo get lab info
func (c *Caller) LabInfo(ctx context.Context, lReq *LabInfoRequest) (WujieCode, *LabInfo, error) {
	resp, err := c.Client.LabInfo(ctx, lReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.LabInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var lResp LabInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&lResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(lResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s",
			getTraceID(resp), err, lResp.Message)
	}
	return code, &lResp.Data, nil
}

// CreateSegmentation create segmentation
func (c *Caller) CreateSegmentation(ctx context.Context, cReq *CreateSegmentationRequest) (WujieCode, *CreateSegmentationResult, error) {
	resp, err := c.Client.CreateSegmentation(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateSegmentation: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateSegmentationResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateSegmentationRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data.AiLabMutation.SegmentAnythingCreateV2, nil
}

// CreateInfiniteZoom create infinite zoom
func (c *Caller) CreateInfiniteZoom(ctx context.Context, cReq *CreateInfiniteZoomRequest) (WujieCode, *CreateInfiniteZoomResult, error) {
	resp, err := c.Client.CreateInfiniteZoom(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateInfiniteZoom: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateInfiniteZoomResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateInfiniteZoomRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data.AiLabMutation.InfiniteZoomCreateV2, nil
}

// CreateVectorStudio create vector studio
func (c *Caller) CreateVectorStudio(ctx context.Context, cReq *CreateVectorStudioRequest) (WujieCode, *CreateVectorStudioResult, error) {
	resp, err := c.Client.CreateVectorStudio(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateVectorStudio: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateVectorStudioResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateVectorStudioRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp.Data.AiLabMutation.VectorStudioCreateV2, nil
}

// CreateSVD creates svd
func (c *Caller) CreateSVD(ctx context.Context, cReq *CreateSVDRequest) (WujieCode, string, error) {
	resp, err := c.Client.CreateSVD(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, "", fmt.Errorf("c.Client.CreateSVD: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateSVDResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, "", fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateSVDRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, cResp.Data.Key, nil
}

// SVDInfo get svd info
func (c *Caller) SVDInfo(ctx context.Context, key string) (WujieCode, *SVDInfo, error) {
	resp, err := c.Client.SVDInfo(ctx, key)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.SVDInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var sResp SVDInfoResponse
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

// CreateMidjourney create midjourney image
func (c *Caller) CreateMidjourney(ctx context.Context, cReq *CreateMidjourneyRequest) (WujieCode, *CreateMidjourneyResponse, error) {
	resp, err := c.Client.CreateMidjourney(ctx, cReq)
	if err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("c.Client.CreateMidjourney: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateMidjourneyResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return ErrorWujieCode, nil, fmt.Errorf("json.NewDecoder: %w", err)
	}

	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return code, nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, Message: %s, CreateMidjourneyRequest: %s",
			getTraceID(resp), err, cResp.Message, cReq.String())
	}
	return code, &cResp, nil
}

func getTraceID(resp *http.Response) string {
	return resp.Header.Get(TraceID)
}
