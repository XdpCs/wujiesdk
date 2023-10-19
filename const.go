package wujiesdk

// @Title        const.go
// @Description  wujie sdk's const
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-18 21:18

import (
	"fmt"
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
	DefaultResourceModelWujieRouter      WujieRouter = "/ai/default_resource"
	ModelBaseInfosWujieRouter            WujieRouter = "/ai/model_base_infos"
	GeneratingInfoWujieRouter            WujieRouter = "/ai/generating_info"
	CreateParamsWujieRouter              WujieRouter = "/ai/create_params"
	ImageInfoWujieRouter                 WujieRouter = "/ai/info"
	ImageModelQueueInfoWujieRouter       WujieRouter = "/ai/model_info"
	DefaultResourceStyleModelWujieRouter WujieRouter = "/ai/default_resource_style_model"
	CreateImageWujieRouter               WujieRouter = "/ai/create"
	AccelerateImageWujieRouter           WujieRouter = "/ai/accelerate"
	CancelImageWujieRouter               WujieRouter = "/ai/cancel"
	ImagePriceInfoWujieRouter            WujieRouter = "/ai/price_info"
	SuperSizeWujieRouter                 WujieRouter = "/ai/supersize"
	PromptOptimizeSubmitWujieRouter      WujieRouter = "/ai/optimize/prompt/submit"
	PromptOptimizeResultWujieRouter      WujieRouter = "/ai/optimize/prompt/result"
	YouthifyWujieRouter                  WujieRouter = "/ai/youthify"
	QuerySpellWujieRouter                WujieRouter = "/ai/spell"
)

// pro ai WujieRouter
const (
	CreateImageProWujieRouter    WujieRouter = "/ai/pro/create"
	GeneratingInfoProWujieRouter WujieRouter = "/ai/pro/generating_info"
)

// avatar create artwork WujieRouter
const (
	CreateAvatarArtworkWujieRouter   WujieRouter = "/avatar/artwork/create"
	AvatarDefaultResourceWujieRouter WujieRouter = "/avatar/artwork/default_resource"
)

// spell analysis WujieRouter
const (
	CreateSpellAnalysisWujieRouter WujieRouter = "/ai/spell_analysis/create"
	SpellAnalysisInfoWujieRouter   WujieRouter = "/ai/spell_analysis/info"
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
	ErrorWujieCode                               WujieCode = "0"
	OKWujieCode                                  WujieCode = "200"
	InvalidParameterWujieCode                    WujieCode = "20010001"
	UnsupportedResolutionWujieCode               WujieCode = "20010015"
	LockRaceConditionWujieCode                   WujieCode = "20010018"
	PromptTranslationFailedWujieCode             WujieCode = "20110000"
	PromptContainsSensitiveWordsWujieCode        WujieCode = "20110001"
	InitImageLinkIncorrectOrUnsupportedWujieCode WujieCode = "20110002"
	InitImageContainsSensitiveInfoWujieCode      WujieCode = "20110003"
	ImageStatusChange                            WujieCode = "20110009"
	InsufficientPointsBalanceWujieCode           WujieCode = "20110010"
	JobNotInQueueAndCannotCancelWujieCode        WujieCode = "20110011"
	CheckResourcesWujieCode                      WujieCode = "20110017"
	ImageRecognitionAbnormalityWujieCode         WujieCode = "20110018"
	NoFaceOrFaceIsSmallWujieCode                 WujieCode = "20110019"
	MultipleFacesDetectedWujieCode               WujieCode = "20110020"
	SideFaceDetectedWujieCode                    WujieCode = "20110021"
)

func (w WujieCode) String() string {
	switch w {
	case InvalidParameterWujieCode:
		return "非法参数"
	case UnsupportedResolutionWujieCode:
		return "暂时无法支持的尺寸/分辨率"
	case LockRaceConditionWujieCode:
		return "由锁竞争导致的作画失败（需要重新发起）"
	case PromptTranslationFailedWujieCode:
		return "文本语言翻译失败"
	case PromptContainsSensitiveWordsWujieCode:
		return "画面描述含有敏感词"
	case InitImageLinkIncorrectOrUnsupportedWujieCode:
		return "底图链接信息有误或不支持"
	case InitImageContainsSensitiveInfoWujieCode:
		return "参考图含有敏感信息"
	case InsufficientPointsBalanceWujieCode:
		return "积分余额不足"
	case CheckResourcesWujieCode:
		return "检测资源"
	case ImageRecognitionAbnormalityWujieCode:
		return "图片识别异常"
	case NoFaceOrFaceIsSmallWujieCode:
		return "未检测到人脸或人脸太小"
	case MultipleFacesDetectedWujieCode:
		return "检测到多张人脸"
	case SideFaceDetectedWujieCode:
		return "检测到侧脸"
	case JobNotInQueueAndCannotCancelWujieCode:
		return "该作品不在排队中，无法撤销"
	case ErrorWujieCode:
		return "非无界报错返回"
	case ImageStatusChange:
		return "作品状态改变，请刷新后查看"
	default:
		return fmt.Sprintf("code: %v 未知错误", string(w))
	}
}

func (w WujieCode) Err() error {
	if w != OKWujieCode {
		return errors.New(w.String())
	}
	return nil
}
