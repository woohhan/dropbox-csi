package dropbox

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
)

type identityServer struct {
	name    string
	version string
}

func NewIdentityServer(name, version string) *identityServer {
	return &identityServer{
		name:    name,
		version: version,
	}
}

func (i identityServer) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	panic("implement me")
}

func (i identityServer) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	panic("implement me")
}

func (i identityServer) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	panic("implement me")
}
