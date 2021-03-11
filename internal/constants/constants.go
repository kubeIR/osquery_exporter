package constants

import "fmt"

var (
	version = "0.0.1"
	build   = "dirty"
	commit  = "0000000"
)

func Version() string {
	return fmt.Sprintf("%s-%s+%s", version, build, commit)
}
