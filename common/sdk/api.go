package sdk

import (
	"encoding/json"
	"net"
	"sync"
	"time"

	"github.com/byrantz/plato/common/idl/message"
	"github.com/byrantz/plato/common/tcp"
	"github.com/golang/protobuf/proto"
)

const (
	MsgTypeText      = "text"
	MsgTypeAck       = "ack"
	MsgTypeReConn    = "reConn"
	MsgTypeHeartbeat = "heartbeat"
	MsgLogin         = "loginMsg"
)

type Chat struct {
	Nick            string
	UserID          string
	SessionID       string
	conn            *connect
	closeChan       chan struct{}
	MsgClientDTable map[string]uint64
	sync.RWMutex
}

type Message struct {
	Type       string
	Name       string
	FormUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(ip net.IP, port int, nick, userID, sessionID string) *Chat {
	chat := &Chat{
		Nick:            nick,
		UserID:          userID,
		SessionID:       sessionID,
		conn:            newConnet(ip, port),
		closeChan:       make(chan struct{}),
		MsgClientDTable: make(map[string]uint64),
	}
	go chat.loop()
	chat.login()
	go chat.heartbeat()
	return chat
}
func (chat *Chat) Send(msg *Message) {
	// chat.conn.send(msg)
	data, _ := json.Marshal(msg)
	upMsg := &message.UPMsg{
		Head: &message.UPMsgHead{
			ClientID: chat.getClientID(msg.Session),
			ConnID:   chat.conn.connID,
		},
		UPMsgBody: data,
	}
	palyload, _ := proto.Marshal(upMsg)
	chat.conn.send(message.CmdType_UP, palyload)
}

// Close close chat
func (chat *Chat) Close() {
	chat.conn.close()
	close(chat.closeChan)
	close(chat.conn.recvChan)
	close(chat.conn.sendChan)
}

func (chat *Chat) ReConn() {
	chat.Lock()
	defer chat.Unlock()
	chat.conn.reConn()
	chat.reConn()
}

// Recv receive message
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}

// loop 主要处理从连接中读取的消息，并根据消息类型分别处理 ACK 和 下行消息。处理后的消息会被发送到 recvChan 中
func (chat *Chat) loop() {
Loop:
	for {
		select {
		case <-chat.closeChan:
			return
		default:
			mc := &message.MsgCmd{}
			// 没有数据的话，会一直阻塞在这里
			data, err := tcp.ReadData(chat.conn.conn)
			if err != nil {
				goto Loop
			}
			err = proto.Unmarshal(data, mc)
			if err != nil {
				panic(err)
			}
			var msg *Message
			switch mc.Type {
			case message.CmdType_ACK:
				msg = handAckMsg(chat.conn, mc.Payload)
			case message.CmdType_Push:
				msg = handPushMsg(chat.conn, mc.Payload)
			}
			/*无论是 ACK 还是下行消息，都会交给 conn.recvChan ， Chat 的消费函数会监视这个 channel，如果是下行消息则显示，是 ACK 不做处理*/
			chat.conn.recvChan <- msg
		}
	}
}

func (chat *Chat) getClientID(sessionID string) uint64 {
	chat.Lock()
	defer chat.Unlock()
	var res uint64
	if id, ok := chat.MsgClientDTable[sessionID]; ok {
		res = id
	}
	res++
	chat.MsgClientDTable[sessionID] = res
	return res
}

func (chat *Chat) login() {
	loginMsg := message.LoginMsg{
		Head: &message.LoginMsgHead{
			DeviceID: 123,
		},
	}
	palyload, err := proto.Marshal(&loginMsg)
	if err != nil {
		panic(err)
	}
	chat.conn.send(message.CmdType_Login, palyload)
}

func (chat *Chat) reConn() {
	reConn := message.ReConnMsg{
		Head: &message.ReConnMsgHead{
			ConnID: chat.conn.connID,
		},
	}
	palyload, err := proto.Marshal(&reConn)
	if err != nil {
		panic(err)
	}
	chat.conn.send(message.CmdType_ReConn, palyload)
}

func (chat *Chat) heartbeat() {
	// 每隔一秒钟发送一次心跳
	tc := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-chat.closeChan:
			return
		case <-tc.C:
			hearbeat := message.HeartbeatMsg{
				Head: &message.HeartbeatMsgHead{},
			}
			palyload, err := proto.Marshal(&hearbeat)
			if err != nil {
				panic(err)
			}
			chat.conn.send(message.CmdType_Heartbeat, palyload)
		}
	}
}
