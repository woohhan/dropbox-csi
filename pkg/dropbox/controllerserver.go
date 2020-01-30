package dropbox

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

type controllerServer struct {
	nodeID string
}

func NewControllerServer(nodeID string) *controllerServer {
	return &controllerServer{
		nodeID: nodeID,
	}
}

func (c controllerServer) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: getControllerServiceCapabilities(
			[]csi.ControllerServiceCapability_RPC_Type{}),
	}, nil
}

func getControllerServiceCapabilities(cl []csi.ControllerServiceCapability_RPC_Type) []*csi.ControllerServiceCapability {
	var csc []*csi.ControllerServiceCapability

	for _, cap := range cl {
		glog.Infof("Enabling controller service capability: %v", cap.String())
		csc = append(csc, &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		})
	}

	return csc
}

func (c controllerServer) CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	panic("implement me")
}

func (c controllerServer) DeleteVolume(context.Context, *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	panic("implement me")
}

func (c controllerServer) ControllerPublishVolume(context.Context, *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	panic("implement me")
}

func (c controllerServer) ControllerUnpublishVolume(context.Context, *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	panic("implement me")
}

func (c controllerServer) ValidateVolumeCapabilities(context.Context, *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	panic("implement me")
}

func (c controllerServer) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	panic("implement me")
}

func (c controllerServer) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	panic("implement me")
}

func (c controllerServer) CreateSnapshot(context.Context, *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	panic("implement me")
}

func (c controllerServer) DeleteSnapshot(context.Context, *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	panic("implement me")
}

func (c controllerServer) ListSnapshots(context.Context, *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	panic("implement me")
}

func (c controllerServer) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	panic("implement me")
}
