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
	ver      = flag.Bool("version", false, "Version")
)

func main() {
	flag.Parse()
	if *ver {
		fmt.Fprintln(os.Stdout, constants.GetVersion())
		os.Exit(0)
	}

	timeoutDuration := time.Duration(*timeout) * time.Second

	var (
		c   client.Client
		err error
	)

	c, err = client.NewOsqueryClient(*socket, timeoutDuration)
	if err != nil {
		glog.Fatal(err)
	}
	api.Server(context.Background(), "0.0.0.0:5000", c)
}
