package wujiesdk

// @Title        enums.go
// @Description  enums
// @Create       XdpCs 2023-10-17 20:48
// @Update       XdpCs 2023-11-26 15:08

// PromptSubmitType prompt submit type /ai/optimize/prompt/submit
type PromptSubmitType int8

const (
	CommonPromptSubmitType PromptSubmitType = iota + 1
	ColorPromptSubmitType
	AnimePromptSubmitType
)

// PromptSubmitLanguage prompt submit language /ai/optimize/prompt/submit
type PromptSubmitLanguage int8

const (
	ChinesePromptSubmitLanguage PromptSubmitLanguage = iota
	EnglishPromptSubmitLanguage
)

// CreateMagicDiceType create magic dice type /ai/magic_dice/search
type CreateMagicDiceType string

const (
	MagicDicCreateMagicDiceType   CreateMagicDiceType = "MAGIC_DICE"
	SmartMagicCreateMagicDiceType CreateMagicDiceType = "SMART_MAGIC"
)

// CreateMagicDiceModel create magic dice model /ai/magic_dice/search
type CreateMagicDiceModel string

const (
	StableDiffusionCreateMagicDiceModel CreateMagicDiceModel = "STABLE_DIFFUSION"
	AnimeDiffusionCreateMagicDiceModel  CreateMagicDiceModel = "ANIME_DIFFUSION"
	StyleCreateMagicDiceModel           CreateMagicDiceModel = "STYLE"
)

// CreateMagicDiceLanguage create magic dice language /ai/magic_dice/search
type CreateMagicDiceLanguage string

const (
	ChineseCreateMagicDiceLanguage CreateMagicDiceLanguage = "CHINESE"
	EnglishCreateMagicDiceLanguage CreateMagicDiceLanguage = "ENGLISH"
)

// SuperSizeType super size type /ai/supersize
type SuperSizeType string

const (
	GeneralSuperSizeType SuperSizeType = "GENERAL"
	AnimeSuperSizeType   SuperSizeType = "ANIME"
)

// SuperSizeCostType super size cost type /ai/supersize
type SuperSizeCostType string

const (
	IntegralSuperSizeCostType SuperSizeCostType = "INTEGRAL"
	DurationSuperSizeCostType SuperSizeCostType = "DURATION"
)

// LabOptionType lab option type /ai/pro/lab/options
type LabOptionType string

const (
	InfiniteZoomModelLabOptionType    LabOptionType = "INFINITE_ZOOM_MODEL"
	InfiniteZoomSamplerLabOptionType  LabOptionType = "INFINITE_ZOOM_SAMPLER"
	SegmentAnythingModelLabOptionType LabOptionType = "SEGMENT_ANYTHING_MODEL"
	VectorStudioStyleLabOptionType    LabOptionType = "VECTOR_STUDIO_STYLE"
)

// LabInfoType lab info type /ai/pro/lab/info
type LabInfoType string

const (
	InfiniteZoomLabInfoType LabInfoType = "AI_LAB_INFINITE_ZOOM"
	SegmentationLabInfoType LabInfoType = "AI_LAB_SEGMENTATION"
	MjDescribeLabInfoType   LabInfoType = "MJ_DESCRIBE"
	PictureLabInfoType      LabInfoType = "PICTURE"
	SuperSizeLabInfoType    LabInfoType = "SUPER_SIZE"
	VectorLabInfoType       LabInfoType = "VECTOR"
	VideoLabInfoType        LabInfoType = "VIDEO"
)
