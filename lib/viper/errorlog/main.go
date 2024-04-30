package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// 로거 설정
	logFile, err := os.OpenFile("./errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("로그 파일을 열 수 없습니다:", err)
	}
	defer logFile.Close()

	logger := logrus.New()
	logger.Out = logFile
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	// Viper 설정
	viper.SetConfigName("config") // 설정 파일의 이름 설정
	viper.AddConfigPath(".")      // 현재 디렉토리를 설정 파일 경로로 추가
	viper.SetConfigType("json")   // 설정 파일 형식은 json

	// 설정 로드
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("설정 로드 실패: %v\n", err)
	} else {
		logger.Info("설정 로드 성공")
	}

}
