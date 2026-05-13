package v2

import rpchttp "github.com/cometbft/cometbft/rpc/client/http"

// New preserves the legacy v2 import path used by downstream callers while
// reusing the current HTTP/WebSocket client implementation.
func New(remote, wsEndpoint string) (*rpchttp.HTTP, error) {
	return rpchttp.New(remote, wsEndpoint)
}
