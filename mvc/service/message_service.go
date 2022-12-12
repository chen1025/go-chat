package service

import (
	"context"
	"ginchat/common"
	"ginchat/pkg/global/log"
	"ginchat/server"
	"ginchat/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

type messageService struct {
}

var MessageService = new(messageService)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ms *messageService) Chat(writer http.ResponseWriter, request *http.Request) {

	// 1.获取参数
	query := request.URL.Query()
	userId := query.Get("userId")
	idInt, _ := strconv.Atoi(userId)
	uid := int64(idInt)

	// 校验token 暂时忽略
	//token := query.Get("token")
	conn, err := upGrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Logger.Error("Chat Upgrade ", log.Any("websocket err", err))
		return
	}
	// 2.获取连接
	currentTime := uint64(time.Now().Unix())
	client := &server.Client{
		Uid:           uid,
		Conn:          conn,
		DataQueue:     make(chan []byte, 50),
		Addr:          conn.RemoteAddr().String(),
		HeartbeatTime: currentTime,
	}
	// 3.开启服务
	server.MyServer.Register <- client
	go client.Read()
	go client.Write()
}

// 查询用户redis消息
func (ms *messageService) GetMsgByRedis(userId, targetId int, start, end int, isRev bool) []string {
	userIdStr := strconv.Itoa(userId)
	targetIdStr := strconv.Itoa(targetId)
	var key string
	if userId > targetId {
		key = common.MESSAGE + userIdStr + "_" + targetIdStr
	} else {
		key = common.MESSAGE + targetIdStr + "_" + userIdStr
	}
	var result []string
	var err error
	if isRev {
		result, err = utils.RedisClient.ZRange(context.Background(), key, int64(start), int64(end)).Result()
	} else {
		result, err = utils.RedisClient.ZRevRange(context.Background(), key, int64(start), int64(end)).Result()
	}
	if err != nil {
		log.Logger.Error("GetMsgByRedis", log.Any("redis err", err))
		return nil
	}
	return result

}
