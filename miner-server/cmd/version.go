package cmd

import "os"

var CurrentCommit string
var BuildType int

const (
	BuildDefault = 0
	BuildMainnet = 0x1
	BuildTestnet = 0x2
)

func BuildTypeString() string {
	switch BuildType {
	case BuildDefault:
		return ""
	case BuildMainnet:
		return "+mainnet"
	case BuildTestnet:
		return "+testnet"
	default:
		return "+huh?"
	}
}

// BuildVersion is the local build version
const BuildVersion = "1.0.0-dev"

func UserVersion() string {
	if os.Getenv("UTILITY_VERSION_IGNORE_COMMIT") == "1" {
		return BuildVersion
	}

	return BuildVersion + BuildTypeString() + CurrentCommit
}
