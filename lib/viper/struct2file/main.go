package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config 구조체 정의
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig 구조체 정의
type ServerConfig struct {
	Host string
	Port int
}

// DatabaseConfig 구조체 정의
type DatabaseConfig struct {
	User     string
	Password string
	Dbname   string
}

func main() {
	viper.SetConfigName("config") // 설정 파일의 이름
	viper.SetConfigType("json")   // 설정 파일의 타입
	viper.AddConfigPath(".")      // 설정 파일의 위치

	// 설정 파일을 읽습니다.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Config 구조체 인스턴스 생성
	var config Config

	// Viper에서 읽은 설정을 구조체에 매핑합니다.
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	// 구조체를 JSON으로 변환
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling config to JSON, %s", err)
	}

	// JSON 데이터를 파일에 저장
	if err := os.WriteFile("config_make.json", jsonData, 0644); err != nil {
		log.Fatalf("Error writing JSON to file, %s", err)
	}

	fmt.Println("Configuration saved to config.json")
}
