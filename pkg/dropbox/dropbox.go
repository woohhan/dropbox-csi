package dropbox

import (
	"fmt"

	"github.com/golang/glog"
)

type dropbox struct {
	name     string
	nodeID   string
	version  string
	endpoint string

	ids *identityServer
	ns  *nodeServer
	cs  *controllerServer
}

func NewDropboxDriver(driverName, nodeID, endpoint, version string) (*dropbox, error) {
	if driverName == "" {
		return nil, fmt.Errorf("No driver name provided")
	}

	if nodeID == "" {
		return nil, fmt.Errorf("No node id provided")
	}

	if endpoint == "" {
		return nil, fmt.Errorf("No driver endpoint provided")
	}

	glog.Infof("Driver: %v ", driverName)
	glog.Infof("Version: %s", version)

	return &dropbox{
		name:     driverName,
		version:  version,
		nodeID:   nodeID,
		endpoint: endpoint,
	}, nil
}

func (d *dropbox) Run() {
	// Create GRPC servers
	d.ids = NewIdentityServer(d.name, d.version)
	d.ns = NewNodeServer(d.nodeID)
	d.cs = NewControllerServer(d.nodeID)

	s := NewNonBlockingGRPCServer()
	s.Start(d.endpoint, d.ids, d.cs, d.ns)
	s.Wait()
}
