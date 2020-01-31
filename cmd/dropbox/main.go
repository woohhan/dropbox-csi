package main

import (
	"github.com/woohhan/dropbox-csi/pkg/dropbox"

	"flag"
	"fmt"
	"os"
	"path"
)

const (
	version = "v1.0.0"
)

var (
	// TODO: change endpoint
	endpoint    = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	driverName  = flag.String("drivername", "dropbox.csi.k8s.io", "name of the driver")
	nodeID      = flag.String("nodeid", "", "node id")
	showVersion = flag.Bool("version", false, "Show version.")
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	if *showVersion {
		baseName := path.Base(os.Args[0])
		fmt.Println(baseName, version)
		return
	}

	handle()
	os.Exit(0)
}

func handle() {
	driver, err := dropbox.NewDropboxDriver(*driverName, *nodeID, *endpoint, version)
	if err != nil {
		fmt.Printf("Failed to initialize driver: %s", err.Error())
		os.Exit(1)
	}
	driver.Run()
}
