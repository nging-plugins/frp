//go:build frpRPC
// +build frpRPC

package model

import (
	"context"
	"fmt"

	rpc "github.com/coscms/rpcxx"
)

func (f *FrpClient) CallRPC(ctx context.Context, serviceMethod string, args interface{}, reply interface{}) error {
	if f.AdminPort > 0 {
		address := fmt.Sprintf(`%s:%d`, f.AdminAddr, f.AdminPort)
		rpcClient := rpc.NewClient(address, f.AdminPwd, nil)
		if args == nil {
			args = &rpc.Empty{}
		}
		if reply == nil {
			reply = &rpc.Empty{}
		}
		return rpcClient.Call(ctx, `ClientRPCService`, serviceMethod, args, reply)
	}
	return rpc.ErrRPCServerDisabled
}
