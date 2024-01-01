package wssvr

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"pawnet/pawlog"
	"strconv"
	"strings"
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

var upgrader = websocket.Upgrader{ReadBufferSize: 81920, WriteBufferSize: 81920, EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

// Account Data GateWay 서버 ( MTC -> )
type AccountDataGateWaySvr struct {
	MaxClientCount        int64                                //최대 사용자 클라이언트 수
	nTotalConnectionCount int64                                //현재 접속된 전체 카운트
	ClientSessionsMap     map[int64]*AccountDataGateWaySession //클라이언트 맵

	ClientSessionsQ *queue.Queue
	Mmutex          *sync.Mutex
	ChanEnter       chan *websocket.Conn
	ChanLeave       chan *AccountDataGateWaySession
	ChanAddSession  chan int64 //섹션 추가

}

var AccountDataGateWaySvrPTR *AccountDataGateWaySvr

func GetAccountDataGateWaySvr() *AccountDataGateWaySvr {
	if AccountDataGateWaySvrPTR == nil {
		AccountDataGateWaySvrPTR = new(AccountDataGateWaySvr)
	}
	return AccountDataGateWaySvrPTR
}

// Init 초기화
func (ty *AccountDataGateWaySvr) Init() error {

	ty.Mmutex = new(sync.Mutex)

	cnfData := cwhookconfig.GetConfigData()
	nMaxCount := cnfData.ADCSessionMax
	if nMaxCount <= 0 {
		nMaxCount = 100
	}

	ty.MaxClientCount = nMaxCount
	ty.ClientSessionsMap = make(map[int64]*AccountDataGateWaySession) //클라이언트 맵
	//ty.ClientLoginMap = make(map[string]*AccountDataGateWaySession)   //로그인된 사용자 클라이언트 리스트

	ty.ChanEnter = make(chan *websocket.Conn)
	ty.ChanLeave = make(chan *AccountDataGateWaySession)
	ty.ChanAddSession = make(chan int64)

	ty.ClientSessionsQ = queue.New()

	for i := int64(0); i < ty.MaxClientCount; i++ {
		gSession := newAccountDataGateWaySession(i)
		ty.ClientSessionsQ.PushBack(gSession)
	}
	pawlog.Info("Start AccountDataGateWaySvr", "Create Sessions ", strconv.FormatInt(ty.MaxClientCount, 10))

	go ty.sessionCtrlRUN()
	return nil
}

// StartWebSocketServer 시세 웹소켓 서버 시작
func (ty *AccountDataGateWaySvr) StartWebSocketServer() {
	http.HandleFunc("/accGateway", func(w http.ResponseWriter, r *http.Request) {
		ty.sessionWs(w, r)
	})

	pawlog.Info("Status WebSocket Server Starting....")
	svrport := ":5105"
	var addr = flag.String("ADC Svr addr", svrport, "http service address")
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		pawlog.Error("Error", "ListenAndServe: ", err)
		os.Exit(1)
		return
	}

}

func (ty *AccountDataGateWaySvr) sessionWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "Error ", err)
			}
		}
	}()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if c != nil {
			c.Close()
		}
		pawlog.Error("ERROR", "upgrade", err.Error())
		return
	}

	//-- 현재 연결가능한 섹션이 존재하는가부터 체크
	if ty.ClientSessionsQ.Len() <= 0 {
		//-- 더이상 접속 못함
		pawlog.Error("Crit Error", "Ctrl Client Session", "Not empty Session!!!!!!!!!!!!!!")
		c.Close()
		return
	}

	//블럭 IP 처리
	if cwhookconfig.GetConfigData().IPLimitUse {
		strIP := r.RemoteAddr
		if strings.Contains(strIP, ":") == true {
			sar := strings.Split(strIP, ":")
			strIP = sar[0]
		}
		if strIP != "127.0.0.1" {

			limitip := cwhookconfig.GetConfigData().IPLimitAddres
			limitLen := len(limitip)
			if limitLen > len(strIP) {
				pawlog.Warn("Warn", "미등록 서버 IP", strIP)
				c.Close()
				return
			}
			chkIP := strIP[:limitLen]
			if limitip != chkIP {
				pawlog.Warn("Warn", "미등록 서버 IP", strIP)
				c.Close()
				return
			}
		}
	}

	ty.ChanEnter <- c
}

