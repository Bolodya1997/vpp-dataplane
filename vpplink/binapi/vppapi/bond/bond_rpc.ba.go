// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package bond

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/vpe"
)

// RPCService defines RPC service bond.
type RPCService interface {
	BondAddMember(ctx context.Context, in *BondAddMember) (*BondAddMemberReply, error)
	BondCreate(ctx context.Context, in *BondCreate) (*BondCreateReply, error)
	BondCreate2(ctx context.Context, in *BondCreate2) (*BondCreate2Reply, error)
	BondDelete(ctx context.Context, in *BondDelete) (*BondDeleteReply, error)
	BondDetachMember(ctx context.Context, in *BondDetachMember) (*BondDetachMemberReply, error)
	BondDetachSlave(ctx context.Context, in *BondDetachSlave) (*BondDetachSlaveReply, error)
	BondEnslave(ctx context.Context, in *BondEnslave) (*BondEnslaveReply, error)
	SwBondInterfaceDump(ctx context.Context, in *SwBondInterfaceDump) (RPCService_SwBondInterfaceDumpClient, error)
	SwInterfaceBondDump(ctx context.Context, in *SwInterfaceBondDump) (RPCService_SwInterfaceBondDumpClient, error)
	SwInterfaceSetBondWeight(ctx context.Context, in *SwInterfaceSetBondWeight) (*SwInterfaceSetBondWeightReply, error)
	SwInterfaceSlaveDump(ctx context.Context, in *SwInterfaceSlaveDump) (RPCService_SwInterfaceSlaveDumpClient, error)
	SwMemberInterfaceDump(ctx context.Context, in *SwMemberInterfaceDump) (RPCService_SwMemberInterfaceDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) BondAddMember(ctx context.Context, in *BondAddMember) (*BondAddMemberReply, error) {
	out := new(BondAddMemberReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondCreate(ctx context.Context, in *BondCreate) (*BondCreateReply, error) {
	out := new(BondCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondCreate2(ctx context.Context, in *BondCreate2) (*BondCreate2Reply, error) {
	out := new(BondCreate2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondDelete(ctx context.Context, in *BondDelete) (*BondDeleteReply, error) {
	out := new(BondDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondDetachMember(ctx context.Context, in *BondDetachMember) (*BondDetachMemberReply, error) {
	out := new(BondDetachMemberReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondDetachSlave(ctx context.Context, in *BondDetachSlave) (*BondDetachSlaveReply, error) {
	out := new(BondDetachSlaveReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BondEnslave(ctx context.Context, in *BondEnslave) (*BondEnslaveReply, error) {
	out := new(BondEnslaveReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwBondInterfaceDump(ctx context.Context, in *SwBondInterfaceDump) (RPCService_SwBondInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwBondInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwBondInterfaceDumpClient interface {
	Recv() (*SwBondInterfaceDetails, error)
	api.Stream
}

type serviceClient_SwBondInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwBondInterfaceDumpClient) Recv() (*SwBondInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwBondInterfaceDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceBondDump(ctx context.Context, in *SwInterfaceBondDump) (RPCService_SwInterfaceBondDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceBondDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceBondDumpClient interface {
	Recv() (*SwInterfaceBondDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceBondDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceBondDumpClient) Recv() (*SwInterfaceBondDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceBondDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceSetBondWeight(ctx context.Context, in *SwInterfaceSetBondWeight) (*SwInterfaceSetBondWeightReply, error) {
	out := new(SwInterfaceSetBondWeightReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSlaveDump(ctx context.Context, in *SwInterfaceSlaveDump) (RPCService_SwInterfaceSlaveDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceSlaveDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceSlaveDumpClient interface {
	Recv() (*SwInterfaceSlaveDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceSlaveDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceSlaveDumpClient) Recv() (*SwInterfaceSlaveDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceSlaveDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwMemberInterfaceDump(ctx context.Context, in *SwMemberInterfaceDump) (RPCService_SwMemberInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwMemberInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwMemberInterfaceDumpClient interface {
	Recv() (*SwMemberInterfaceDetails, error)
	api.Stream
}

type serviceClient_SwMemberInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwMemberInterfaceDumpClient) Recv() (*SwMemberInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwMemberInterfaceDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
