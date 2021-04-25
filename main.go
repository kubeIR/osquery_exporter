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
)

var versionString = "unknown"

var (
	socket   = flag.String("socket", "", "osquery Socket file")
	timeout  = flag.Int("timeout", 3, "Timeout to connect to socket file")
	interval = flag.Int("interval", 3, "Delay between connectivity checks")
	verbose  = flag.Bool("verbose", false, "Enable verbose output")
	addr     = flag.String("addr", "0.0.0.0:5000", "host:port to listen on")
	version  = flag.Bool("version", false, "Get Version")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Fprintln(os.Stdout, versionString)
		os.Exit(0)
	}

	timeoutDuration := time.Duration(*timeout) * time.Second

	var (
		c   client.Client
		err error
	)

	glog.Infof("Starting with the socket file: %s", *socket)
	c, err = client.NewOsqueryClient(*socket, timeoutDuration)
	if err != nil {
		glog.Fatal(err)
	}
	api.Server(context.Background(), *addr, c)
}