func (ty *AccountDataGateWaySvr) addNewSession(nCount int64) {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "Error ", err)
			}
		}
	}()
	var i int64
	for i = 0; i < nCount; i++ {
		newObj := newAccountDataGateWaySession(i + ty.MaxClientCount)
		//newObj.Index =

		ty.ClientSessionsQ.PushBack(newObj)
	}
	ty.MaxClientCount = ty.MaxClientCount + nCount

	tot := ty.MaxClientCount
	cur := ty.nTotalConnectionCount
	totq := tot - cur

	strP := fmt.Sprintf("\n -------------- User Session Info ----------------\n ")
	strP2 := fmt.Sprintf("Max Clint= %d \n Current User Session Count = %d \n %d Free Sessions", tot, cur, totq)
	strP3 := fmt.Sprintf("%s%s\n-------------------------------------------------", strP, strP2)
	pawlog.Warn(strP3)

}

// sessionCtrlRUN 섹션 런
func (ty *AccountDataGateWaySvr) sessionCtrlRUN() {
	//ticker := time.NewTicker(1 * time.Second)
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "Error ", err)
			}
		}
	}()
	for {
		select {
		case client := <-ty.ChanEnter:
			ty.addSession(client)
		case client := <-ty.ChanLeave:
			ty.closeClientSession(client)
		case nCount := <-ty.ChanAddSession:
			ty.addNewSession(nCount)
		}
	}
}

// CloseClientSession 클라언트 연결종료시
func (ty *AccountDataGateWaySvr) closeClientSession(se *AccountDataGateWaySession) {
	defer func() {

		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "Error ", err)
			}
		}
	}()

	ty.nTotalConnectionCount--
	if ty.nTotalConnectionCount <= 0 {
		ty.nTotalConnectionCount = 0
	}

	delete(ty.ClientSessionsMap, se.GetIndex())
	se.Conn.Close()
	se.Conn = nil
	ty.ClientSessionsQ.PushBack(se)

	if cwhookconfig.IsDebugmode() == true {
		pawlog.Info("ADC status Web Socket UnConnect Clint ", "Index  ", fmt.Sprintf("%d", se.GetIndex()))
		pawlog.Info("ADC status Socket count ", "total", fmt.Sprintf("%d", ty.nTotalConnectionCount))
	}

	//GetLBSClientCtrl().SendStatus(ty.MaxClientCount, ty.nTotalConnectionCount)
}

// AddSession 클라이언트 추가
func (ty *AccountDataGateWaySvr) addSession(Conn *websocket.Conn) {
	var client *AccountDataGateWaySession

	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "WebSocket Add Session fail ", err)
			}
		}
	}()

	ty.nTotalConnectionCount++

	if ty.ClientSessionsQ.Len() <= 0 {
		//-- 더이상 접속 못함
		pawlog.Error("Crit Error", "Client Session", "Not empty Session!!!!!!!!!!!!!!")
		Conn.Close()
		return
	}

	tempObj := ty.ClientSessionsQ.PopFront()
	if tempObj == nil {
		Conn.Close()
		pawlog.Error("Crit Error", "Client Session", "Not empty Session!!!!!!!!!!!!!!")
		return
	} else {
		client = tempObj.(*AccountDataGateWaySession)
		client.Clear()
		client.Conn = Conn
		ty.ClientSessionsMap[client.GetIndex()] = client
		if cwhookconfig.IsDebugmode() {
			pawlog.Info("ADC status Socket Connect Clint ", "Index  ", fmt.Sprintf("%d", client.GetIndex()))
		}
	}

	if client != nil {
		go client.ReadMessage()
		go client.WriteMessage()
	}
}

// GetSessionCnts 현재 접속자 정보
func (ty *AccountDataGateWaySvr) GetSessionCnts() (maxCnt int64, currCnt int64) {
	return ty.MaxClientCount, ty.nTotalConnectionCount
}

// Broadcast 브로캐스트
func (ty *AccountDataGateWaySvr) Broadcast(msg []byte) {
	defer func() {
		if cwhookconfig.IsDebugmode() == false {
			if err := recover(); err != nil {
				pawlog.Error("Crit Error !!!!! Panic", "Error ", err)
			}
		}
	}()

	for _, value := range ty.ClientSessionsMap {
		if value != nil && value.Conn != nil {
			//value.SendAES256Packet(msg)
			value.SendMessageByte(msg)
		}
	}
}
