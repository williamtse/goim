package server

import (
	"GoIM/pkg/auth"
	"GoIM/pkg/common/config"
	"GoIM/pkg/common/constant"
	"GoIM/pkg/common/log"
	pbMsg "GoIM/pkg/proto/msg"
	pbRelay "GoIM/pkg/proto/relay"
	sdk_ws "GoIM/pkg/proto/sdk_ws"
	"GoIM/pkg/utils"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
)

type RPCServer struct {
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
	platformList    []int
	pushTerminal    []int
	target          string
}
type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	OperationID   string `json:"operationID"`
	ErrCode       int32  `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}

func (r *RPCServer) onInit(rpcPort int) {
	r.rpcPort = rpcPort
	r.rpcRegisterName = config.Config.RpcRegisterName.OpenImRelayName
	r.etcdSchema = config.Config.Etcd.EtcdSchema
	r.etcdAddr = config.Config.Etcd.EtcdAddr
	r.platformList = genPlatformArray()
	r.pushTerminal = []int{constant.IOSPlatformID, constant.AndroidPlatformID}
	fmt.Println("rpc server init:", r)
}
func (r *RPCServer) run() {
	fmt.Println("rpc server run...")
	listenIP := ""
	if config.Config.ListenIP == "" {
		listenIP = "0.0.0.0"
	} else {
		listenIP = config.Config.ListenIP
	}
	address := listenIP + ":" + strconv.Itoa(r.rpcPort)
	fmt.Println("rpc server:", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error() + r.rpcRegisterName)
	}
	defer listener.Close()
	var grpcOpts []grpc.ServerOption

	srv := grpc.NewServer(grpcOpts...)
	defer srv.GracefulStop()
	pbRelay.RegisterRelayServer(srv, r)
	pbMsg.RegisterMsgServer(srv, r)

	// rpcRegisterIP := config.Config.RpcRegisterIP
	// if rpcRegisterIP == "" {
	// 	rpcRegisterIP, err = utils.GetLocalIP()
	// 	if err != nil {
	// 		log.Error("", "GetLocalIP failed ", err.Error())
	// 	}
	// }
	// err = getcdv3.RegisterEtcd4Unique(r.etcdSchema, strings.Join(r.etcdAddr, ","), rpcRegisterIP, r.rpcPort, r.rpcRegisterName, 10)
	// if err != nil {
	// 	log.Error("", "register push message rpc to etcd err", "", "err", err.Error(), r.etcdSchema, strings.Join(r.etcdAddr, ","), rpcRegisterIP, r.rpcPort, r.rpcRegisterName)
	// 	panic(utils.Wrap(err, "register msg_gataway module  rpc to etcd err"))
	// }
	// r.target = getcdv3.GetTarget(r.etcdSchema, rpcRegisterIP, r.rpcPort, r.rpcRegisterName)
	err = srv.Serve(listener)
	if err != nil {
		log.Error("", "push message rpc listening err", "", "err", err.Error())
		return
	}
}

func getUserWsClient(uid string) *Client {
	authUser := auth.GetUserAuthInfoByUid(uid)
	return hub.auths[authUser.AccessToken]
}

func (r *RPCServer) SendMsg(_ context.Context, in *pbMsg.SendMsgReq) (*pbMsg.SendMsgResp, error) {
	return &pbMsg.SendMsgResp{}, nil
}

func (r *RPCServer) OnlinePushMsg(_ context.Context, in *pbRelay.OnlinePushMsgReq) (*pbRelay.OnlinePushMsgResp, error) {
	fmt.Println("rpc relay->OnlinePushMsg", in)
	log.NewInfo(in.OperationID, "PushMsgToUser is arriving", in.String())
	recvID := in.PushToUserID
	client := getUserWsClient(recvID)
	if client != nil {
		client.send <- []byte(in.MsgData.Content)
	} else {
		log.NewDebug(in.OperationID, "push err ,no matched ws conn not in map", in.String())
	}
	return &pbRelay.OnlinePushMsgResp{}, nil
}
func (r *RPCServer) encodeWsData(wsData *sdk_ws.MsgData, operationID string) (bytes.Buffer, error) {
	log.Debug(operationID, "encodeWsData begin", wsData.String())
	msgBytes, err := proto.Marshal(wsData)
	if err != nil {
		log.NewError(operationID, "Marshal", err.Error())
		return bytes.Buffer{}, utils.Wrap(err, "")
	}
	log.Debug(operationID, "encodeWsData begin", wsData.String())
	mReply := Resp{
		ReqIdentifier: constant.WSPushMsg,
		OperationID:   operationID,
		Data:          msgBytes,
	}
	var replyBytes bytes.Buffer
	enc := gob.NewEncoder(&replyBytes)
	err = enc.Encode(mReply)
	if err != nil {
		log.NewError(operationID, "data encode err", err.Error())
		return bytes.Buffer{}, utils.Wrap(err, "")
	}
	return replyBytes, nil
}

func genPlatformArray() (array []int) {
	for i := 1; i <= constant.LinuxPlatformID; i++ {
		array = append(array, i)
	}
	return array
}
