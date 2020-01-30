package dropbox

import (
	"os"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/utils/mount"
)

type nodeServer struct {
	nodeID string
}

func NewNodeServer(nodeId string) *nodeServer {
	return &nodeServer{
		nodeID: nodeId,
	}
}

const (
	volumePath string = "/csi-dropbox-data"
)

func (n *nodeServer) NodeGetInfo(context.Context, *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	return &csi.NodeGetInfoResponse{
		NodeId: n.nodeID,
	}, nil
}

func (n nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	glog.Infof("volumeId: %v", req.GetVolumeId())
	glog.Infof("targetPath: %v", req.GetStagingTargetPath())
	glog.Infof("cap: %v", req.GetVolumeCapability())

	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetStagingTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}
	if req.GetVolumeCapability() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume Capability missing in request")
	}

	err := os.MkdirAll(volumePath, 0777)
	if err != nil {
		return nil, err
	}

	// TODO: dropbox 마운트하기

	return &csi.NodeStageVolumeResponse{}, nil
}

func (n nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetStagingTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (n nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if req.GetVolumeCapability() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capability missing in request")
	}
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}

	targetPath := req.GetTargetPath()

	notMnt, err := mount.New("").IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		return &csi.NodePublishVolumeResponse{}, nil
	}

	options := []string{"bind"}
	if req.GetReadonly() {
		options = append(options, "ro")
	}

	mounter := mount.New("")
	if err := mounter.Mount(volumePath, targetPath, "", options); err != nil {
		var errList strings.Builder
		errList.WriteString(err.Error())
	}
	glog.V(4).Infof("dropbox-csi: volume %s is mount to %s.", volumePath, targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (n nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if len(req.GetTargetPath()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Target path missing in request")
	}
	targetPath := req.GetTargetPath()

	err := mount.New("").Unmount(req.GetTargetPath())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	glog.V(4).Infof("dropbox-csi: volume %s is unmounted,", targetPath)

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (n *nodeServer) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (n nodeServer) NodeGetVolumeStats(context.Context, *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	panic("implement me node volumestats")
}

func (n nodeServer) NodeExpandVolume(context.Context, *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	panic("implement me node expand")
}
