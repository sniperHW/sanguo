package cluster

import (
	"github.com/sniperHW/kendynet/event"
	"github.com/sniperHW/kendynet/golog"
	"github.com/sniperHW/kendynet/rpc"
	//listener "github.com/sniperHW/kendynet/socket/listener/tcp"
	center_client "github.com/sniperHW/sanguo/center/client"
	"github.com/sniperHW/sanguo/cluster/addr"
	"github.com/sniperHW/sanguo/cluster/priority"
	"github.com/sniperHW/sanguo/cluster/rpcerr"
	"net"
	"sync"
)

const EXPORT bool = true
const dialTerminateCount int = 5
const harbarType uint32 = 255

var logger golog.LoggerI

type serviceManager struct {
	sync.RWMutex
	cluster            *Cluster
	idEndPointMap      map[addr.LogicAddr]*endPoint
	ttEndPointMap      map[uint32]*typeEndPointMap
	ttForignServiceMap map[uint32]*typeForignServiceMap
	harborsByGroup     map[uint32][]*endPoint
}

//确保同一逻辑地址只能被唯一进程使用
type UniLocker interface {
	Lock(addr.Addr) bool
	Unlock()
}

type Cluster struct {
	serverState            clusterState
	queue                  *event.EventQueue
	rpcMgr                 rpcManager
	serviceMgr             serviceManager
	msgMgr                 msgManager
	centerClient           *center_client.CenterClient
	l                      net.Listener //*listener.Listener
	uniLocker              UniLocker
	pendingRPCRequestCount int32
}

func NewCluster() *Cluster {
	c := &Cluster{
		queue: event.NewEventQueueWithPriority(priority.HIGH+1, 10000),
		rpcMgr: rpcManager{
			server: rpc.NewRPCServer(&decoder{}, &encoder{}),
			client: rpc.NewClient(&decoder{}, &encoder{}),
		},
		msgMgr: msgManager{
			msgHandlers: map[uint16]MsgHandler{},
		},
		serviceMgr: serviceManager{
			idEndPointMap:      map[addr.LogicAddr]*endPoint{},
			ttEndPointMap:      map[uint32]*typeEndPointMap{},
			ttForignServiceMap: map[uint32]*typeForignServiceMap{},
			harborsByGroup:     map[uint32][]*endPoint{},
		},
	}
	c.serviceMgr.cluster = c
	c.rpcMgr.cluster = c
	c.rpcMgr.server.SetErrorCodeOnMissingMethod(rpcerr.Err_RPC_InvaildMethod)
	return c
}

var defaultCluster *Cluster = NewCluster()
