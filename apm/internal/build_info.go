package internal

import (
	"os"
	"path/filepath"
)

var (
	hostName string
	appName  string
)

func init() {
	hostName, _ = os.Hostname()
	appName = filepath.Base(os.Args[0])
}

type buildInfo struct {
}

var BuildInfo = &buildInfo{}

func (b buildInfo) AppName() string {
	return appName
}

func (b buildInfo) HostName() string {
	return hostName
}
