package constants

import "fmt"

var (
	Version = "0.0.1"
	Build   = "dirty"
	Commit  = "0000000"
)

func GetVersion() string {
	return fmt.Sprintf("%s-%s+%s", Version, Build, Commit)
}
