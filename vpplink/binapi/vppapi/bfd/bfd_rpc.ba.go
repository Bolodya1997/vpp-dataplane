// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package bfd

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/vpe"
)

// RPCService defines RPC service bfd.
type RPCService interface {
	BfdAuthDelKey(ctx context.Context, in *BfdAuthDelKey) (*BfdAuthDelKeyReply, error)
	BfdAuthKeysDump(ctx context.Context, in *BfdAuthKeysDump) (RPCService_BfdAuthKeysDumpClient, error)
	BfdAuthSetKey(ctx context.Context, in *BfdAuthSetKey) (*BfdAuthSetKeyReply, error)
	BfdUDPAdd(ctx context.Context, in *BfdUDPAdd) (*BfdUDPAddReply, error)
	BfdUDPAuthActivate(ctx context.Context, in *BfdUDPAuthActivate) (*BfdUDPAuthActivateReply, error)
	BfdUDPAuthDeactivate(ctx context.Context, in *BfdUDPAuthDeactivate) (*BfdUDPAuthDeactivateReply, error)
	BfdUDPDel(ctx context.Context, in *BfdUDPDel) (*BfdUDPDelReply, error)
	BfdUDPDelEchoSource(ctx context.Context, in *BfdUDPDelEchoSource) (*BfdUDPDelEchoSourceReply, error)
	BfdUDPGetEchoSource(ctx context.Context, in *BfdUDPGetEchoSource) (*BfdUDPGetEchoSourceReply, error)
	BfdUDPMod(ctx context.Context, in *BfdUDPMod) (*BfdUDPModReply, error)
	BfdUDPSessionDump(ctx context.Context, in *BfdUDPSessionDump) (RPCService_BfdUDPSessionDumpClient, error)
	BfdUDPSessionSetFlags(ctx context.Context, in *BfdUDPSessionSetFlags) (*BfdUDPSessionSetFlagsReply, error)
	BfdUDPSetEchoSource(ctx context.Context, in *BfdUDPSetEchoSource) (*BfdUDPSetEchoSourceReply, error)
	WantBfdEvents(ctx context.Context, in *WantBfdEvents) (*WantBfdEventsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) BfdAuthDelKey(ctx context.Context, in *BfdAuthDelKey) (*BfdAuthDelKeyReply, error) {
	out := new(BfdAuthDelKeyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdAuthKeysDump(ctx context.Context, in *BfdAuthKeysDump) (RPCService_BfdAuthKeysDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BfdAuthKeysDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BfdAuthKeysDumpClient interface {
	Recv() (*BfdAuthKeysDetails, error)
	api.Stream
}

type serviceClient_BfdAuthKeysDumpClient struct {
	api.Stream
}

func (c *serviceClient_BfdAuthKeysDumpClient) Recv() (*BfdAuthKeysDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BfdAuthKeysDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BfdAuthSetKey(ctx context.Context, in *BfdAuthSetKey) (*BfdAuthSetKeyReply, error) {
	out := new(BfdAuthSetKeyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPAdd(ctx context.Context, in *BfdUDPAdd) (*BfdUDPAddReply, error) {
	out := new(BfdUDPAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPAuthActivate(ctx context.Context, in *BfdUDPAuthActivate) (*BfdUDPAuthActivateReply, error) {
	out := new(BfdUDPAuthActivateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPAuthDeactivate(ctx context.Context, in *BfdUDPAuthDeactivate) (*BfdUDPAuthDeactivateReply, error) {
	out := new(BfdUDPAuthDeactivateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPDel(ctx context.Context, in *BfdUDPDel) (*BfdUDPDelReply, error) {
	out := new(BfdUDPDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPDelEchoSource(ctx context.Context, in *BfdUDPDelEchoSource) (*BfdUDPDelEchoSourceReply, error) {
	out := new(BfdUDPDelEchoSourceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPGetEchoSource(ctx context.Context, in *BfdUDPGetEchoSource) (*BfdUDPGetEchoSourceReply, error) {
	out := new(BfdUDPGetEchoSourceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPMod(ctx context.Context, in *BfdUDPMod) (*BfdUDPModReply, error) {
	out := new(BfdUDPModReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPSessionDump(ctx context.Context, in *BfdUDPSessionDump) (RPCService_BfdUDPSessionDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BfdUDPSessionDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BfdUDPSessionDumpClient interface {
	Recv() (*BfdUDPSessionDetails, error)
	api.Stream
}

type serviceClient_BfdUDPSessionDumpClient struct {
	api.Stream
}

func (c *serviceClient_BfdUDPSessionDumpClient) Recv() (*BfdUDPSessionDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BfdUDPSessionDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BfdUDPSessionSetFlags(ctx context.Context, in *BfdUDPSessionSetFlags) (*BfdUDPSessionSetFlagsReply, error) {
	out := new(BfdUDPSessionSetFlagsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BfdUDPSetEchoSource(ctx context.Context, in *BfdUDPSetEchoSource) (*BfdUDPSetEchoSourceReply, error) {
	out := new(BfdUDPSetEchoSourceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantBfdEvents(ctx context.Context, in *WantBfdEvents) (*WantBfdEventsReply, error) {
	out := new(WantBfdEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
