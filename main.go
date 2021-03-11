package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/prateeknischal/osqueryexporter/internal/api"
	"github.com/prateeknischal/osqueryexporter/internal/client"
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

	timeoutDuration := time.Duration(*timeout) * time.Second

	var (
		c   client.Client
		err error
	)

	c, err = client.NewQsqueryClient(*socket, timeoutDuration)
	if err != nil {
		glog.Fatal(err)
	}
	api.Server(context.Background(), "localhost:5000", c)
}
