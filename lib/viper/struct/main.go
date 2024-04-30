package main

import (
	"fmt"
	"log"

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

	// 구조체의 내용을 출력하여 확인
	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Database User: %s\n", config.Database.User)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
	fmt.Printf("Database Name: %s\n", config.Database.Dbname)

	fmt.Printf("%+v", config)
}
