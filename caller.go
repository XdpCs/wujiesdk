package wujiesdk

// @Title        caller.go
// @Description  handle wujie sdk's response
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-10 20:47

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

func (c *Caller) AvailableIntegralBalance(ctx context.Context) (int, error) {
	resp, err := c.Client.AvailableIntegralBalance(ctx)
	if err != nil {
		return 0, fmt.Errorf("c.Client.AvailableIntegralBalance: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var bResp AvailableIntegralBalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&bResp); err != nil {
		return 0, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(bResp.Code)
	if err := code.Err(); err != nil {
		return 0, fmt.Errorf("TRACE_ID: %s, WujieCode: %w", getTraceID(resp), err)
	}

	return bResp.Data.Balance, nil
}

func (c *Caller) ExchangePoint(ctx context.Context, eReq *ExchangePointRequest) (bool, error) {
	resp, err := c.Client.ExchangePoint(ctx, eReq)
	if err != nil {
		return false, fmt.Errorf("c.Client.ExchangePoint: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var eResp BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&eResp); err != nil {
		return false, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(eResp.Code)
	if err := code.Err(); err != nil {
		return false, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, ExchangePointRequest: %s",
			getTraceID(resp), err, eReq.String())
	}
	return true, nil
}

func (c *Caller) ModelBaseInfos(ctx context.Context) ([]ModelBaseInfoData, error) {
	resp, err := c.Client.ModelBaseInfos(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.Client.ModelBaseInfos: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ModelBaseInfosResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w", getTraceID(resp), err)
	}

	return iResp.Data, nil
}

func (c *Caller) DefaultResourceStyleModel(ctx context.Context) ([]StyleModel, error) {
	resp, err := c.Client.DefaultResourceStyleModel(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.Client.DefaultResourceStyleModel: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp DefaultResourceStyleModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w", getTraceID(resp), err)
	}
	return mResp.Data.StyleModels, nil
}

func (c *Caller) DefaultResourceModel(ctx context.Context, model int32) (*DefaultResourceModelData, error) {
	resp, err := c.Client.DefaultResourceModel(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("c.Client.DefaultResourceModel: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var mResp DefaultResourceModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&mResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(mResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w", getTraceID(resp), err)
	}
	return &mResp.Data, nil
}

func (c *Caller) CreateImage(ctx context.Context, cReq *CreateImageRequest) (*CreateImageData, error) {
	resp, err := c.Client.CreateImage(ctx, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.Client.CreateImage: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, CreateImageRequest: %s",
			getTraceID(resp), err, cReq.String())
	}
	return &cResp.Data, nil
}

func (c *Caller) GeneratingInfo(ctx context.Context, gReq *GeneratingInfoRequest) ([]ImageGeneratingInfo, error) {
	resp, err := c.Client.GeneratingInfo(ctx, gReq)
	if err != nil {
		return nil, fmt.Errorf("c.Client.GeneratingInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GeneratingInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, GeneratingInfoRequest: %s",
			getTraceID(resp), err, gReq.String())
	}
	return gResp.Data.List, nil
}

func (c *Caller) ImageInfo(ctx context.Context, key string) (*ImageInfoData, error) {
	resp, err := c.Client.ImageInfo(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("c.Client.ImageInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImageInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, key: %s",
			getTraceID(resp), err, key)
	}
	return &iResp.Data, nil
}

func (c *Caller) ImagePriceInfo(ctx context.Context, iReq *ImagePriceInfoRequest) (*ImagePriceInfoData, error) {
	resp, err := c.Client.ImagePriceInfo(ctx, iReq)
	if err != nil {
		return nil, fmt.Errorf("c.Client.ImagePriceInfo: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var iResp ImagePriceInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&iResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(iResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, ImagePriceInfoRequest: %s",
			getTraceID(resp), err, iReq.String())
	}
	return &iResp.Data, nil
}

func (c *Caller) PostSuperSize(ctx context.Context, sReq *PostSuperSizeRequest) (string, error) {
	resp, err := c.Client.PostSuperSize(ctx, sReq)
	if err != nil {
		return "", fmt.Errorf("c.Client.PostSuperSize: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var pResp PostSuperSizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(pResp.Code)
	if err := code.Err(); err != nil {
		return "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, PostSuperSizeRequest: %s",
			getTraceID(resp), err, sReq.String())
	}
	return pResp.Data.Key, nil
}

func (c *Caller) GetSuperSize(ctx context.Context, keys []string) ([]SuperSizeInfo, error) {
	resp, err := c.Client.GetSuperSize(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("c.Client.GetSuperSize: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GetSuperSizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, keys: %v",
			getTraceID(resp), err, keys)
	}
	return gResp.Data, nil
}

func (c *Caller) CancelImage(ctx context.Context, cReq *CancelImageRequest) (string, error) {
	resp, err := c.Client.CancelImage(ctx, cReq)
	if err != nil {
		return "", fmt.Errorf("c.Client.CancelImage: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CancelImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return "", fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return "", fmt.Errorf("TRACE_ID: %s, WujieCode: %w, CancelImageRequest: %s",
			getTraceID(resp), err, cReq.String())
	}
	return cResp.Data, nil
}

func (c *Caller) CreateImagePro(ctx context.Context, cReq *CreateImageProRequest) (*CreateImageProResponse, error) {
	resp, err := c.Client.CreateImagePro(ctx, cReq)
	if err != nil {
		return nil, fmt.Errorf("c.Client.CreateImagePro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var cResp CreateImageProResponse
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(cResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, CreateImageProRequest: %s",
			getTraceID(resp), err, cReq.String())
	}
	return &cResp, nil
}

func (c *Caller) GeneratingInfoPro(ctx context.Context, gReq *GeneratingInfoProRequest) (*GeneratingInfoProResponse, error) {
	resp, err := c.Client.GeneratingInfoPro(ctx, gReq)
	if err != nil {
		return nil, fmt.Errorf("c.Client.GeneratingInfoPro: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var gResp GeneratingInfoProResponse
	if err := json.NewDecoder(resp.Body).Decode(&gResp); err != nil {
		return nil, fmt.Errorf("json.NewDecoder: %w", err)
	}
	code := WujieCode(gResp.Code)
	if err := code.Err(); err != nil {
		return nil, fmt.Errorf("TRACE_ID: %s, WujieCode: %w, GeneratingInfoProRequest: %s",
			getTraceID(resp), err, gReq.String())
	}
	return &gResp, nil
}

func getTraceID(resp *http.Response) string {
	return resp.Header.Get(TraceID)
}
