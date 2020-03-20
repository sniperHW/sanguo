package cluster

import (
	"errors"
)

var (
	ERR_STARTED              = errors.New("already started")
	ERR_NO_AVAILABLE_SERVICE = errors.New("no available service")
	ERR_INVAILD_ENDPOINT     = errors.New("invaild endpoint")
	ERR_DUP_CONN             = errors.New("duplicate endPoint connection")
	ERR_RPC_NO_CONN          = errors.New("rpc no connection")
	ERR_DIAL                 = errors.New("dial failed")
)
