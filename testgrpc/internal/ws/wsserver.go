package ws

import (
	"context"
	"gopkg.in/olahol/melody.v1"
	"log"
	"net/http"
)

type Server struct {
	svr  *melody.Melody
	mgr  *ClientManager
	msgC chan *msgChan
	ctx  context.Context
}

type msgChan struct {
	kind int
	msg  []byte
}

func New(option *Option) {

	server := &Server{}

	m := melody.New()
	server.svr = m

	server.mgr = new(ClientManager)
	server.msgC = make(chan *msgChan, 1024)
	ctx, done := context.WithCancel(context.Background())
	defer func() {
		recover()
		done()
	}()
	server.ctx = ctx

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		m.HandleRequest(writer, request)
	})

	// 处理websocket客户端新连接，并为每一个新连接创建一个 双向数据流
	m.HandleConnect(func(s *melody.Session) {
		log.Println("有新用户接入")
		server.mgr.clients.Store(s, nil)
	})

	// 处理用户发来的消息
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Println("收到消息:", string(msg))
		// 把用户输入的信息原样返回 websocket 客户端
		//s.Write(msg)
		server.msgC <- &msgChan{
			kind: 1,
			msg:  msg,
		}
	})

	// 处理 websocket 连接断开事件，并关闭session 中 stream的连接
	m.HandleDisconnect(func(s *melody.Session) {
		log.Println("websocket客户端断开连接")
		server.mgr.clients.Delete(s)
	})
	go server.process()
	log.Fatal(http.ListenAndServe(option.Addr, nil))
}

func (s *Server) process() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case cc := <-s.msgC:
			switch cc.kind {
			case 1:
				s.mgr.clients.Range(func(key, value interface{}) bool {
					sess, _ := key.(*melody.Session)
					if sess != nil {
						sess.Write(cc.msg)
					}
					return true
				})
			}
		}
	}
}
