package clientconnect

import (
	"encoding/json"
	"fmt"

	// "go/pawnet///pawlog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/phf/go-queue/queue"
	"woungbe.utils/websocket/cwhookconfig"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 81920
)

type ClientConnector struct {
	con           *websocket.Conn
	send          chan []byte
	bBreakWrite   chan bool
	bConnected    bool         //연결 상태
	bLogin        bool         //서버 로그인상태
	bReconnecting bool         //자동 재접속중인가
	packetQueue   *queue.Queue //전송할 패킷 큐
	PacketMmutex  *sync.Mutex
	mErrorMsgFlg  bool   //true이면 같은 에러 전송됐다
	bGoSendQ      bool   //쓰레드큐를 이용한 전송 사용중인가.
	StrUrl        string // url
}

var ClientConnectorPTR *ClientConnector

func GetClientConnector() *ClientConnector {
	if ClientConnectorPTR == nil {
		ClientConnectorPTR = new(ClientConnector)
	}
	return ClientConnectorPTR
}

func (ty *ClientConnector) Init(url string) error {
	ty.send = make(chan []byte, maxMessageSize)
	ty.bBreakWrite = make(chan bool)
	ty.bConnected = false
	ty.bLogin = false
	ty.bReconnecting = false
	ty.PacketMmutex = new(sync.Mutex)
	ty.packetQueue = queue.New()

	ty.mErrorMsgFlg = false
	ty.bGoSendQ = false
	ty.StrUrl = url
	return nil
}

// 서버 연결
func (ty *ClientConnector) SvrConnect() {
	ty.startConnectingServer()
}

// IsConnected 서버와 연결상태 체크
func (ty *ClientConnector) IsConnected() bool {
	if ty.con != nil && ty.bConnected == true {
		return true
	}
	return false
}

// ProcClient R/W 프로세스 시작
func (ty *ClientConnector) procClient() {
	go ty.ReadMessage()
	go ty.WriteMessage()
}

func (ty *ClientConnector) clientConnect() bool {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()

	r, _ := http.NewRequest("GET", ty.StrUrl, nil)
	r.Header.Add("Content-Type", "application/json")
	//pawlog.Info("Connecting", " server Connecting...", strUrl)

	c, _, err := websocket.DefaultDialer.Dial(ty.StrUrl, nil)
	ty.con = c
	if err != nil {
		if !ty.mErrorMsgFlg {
			//pawlog.Error("Error", "msg", err.Error())
			ty.mErrorMsgFlg = true
		}
		return false
	}
	/*
		if bconf.Ssluse == "Y" {
			strUrl = fmt.Sprintf("wss://%s/accGateway", bconf.SvrIP)
			////pawlog.Info(strUrl)
			r, _ := http.NewRequest("GET", strUrl, nil)
			r.Header.Add("Content-Type", "application/json")
			//pawlog.Info("Connecting", " server Connecting...", strUrl)

			//c, _, err := websocket.DefaultDialer.Dial(strUrl, nil)

			roots := x509.NewCertPool()
			severCert, _ := os.ReadFile(bconf.Sslcrt)
			ok := roots.AppendCertsFromPEM(severCert)
			if !ok {
				//pawlog.Error("failed to parse root certificate")
				return false
			}
			d := websocket.Dialer{TLSClientConfig: &tls.Config{RootCAs: roots}}
			c, _, err := d.Dial(strUrl, nil)

			ty.con = c

			if err != nil {
				if !ty.mErrorMsgFlg {
					//pawlog.Error("Error", "msg", err.Error())
					ty.mErrorMsgFlg = true
				}

				return false
			}

		} else {
			strUrl = fmt.Sprintf("ws://%s/accGateway", bconf.SvrIP)

			r, _ := http.NewRequest("GET", strUrl, nil)
			r.Header.Add("Content-Type", "application/json")
			//pawlog.Info("Connecting", " server Connecting...", strUrl)

			c, _, err := websocket.DefaultDialer.Dial(strUrl, nil)
			ty.con = c
			if err != nil {
				if !ty.mErrorMsgFlg {
					//pawlog.Error("Error", "msg", err.Error())
					ty.mErrorMsgFlg = true
				}
				return false
			}
		}
	*/

	ty.bLogin = false
	ty.bReconnecting = false
	ty.bConnected = true

	ty.mErrorMsgFlg = false
	return true
}

