package wssvr

import (
	"dfinForBinance/COEExchange/COE_ConcluerSystem/spotAccountDataCtrl/adcConfig"
	"fmt"
	"pawnet/pawlog"
	"time"

	"github.com/gorilla/websocket"
)

type AccountDataGateWaySession struct {
	index int64
	Conn  *websocket.Conn
	//로그인 되어있나
	BLogin         bool
	BChkLogin      bool //정상적인 로그인 명령이 왔는가
	send           chan []byte
	bBreakWrite    chan bool
	bConnected     bool
	bLoginCloseFlg bool   //다른 커넥션에서 로그인햇을때 =true
	svrIP          string //접근IP
}

func newAccountDataGateWaySession(index int64) *AccountDataGateWaySession {

	newObj := new(AccountDataGateWaySession)
	newObj.initObj(index)
	return newObj
}

func (ty *AccountDataGateWaySession) initObj(index int64) {
	ty.send = make(chan []byte, maxMessageSize)
	ty.bBreakWrite = make(chan bool)
	ty.bConnected = false
	ty.bLoginCloseFlg = false
	ty.index = index
}

func (ty *AccountDataGateWaySession) GetIndex() int64 {
	return ty.index
}

// Clear 클리어
func (ty *AccountDataGateWaySession) Clear() {
	ty.Conn = nil
	ty.BLogin = false
	ty.bConnected = false
	ty.bLoginCloseFlg = false
	ty.svrIP = ""
}

// ReadMessage 소켓에서 메시지 읽기
func (ty *AccountDataGateWaySession) ReadMessage() {

	defer func() {
		if ty != nil {
			if ty.Conn != nil {
				ty.Conn.Close()
				ty.onUnconnected()
			}
		}
		if adcConfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()

	ty.Conn.SetReadLimit(maxMessageSize)
	ty.Conn.SetReadDeadline(time.Now().Add(pongWait))
	ty.Conn.SetPongHandler(func(string) error { ty.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	ty.onConnected()
	for {
		_, message, err := ty.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				if adcConfig.IsDebugmode() == true {
					pawlog.Info("Client Session Close ", "Msg =", err)
				}
			}
			/* else {
				log.Error("InfoProxyWSSession Read Error ", "Error Msg =", err)
			}
			ty.onUnconnected()
			ty.Conn.Close()
			return
			*/
			break
		}
		//- 메시지 처리
		//-- 만약 메시지 사이즈가 MAX 사이즈 이상이면 연결종료 시킴
		if len(message) > maxMessageSize {
			pawlog.Error("Error", "Packet size over ", fmt.Sprintf("%d", len(message)))
			ty.Conn.Close()
		} else {
			ty.messagePasering(message)
		}
	}
}

func (ty *AccountDataGateWaySession) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)
	aliveChk := time.NewTicker(time.Second * 10)
	defer func() {
		//log.Info("쓰기정지")
		ticker.Stop()
		aliveChk.Stop()
		if ty != nil {
			/*
				if ty.Conn != nil {
					ty.Conn.Close()
					ty.onUnconnected()
				}
			*/
		}
		//	if futuMFconfig.IsDebugmode() == false {
		if err := recover(); err != nil {
			pawlog.Error("Crit Panic", "Error", err)
		}
		//	}
	}()
	for {
		select {
		case <-ty.bBreakWrite:
			return
		case message, ok := <-ty.send:
			ty.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				ty.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			ty.Conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			ty.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ty.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		case <-aliveChk.C:
			if ty.BChkLogin == false {
				ty.Conn.Close()
			}
		}
	}
}

func (ty *AccountDataGateWaySession) onConnected() {
	ty.BLogin = false
	ty.bConnected = true
	ty.svrIP = ""
	ty.BChkLogin = true
}
func (ty *AccountDataGateWaySession) onUnconnected() {

	defer func() {
		if err := recover(); err != nil {
			pawlog.Error("Crit Panic", "Error", err)
		}
	}()

	if ty.bConnected == false {
		return
	}
	ty.bConnected = false
	ty.bBreakWrite <- true

	ty.CloseJob()
	GetAccountDataGateWaySvr().ChanLeave <- ty
}

// CloseJob 유저 연결종료시 처리
func (ty *AccountDataGateWaySession) CloseJob() {

	defer func() {
		if err := recover(); err != nil {
			pawlog.Error("Crit Panic", "Error", err)
		}
	}()

	ty.svrIP = ""
	if ty.bLoginCloseFlg == true {
		//같은 아이디 로그인
	}
}

// InfoProxyWSSession 메시지 파싱
func (ty *AccountDataGateWaySession) messagePasering(msg []byte) {
	defer func() {
		if !adcConfig.IsDebugmode() {
			if err := recover(); err != nil {
				pawlog.Error("Crit Panic", "Error", err)
				if ty != nil {
					if ty.Conn != nil {
						ty.Conn.Close()
					}
				}
			}
		}
	}()

	//--> 전송된 데이터를(account Data) ADWS서버로 브로드케스팅 한다.
	GetAccountDataBroadcastSvr().Broadcast(msg)
}

// SendMessage 메시지전송
func (ty *AccountDataGateWaySession) SendMessage(msg string) {

	defer func() {
		if err := recover(); err != nil {
			pawlog.Error("Crit Panic", "Error", err)
			if ty != nil {
				if ty.Conn != nil {
					ty.Conn.Close()
				}
			}
		}
	}()

	ty.SendMessageByte([]byte(msg))
}

// SendMessageByte 메시지전송
func (ty *AccountDataGateWaySession) SendMessageByte(msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			pawlog.Error("Crit Panic", "Error", err)
			if ty != nil {
				if ty.Conn != nil {
					ty.Conn.Close()
				}
			}
		}
	}()

	if ty == nil {
		return
	}
	if ty.Conn == nil {
		return
	}
	ty.send <- msg
}
