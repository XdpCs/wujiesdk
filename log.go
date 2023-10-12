package wujiesdk

// @Title        log.go
// @Description  log
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-10 20:47

// Define the level of the output log
const (
	LogOff = iota
	LogError
	LogWarn
	LogInfo
	LogDebug
)

// LogTag Tag for each level of log
var LogTag = []string{"[error]", "[warn]", "[info]", "[debug]"}
