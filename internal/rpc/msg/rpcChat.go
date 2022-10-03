package msg

import (
	"GoIM/pkg/common/config"
	"GoIM/pkg/common/constant"
	"GoIM/pkg/common/log"
	"GoIM/pkg/proto/msg"
	"GoIM/pkg/utils"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

//var (
//	sendMsgSuccessCounter prometheus.Counter
//	sendMsgFailedCounter  prometheus.Counter
//)

type rpcChat struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
	delMsgCh        chan deleteMsg
}

type deleteMsg struct {
	UserID      string
	OpUserID    string
	SeqList     []uint32
	OperationID string
}

func NewRpcChatServer(port int) *rpcChat {
	log.NewPrivateLog(constant.LogFileName)
	rc := rpcChat{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.OpenImMsgName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
	//rc.offlineProducer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2mschatOffline.Addr, config.Config.Kafka.Ws2mschatOffline.Topic)
	rc.delMsgCh = make(chan deleteMsg, 1000)
	return &rc
}

func (rpc *rpcChat) Run() {
	log.Info("", "rpcChat init...")
	listenIP := ""
	if config.Config.ListenIP == "" {
		listenIP = "0.0.0.0"
	} else {
		listenIP = config.Config.ListenIP
	}
	address := listenIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error() + rpc.rpcRegisterName)
	}
	log.Info("", "listen network success, address ", address)

	var grpcOpts []grpc.ServerOption
	srv := grpc.NewServer(grpcOpts...)
	defer srv.GracefulStop()

	rpcRegisterIP := config.Config.RpcRegisterIP
	msg.RegisterMsgServer(srv, rpc)
	if rpcRegisterIP == "" {
		rpcRegisterIP, err = utils.GetLocalIP()
		if err != nil {
			log.Error("", "GetLocalIP failed ", err.Error())
		}
	}
	if err != nil {
		log.Error("", "register rpcChat to etcd failed ", err.Error())
		panic(utils.Wrap(err, "register chat module  rpc to etcd err"))
	}
	go rpc.runCh()
	err = srv.Serve(listener)
	if err != nil {
		log.Error("", "rpc rpcChat failed ", err.Error())
		return
	}
	log.Info("", "rpc rpcChat init success")
}

func (rpc *rpcChat) runCh() {
	log.NewInfo("", "start del msg chan ")
	for {
		select {
		case msg := <-rpc.delMsgCh:
			log.NewInfo(msg.OperationID, utils.GetSelfFuncName(), "delmsgch recv new: ", msg)
		}
	}
}