func (ty *ClientConnector) onConnected() {
	//pawlog.Info(" server Connected ")
	ty.mErrorMsgFlg = false
	ty.bLogin = true
}
func (ty *ClientConnector) onUnconnected() {
	//pawlog.Error(" server UnConnected ")
	if ty.bConnected == true {
		ty.bBreakWrite <- true
		ty.bConnected = false
		if ty.bReconnecting == false {
			go ty.reconnectingServer()
		}
	}
}

// reconnectingServer 서버 재젒고
func (ty *ClientConnector) reconnectingServer() {
	ty.bReconnecting = true
	for {
		time.Sleep(time.Duration(time.Second * 4))
		//pawlog.Info("Try Reconnect ")
		if ty.bConnected == true {
			return
		}

		if ty.clientConnect() == true {
			ty.procClient()
			return
		}
	}
}

// 최초 접속시 사용
func (ty *ClientConnector) startConnectingServer() {
	ty.bReconnecting = true
	if ty.bConnected {
		return
	}
	if ty.clientConnect() {
		ty.procClient()
		return
	}

	go func() {
		for {
			if ty.bConnected {
				return
			}
			if ty.clientConnect() {
				ty.procClient()
				return
			}
			time.Sleep(time.Duration(time.Second * 2))
		}
	}()

}

func (ty *ClientConnector) ReadMessage() {
	defer func() {
		ty.con.Close()
		ty.onUnconnected()

		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()

	ty.onConnected()
	ty.con.SetReadLimit(maxMessageSize)
	for {
		if ty.con != nil {
			_, message, err := ty.con.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					if cwhookconfig.IsDebugmode() == true {
						//pawlog.Error("client close", "Msg =", err)
					}
				}
				break
			}

			if len(message) > maxMessageSize {
				//pawlog.Error("Error", "Packet size over ", fmt.Sprintf("%d", len(message)))
				ty.con.Close()
			} else {
				ty.messagePasering(message)
			}
		}
	}
}

func (ty *ClientConnector) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		if ty != nil {
			if ty.con != nil {
				ty.con.Close()
				ty.onUnconnected()
			}
		}
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Error !!!!! Panic", "WriteMessage ", err)
			}
		}
	}()
	for {
		select {
		case <-ty.bBreakWrite:
			return
		case message, ok := <-ty.send:
			ty.con.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				ty.con.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			ty.con.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			ty.con.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ty.con.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// SendMessageByte 메시지전송
func (ty *ClientConnector) SendMessageByte(msg []byte) {

	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
				if ty != nil {
					if ty.con != nil {
						ty.con.Close()
					}
				}
			}
		}
	}()

	if ty == nil {
		return
	}
	if ty.con == nil {
		return
	}
	ty.send <- msg
}

// 메시지 전송( 쓰레드 큐 이용)
func (ty *ClientConnector) SendMessageForQueue(packetData []byte) {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()

	if !ty.bGoSendQ {
		ty.bGoSendQ = true
		go ty.goFuncSendQueue()
	}

	ty.PacketMmutex.Lock()
	ty.packetQueue.PushBack(packetData)
	ty.PacketMmutex.Unlock()
}
func (ty *ClientConnector) SendDataForQueue(packetData interface{}) error {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()
	jsonData, er := json.Marshal(packetData)
	if er != nil {
		return er
	}
	ty.SendMessageForQueue(jsonData)
	return nil
}

// goFuncSendQuMiddleSvrProc 전송용 쓰레드
func (ty *ClientConnector) goFuncSendQueue() {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
			}
		}
	}()
	for {
		//nc := ty.packetQueue.Len()
		if ty.packetQueue.Len() > 0 {
			for {
				if ty.IsConnected() {
					ty.PacketMmutex.Lock()
					sndData := ty.packetQueue.PopFront() //.([]byte)
					ty.PacketMmutex.Unlock()
					if sndData != nil {
						ty.SendMessageByte(sndData.([]byte))
					} else {
						break
					}
				} else {
					break
				}
			}
		}
		time.Sleep(time.Millisecond)
	}
}

// messagePasering 메시지 파싱
func (ty *ClientConnector) messagePasering(msg []byte) {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				//pawlog.Error("Crit Panic", "Error", err)
				if ty != nil {
					if ty.con != nil {
						ty.con.Close()
					}
				}
			}
		}
	}()

	var packetdata map[string]interface{}
	jer := json.Unmarshal(msg, &packetdata)
	if jer != nil {
		ty.con.Close()
		return
	}
	if val, ok := packetdata["p"]; ok {
		price := val.(string)
		fmt.Println("message : ", price)
	}

}
