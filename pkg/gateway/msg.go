package gateway

import (
	"GoIM/pkg/common/constant"
	"GoIM/pkg/common/db"
	"GoIM/pkg/common/host"
	pbRelay "GoIM/pkg/proto/relay"
	sdk_ws "GoIM/pkg/proto/sdk_ws"
	"GoIM/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type ClientMsg struct {
	msgType    string
	msgTo      string
	msgContent string
}

func NewClientMsg(msgType string, to string, content string) *ClientMsg {
	return &ClientMsg{
		msgType:    msgType,
		msgTo:      to,
		msgContent: content,
	}
}

func SendMsg(msg *ClientMsg) error {
	var uHost host.UserConnectedHost
	msgToUid := msg.msgTo
	hostKey := constant.UserWsConnectedHostPrefix + msgToUid
	fmt.Println("查询用户连接的服务器", hostKey)
	uHostJson, err := db.DB.RDB.Get(context.Background(), hostKey).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(uHostJson), &uHost)
	if err != nil {
		panic(err.Error())
	}
	// rpc调用那台服务器的SendMsg方法，往它的websocket连接中写入消息
	_, err = sendMsgByRpc(&uHost, msgToUid, msg.msgContent)
	if err != nil {
		return err
	}
	return nil
}

func sendMsgByRpc(uHost *host.UserConnectedHost, toUid string, content string) (*pbRelay.OnlinePushMsgResp, error) {
	fmt.Println("grpc dial:", uHost)
	conn, err := grpc.Dial(uHost.IP+":"+utils.Int32ToString(uHost.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbRelay.NewRelayClient(conn)
	msgData := &sdk_ws.MsgData{
		RecvID:  toUid,
		Content: []byte(content),
	}
	pbReq := &pbRelay.OnlinePushMsgReq{
		PushToUserID: toUid,
		MsgData:      msgData,
	}
	return client.OnlinePushMsg(context.TODO(), pbReq)
}
