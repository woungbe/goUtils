package cwhookconfig

var debugmodeflg bool //디버그 모드 플레그

type APISvrConfigData struct {
	ADCSvrIP       string //ADC 웹소켓 서버 IP (포트 포함  = 127.0.0.1:5106)
	ADWSSessionMax int64  //ADWS 클라이언트 최대 섹션수 ( 기본 5000)

}

type APISvrConfig struct {
	cnf           APISvrConfigData
	configOrgData string //설정 정보 데이터
}

var APISvrConfigPtr *APISvrConfig

// IsDebugmode 현재  디버그 모드이가 체크
func IsDebugmode() bool {
	return debugmodeflg
}

// SetDebugmode 디버그모드 설정
func SetDebugmode(b bool) {
	debugmodeflg = b
}

// GetConfigData 서버 기본정보
func GetConfigData() *APISvrConfigData {
	return &APISvrConfigPtr.cnf
}
