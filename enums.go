package wujiesdk

// @Title        enums.go
// @Description
// @Create       XdpCs 2023-10-17 20:48
// @Update       XdpCs 2023-10-17 20:48

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
