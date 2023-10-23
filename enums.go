package wujiesdk

// @Title        enums.go
// @Description  enums
// @Create       XdpCs 2023-10-17 20:48
// @Update       XdpCs 2023-10-23 15:08

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
	GeneralSuperSizeType = "GENERAL"
	AnimeSuperSizeType   = "ANIME"
)

// SuperSizeCostType super size cost type /ai/supersize
type SuperSizeCostType string

const (
	IntegralSuperSizeCostType = "INTEGRAL"
	DurationSuperSizeCostType = "DURATION"
)
