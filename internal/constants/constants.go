package constants

import "fmt"

var (
	Version = "0.0.1"
	Commit  = "0000000"
)

func GetVersion() string {
	return fmt.Sprintf("%s+%s", Version, Commit)
}
