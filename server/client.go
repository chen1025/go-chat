package server

import (
	"encoding/json"
	"ginchat/models"
	"ginchat/pkg/global/log"
	"ginchat/utils"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Uid           int64
	Conn          *websocket.Conn
	DataQueue     chan []byte
	Addr          string
	HeartbeatTime uint64
}

// 发送消息
func (c *Client) Write() {
	for {
		select {
		case data := <-c.DataQueue:
			//log.Logger.Info("[ws] sendProc >>>" + string(data))
			err := c.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Logger.Error("sendProc ", log.Any("websocket write err", err))
				return
			}
		}
	}
}

// 接受消息
func (c *Client) Read() {
	for {
		// 读取数据
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			// socket 关闭注销
			MyServer.SignOut <- c
			log.Logger.Error("receiveProc ", log.Any("websocket read err", err))
			return
		}
		msg := models.Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Logger.Error("receiveProc ", log.Any("json err", err))
		}
		if msg.Type == 3 {
			// 心跳
			c.Heartbeat()
		} else {
			broadMsg(data)
		}
		log.Logger.Info("[ws] receiveProc <<<< " + string(data))
	}
}

// 广播消息
func broadMsg(data []byte) {
	MyServer.Broadcast <- data
}

func (c *Client) Heartbeat() {
	c.HeartbeatTime = uint64(time.Now().Unix())
}

func (c *Client) IsTimeout() (timeout bool) {
	currentTime := uint64(time.Now().Unix())
	if c.HeartbeatTime+uint64(utils.Conf.Timeout.HeartbeatMaxTime) <= currentTime {
		log.Logger.Info("心跳超时。。。自动下线 node=" + c.Addr)
		timeout = true
	}
	return
}
