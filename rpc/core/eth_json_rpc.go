package core

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

func (env *Environment) EthQuery(
	_ *rpctypes.Context,
	request []byte,
) (*ctypes.ResultEthQuery, error) {
	resEthQuery, err := env.ProxyAppEthQuery.EthQuerySync(context.TODO(), &abci.RequestEthQuery{
		Request: request,
	})
	if err != nil {
		return nil, err
	}

	return &ctypes.ResultEthQuery{Response: *resEthQuery}, nil
}
