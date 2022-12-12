package server

import (
	"context"
	"encoding/json"
	"ginchat/common"
	"ginchat/models"
	"ginchat/mvc/dao"
	"ginchat/pkg/global/log"
	"ginchat/utils"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

var MyServer = NewServer()

type Server struct {
	Register  chan *Client
	SignOut   chan *Client
	Broadcast chan []byte
	Clients   map[int64]*Client
	mutex     *sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		Register:  make(chan *Client),
		SignOut:   make(chan *Client),
		Broadcast: make(chan []byte),
		Clients:   make(map[int64]*Client),
		mutex:     &sync.RWMutex{},
	}
}

func (s *Server) Start() {
	log.Logger.Info("Server Start...")
	for {
		select {
		case c := <-s.Register:
			userId := strconv.Itoa(int(c.Uid))
			log.Logger.Info("用户注册 uid=" + userId)
			// 注册
			s.mutex.Lock()
			s.Clients[c.Uid] = c
			s.mutex.Unlock()
			// 设置在线状态
			utils.SetUserOnlineInfo(common.ONLINE+userId, []byte(c.Addr), time.Duration(utils.Conf.Timeout.RedisOnlineTime)*time.Hour)
		case c := <-s.SignOut:
			userId := strconv.Itoa(int(c.Uid))
			log.Logger.Info("用户退出 uid=" + userId)
			s.mutex.Lock()
			delete(s.Clients, c.Uid)
			s.mutex.Unlock()
			c.Conn.Close()
			// 删除在线状态
			utils.RedisClient.Del(context.Background(), common.ONLINE+userId)
		case msg := <-s.Broadcast:
			dispatch(msg, s)
		}
	}
}

// 调度
func dispatch(data []byte, server *Server) {
	msg := models.Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Logger.Error("dispatch ", log.Any("json err", err))
		return
	}
	switch msg.Type {
	case 1:
		sendMsg(msg.TargetId, data, server)
	case 2:
		sendGroupMsg(msg.TargetId, data, server)
	case 3:
		sendAllMsg()
	case 4:

	}
}

// 私信
func sendMsg(targetId int64, msg []byte, server *Server) {
	server.mutex.RLock()
	client, ok := server.Clients[targetId]
	server.mutex.RUnlock()
	// 判断是否在线
	jsonMsg := models.Message{}
	json.Unmarshal(msg, &jsonMsg)
	cxt := context.Background()
	fromIdStr := strconv.Itoa(int(jsonMsg.UserId))
	targetIdStr := strconv.Itoa(int(targetId))
	result, err := utils.RedisClient.Get(cxt, common.ONLINE+targetIdStr).Result()
	if err != nil {
		log.Logger.Error("sendMsg redis1", log.Any("redis err", err))
		return
	}
	if result != "" {
		if ok {
			client.DataQueue <- msg
		}
	}
	// 存入redis 中
	// 需要双方都能看到全部消息，把key存成一个通用的
	var key string
	if jsonMsg.UserId > targetId {
		key = common.MESSAGE + fromIdStr + "_" + targetIdStr
	} else {
		key = common.MESSAGE + targetIdStr + "_" + fromIdStr
	}

	re, err := utils.RedisClient.ZRevRange(cxt, key, 0, -1).Result()
	if err != nil {
		log.Logger.Error("sendMsg redis2", log.Any("redis err", err))
		return
	}
	score := float64(cap(re) + 1)
	_, err = utils.RedisClient.ZAdd(cxt, key, &redis.Z{Member: msg, Score: score}).Result()
	if err != nil {
		log.Logger.Error("sendMsg redis3", log.Any("redis err", err))
		return
	}
}

// 群组
func sendGroupMsg(targetId int64, msg []byte, server *Server) {
	// {"targetId":3,"type":2,"createTime":1670140148706,"userId":1,"media":1,"content":"111"}
	// targetId 群id
	userIds := dao.Group.GetGroupUserById(uint(targetId))
	for i := 0; i < len(userIds); i++ {
		sendMsg(int64(userIds[i]), msg, server)
	}
}

// 广播
func sendAllMsg() {

}

// ClearConn 清理过期链接
func ClearConn(params interface{}) bool {
	defer func() {
		var n any = nil
		if r := recover(); r != n {
			log.Logger.Error("ClearConn", log.Any("recover err", r))
		}
	}()
	for _, node := range MyServer.Clients {
		if node.IsTimeout() {
			// 退出
			MyServer.SignOut <- node
		}
	}
	return true
}
