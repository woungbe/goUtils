package wsSample

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const BaseSpotWSURL = "wss://stream.binance.com:9443/ws/" //사용자 웹소켓 연결 URL

type callBackFunc_onConnect func()
type callBackFunc_onUnConnect func()
type callBackFunc_onMessage func(msg []byte)

type BinanceUserWSObject struct {
	con         *websocket.Conn
	bBreakWrite chan bool
	mListenKey  string //웹소켓 리슨키

	bConnected   bool //연결 상태
	mMenualClose bool //수동 연결 해지 시 true

	mCB_onConnect   callBackFunc_onConnect
	mCB_onUnConnect callBackFunc_onUnConnect
	mCB_onMessage   callBackFunc_onMessage

	// Send pings to peer with this period. Must be less than pongWait.
	mPingPeriod int64 //= (pongWait * 9) / 10
	mwriteWait  int64
}

// Init 객체 초기화
func (ty *BinanceUserWSObject) Init(userListenKey string) error {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()

	ty.mMenualClose = false
	ty.bBreakWrite = make(chan bool)
	ty.mwriteWait = int64(10 * time.Second)
	pongWait := int64(60 * time.Second)
	ty.mPingPeriod = (pongWait * 9) / 10

	ty.mListenKey = userListenKey
	ty.bConnected = false
	return nil
}

// SetCallbackFunc 콜백설정
func (ty *BinanceUserWSObject) SetCallbackFunc(cbOnConnect callBackFunc_onConnect, cbOnUnConnect callBackFunc_onUnConnect, cbOnMessage callBackFunc_onMessage) {
	ty.mCB_onConnect = cbOnConnect
	ty.mCB_onUnConnect = cbOnUnConnect
	ty.mCB_onMessage = cbOnMessage
}
func (ty *BinanceUserWSObject) IsConnect() bool {
	return ty.bConnected
}

// clientConnect 웹소켓연결
func (ty *BinanceUserWSObject) ClientConnect() bool {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()

	strUrl := fmt.Sprintf("%s%s", BaseSpotWSURL, ty.mListenKey)
	r, _ := http.NewRequest("GET", strUrl, nil)
	r.Header.Add("Content-Type", "application/json")
	c, _, err := websocket.DefaultDialer.Dial(strUrl, nil)
	ty.con = c
	if err != nil {
		// pawlog.Error("Error", "msg", err.Error())
		return false
	}
	ty.bConnected = true
	ty.procClient()
	return true
}

// Close 연결해제
func (ty *BinanceUserWSObject) Close() {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()

	ty.mMenualClose = true
	if ty.bBreakWrite != nil {
		ty.bBreakWrite <- true
	}
	if ty.con != nil {
		ty.con.Close()
	}
	ty.bConnected = false
}

func (ty *BinanceUserWSObject) onConnected() {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()
	if ty.mCB_onConnect != nil {
		ty.mCB_onConnect()
	}
}

func (ty *BinanceUserWSObject) onUnconnected() {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()
	//msg := fmt.Sprintf("Binance UnConnected")
	//// pawlog.Info(msg)
	ty.bBreakWrite <- true

	if !ty.mMenualClose && ty.bConnected && ty.mCB_onUnConnect != nil {
		ty.mCB_onUnConnect()
	}
	ty.bConnected = false
}

func (ty *BinanceUserWSObject) procClient() {
	go ty.ReadMessage()
	go ty.WriteMessage()
}

func (ty *BinanceUserWSObject) ReadMessage() {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
		if ty.con != nil {
			ty.con.Close()
			ty.onUnconnected()
		}
	}()

	ty.onConnected()

	ty.con.SetReadLimit(81920) //최대 읽기 버퍼 사이즈
	for {
		if ty.con != nil {
			_, message, err := ty.con.ReadMessage()
			if err != nil {
				return
			}
			ty.messagePasering(message)
		}
	}
}

func (ty *BinanceUserWSObject) WriteMessage() {
	ticker := time.NewTicker(time.Duration(ty.mPingPeriod))
	defer func() {
		ticker.Stop()
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()
	for {
		select {
		case <-ty.bBreakWrite:
			return
		case <-ticker.C:
			ty.con.SetWriteDeadline(time.Now().Add(time.Duration(ty.mwriteWait)))
			if err := ty.con.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// messagePasering 메시지 파싱
func (ty *BinanceUserWSObject) messagePasering(msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			// pawlog.Error("Crit Panic", "Error", err)
		}
	}()
	if ty.mCB_onMessage != nil {
		ty.mCB_onMessage(msg)
	}
}
