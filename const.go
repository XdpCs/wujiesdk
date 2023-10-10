package wujiesdk

// @Title        const.go
// @Description  wujie sdk's const
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-10 20:47

import (
	"time"

	"errors"
)

type WujieRouter string

// account WujieRouter
const (
	AvailableIntegralBalanceWujieRouter WujieRouter = "/account/availableIntegralBalance"
	ExchangePointWujieRouter            WujieRouter = "/account/integral/exchange"
)

// common ai WujieRouter
const (
	ModelBaseInfosWujieRouter            WujieRouter = "/ai/model_base_infos"
	DefaultResourceStyleModelWujieRouter WujieRouter = "/ai/default_resource_style_model"
	DefaultResourceModelWujieRouter      WujieRouter = "/ai/default_resource"
	ImagePriceInfoWujieRouter            WujieRouter = "/ai/price_info"
	CreateImageWujieRouter               WujieRouter = "/ai/create"
	GeneratingInfoWujieRouter            WujieRouter = "/ai/generating_info"
	ImageInfoWujieRouter                 WujieRouter = "/ai/info"
	SuperSizeWujieRouter                 WujieRouter = "/ai/supersize"
)

// image queue
const (
	AccelerateWujieRouter  WujieRouter = "/ai/accelerate"
	CancelImageWujieRouter WujieRouter = "/ai/cancel"
	ImageModelQueueInfo    WujieRouter = "/ai/model_info"
)

// pro ai WujieRouter
const (
	CreateImageProWujieRouter    WujieRouter = "/ai/pro/create"
	GeneratingInfoProWujieRouter WujieRouter = "/ai/pro/generating_info"
)

const DefaultExpiration time.Duration = 4 * time.Minute
const Domain string = "https://gate.wujiebantu.com/wj-open/v1"

const (
	ContentType                   string = "Content-Type"
	ApplicationFormUrlencodedUTF8 string = "application/x-www-form-urlencoded;charset=UTF-8"
	ApplicationJson               string = "application/json"
	HTTPHeaderAuthorization       string = "Authorization"
)
const TraceID string = "TRACE_ID"

type WujieCode string

const (
	OK                                  WujieCode = "200"
	InvalidParameter                    WujieCode = "20010001"
	UnsupportedResolution               WujieCode = "20010015"
	LockRaceCondition                   WujieCode = "20010018"
	PromptTranslationFailed             WujieCode = "20110000"
	PromptContainsSensitiveWords        WujieCode = "20110001"
	InitImageLinkIncorrectOrUnsupported WujieCode = "20110002"
	InitImageContainsSensitiveInfo      WujieCode = "20110003"
	InsufficientPointsBalance           WujieCode = "20110010"
	CheckResources                      WujieCode = "20110017"
	ImageRecognitionAbnormality         WujieCode = "20110018"
	NoFaceOrFaceIsSmall                 WujieCode = "20110019"
	MultipleFacesDetected               WujieCode = "20110020"
	SideFaceDetected                    WujieCode = "20110021"
)

func (w WujieCode) String() string {
	switch w {
	case InvalidParameter:
		return "非法参数"
	case UnsupportedResolution:
		return "暂时无法支持的尺寸/分辨率"
	case LockRaceCondition:
		return "由锁竞争导致的作画失败（需要重新发起）"
	case PromptTranslationFailed:
		return "文本语言翻译失败"
	case PromptContainsSensitiveWords:
		return "画面描述含有敏感词"
	case InitImageLinkIncorrectOrUnsupported:
		return "底图链接信息有误或不支持"
	case InitImageContainsSensitiveInfo:
		return "参考图含有敏感信息"
	case InsufficientPointsBalance:
		return "积分余额不足"
	case CheckResources:
		return "检测资源"
	case ImageRecognitionAbnormality:
		return "图片识别异常"
	case NoFaceOrFaceIsSmall:
		return "未检测到人脸或人脸太小"
	case MultipleFacesDetected:
		return "检测到多张人脸"
	case SideFaceDetected:
		return "检测到侧脸"
	default:
		return "未知错误"
	}
}

func (w WujieCode) Err() error {
	if w != OK {
		return errors.New(w.String())
	}
	return nil
}
