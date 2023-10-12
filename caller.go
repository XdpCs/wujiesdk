package wujiesdk

// @Title        caller.go
// @Description  handle wujie sdk's response
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-12 14:47

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Caller struct {
	Client *Client
}

func NewCaller(c *Client) *Caller {
	return &Caller{Client: c}
}

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

func getTraceID(resp *http.Response) string {
	return resp.Header.Get(TraceID)
}
