package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // 설정 파일 이름 (config.json, config.yaml 등)
	viper.AddConfigPath(".")      // 설정 파일 경로
	viper.SetConfigType("json")   // 여기서 설정 파일 형식을 지정

	err := viper.ReadInConfig() // 설정 파일 읽기
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// 설정 사용
	name := viper.GetString("name") // 'name' 키로 설정된 값을 읽음

	viper.SetEnvPrefix("general.prefix.") // -- 사용법 모름
	// { //
	viper.SetDefault("port", 8080)  // 디폴트 주는거고,
	port := viper.GetString("port") // 이렇게 호출함... 보통 세트로 됨 !!
	// } //

	fmt.Printf("Hello, %s! %s\n", name, port)

	HandlerSetEnvPrefix()
}

func HandlerSetEnvPrefix() {
	viper.SetEnvPrefix("myapp") // 아 Env 가져오는 거구나 ...

	// Viper에 환경 변수 사용 설정
	viper.AutomaticEnv() // 여기서 Env 적용하는거고 ,,

	// 특정 환경 변수 설정 (예제 실행을 위해 임시로 설정)
	os.Setenv("MYAPP_PORT", "8080")
	os.Setenv("MYAPP_HOST", "localhost")

	// 환경 변수에서 설정값 읽기
	port := viper.GetString("PORT") // 'MYAPP_PORT' 환경 변수에서 값을 가져옴
	host := viper.GetString("HOST") // 'MYAPP_HOST' 환경 변수에서 값을 가져옴

	fmt.Printf("Server is running on %s:%s\n", host, port)
}
