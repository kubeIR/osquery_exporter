package osqueryexporter

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/kolide/osquery-go"
	"github.com/prateeknischal/osqueryexporter/internal/constants"
)

var (
	socket   = flag.String("socket", "", "osquery Socket file")
	timeout  = flag.Int("timeout", 3, "Timeout to connect to socket file")
	interval = flag.Int("interval", 3, "Delay between connectivity checks")
	verbose  = flag.Bool("verbose", false, "Enable verbose output")
	version  = flag.Bool("version", false, "Version")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Fprintln(os.Stdout, constants.Version())
		os.Exit(0)
	}

	_, err := osquery.NewExtensionManagerServer("foo", *socket)
	if err != nil {
		glog.Fatalf("Failed to initialize the extension", err)
	}
}
