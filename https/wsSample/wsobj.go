package wsSample

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type callBackFunc_onConnect func()
type callBackFunc_onUnConnect func()
type callBackFunc_onMessage func(msg []byte)

type BinanceUserWSObject struct {
	con         *websocket.Conn
	bBreakWrite chan bool
	mListenKey  string // 웹소켓 리슨키

	bConnected   bool // 연결 상태
	mMenualClose bool // 수동 연결 해지 시 true

	send           chan []byte
	mCB_onConnect   callBackFunc_onConnect
	mCB_onUnConnect callBackFunc_onUnConnect
	mCB_onMessage   callBackFunc_onMessage

	// Send pings to peer with this period. Must be less than pongWait.
	mPingPeriod int64 // = (pongWait * 9) / 10
	mwriteWait  int64
	Url         string
}

// Init initializes the WebSocket object
func (ty *BinanceUserWSObject) Init(listenKey string, url string) {
	ty.mListenKey = listenKey
	ty.Url = url
	ty.bBreakWrite = make(chan bool)
	ty.send = make(chan []byte)
	ty.mwriteWait = int64(10 * time.Second)
	pongWait := int64(60 * time.Second)
	ty.mPingPeriod = (pongWait * 9) / 10
	ty.bConnected = false
	ty.mMenualClose = false
}

// SetCallbackFunc sets callback functions
func (ty *BinanceUserWSObject) SetCallbackFunc(onConnect callBackFunc_onConnect, onUnConnect callBackFunc_onUnConnect, onMessage callBackFunc_onMessage) {
	ty.mCB_onConnect = onConnect
	ty.mCB_onUnConnect = onUnConnect
	ty.mCB_onMessage = onMessage
}

// IsConnect returns the connection status
func (ty *BinanceUserWSObject) IsConnect() bool {
	return ty.bConnected
}

// ClientConnect establishes a WebSocket connection
func (ty *BinanceUserWSObject) ClientConnect() bool {
	strUrl := fmt.Sprintf("%s%s", ty.Url, ty.mListenKey)
	r, _ := http.NewRequest("GET", strUrl, nil)
	r.Header.Add("Content-Type", "application/json")
	c, _, err := websocket.DefaultDialer.Dial(strUrl, nil)
	if err != nil {
		log.Printf("WebSocket connection failed: %v", err)
		return false
	}

	ty.con = c
	ty.bConnected = true
	log.Println("WebSocket connected")
	ty.procClient()
	return true
}

// Close terminates the WebSocket connection
func (ty *BinanceUserWSObject) Close() {
	if ty.mMenualClose {
		return
	}

	ty.mMenualClose = true
	if ty.bBreakWrite != nil {
		close(ty.bBreakWrite)
	}
	if ty.con != nil {
		ty.con.Close()
	}
	ty.bConnected = false
	log.Println("WebSocket closed")
}

func (ty *BinanceUserWSObject) onConnected() {
	if ty.mCB_onConnect != nil {
		ty.mCB_onConnect()
	}
}

func (ty *BinanceUserWSObject) onUnconnected() {
	if ty.mCB_onUnConnect != nil {
		ty.mCB_onUnConnect()
	}
	ty.bConnected = false
	log.Println("WebSocket disconnected")
}

func (ty *BinanceUserWSObject) procClient() {
	go ty.readMessage()
	go ty.writeMessage()
}

func (ty *BinanceUserWSObject) readMessage() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("readMessage panic: %v", err)
		}
		if ty.con != nil {
			ty.con.Close()
			ty.onUnconnected()
		}
	}()

	ty.onConnected()
	ty.con.SetReadLimit(81920)
	for {
		_, message, err := ty.con.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			ty.detectDisconnection(err)
			return
		}
		if ty.mCB_onMessage != nil {
			ty.mCB_onMessage(message)
		}
	}
}

func (ty *BinanceUserWSObject) writeMessage() {
	ticker := time.NewTicker(time.Duration(ty.mPingPeriod))
	defer func() {
		ticker.Stop()
		if err := recover(); err != nil {
			log.Printf("writeMessage panic: %v", err)
		}
	}()

	for {
		select {
		case <-ty.bBreakWrite:
			return
		case message, ok := <-ty.send:
			if !ok {
				return
			}
			if err := ty.con.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Write error: %v", err)
				ty.detectDisconnection(err)
				return
			}
		case <-ticker.C:
			if err := ty.con.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Ping error: %v", err)
				ty.detectDisconnection(err)
				return
			}
		}
	}
}

func (ty *BinanceUserWSObject) SendMessage(message string) {
	if ty == nil || ty.con == nil || !ty.bConnected {
		return
	}
	ty.send <- []byte(message)
}

// Reconnect handles WebSocket reconnection logic
func (ty *BinanceUserWSObject) Reconnect() {
	log.Println("Attempting to reconnect...")
	for {
		if ty.ClientConnect() {
			log.Println("Reconnected successfully")
			return
		}
		log.Println("Reconnect failed, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}

// detectDisconnection detects WebSocket disconnection and triggers reconnection
func (ty *BinanceUserWSObject) detectDisconnection(err error) {
	log.Printf("Connection lost: %v", err)
	if ty.bConnected {
		ty.bConnected = false
		go ty.Reconnect()
	}
}
