package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"log"
)

// OtpParameters 구조체
type OtpParameters struct {
	Secret []byte
	Name   string
	Issuer string
}

// MigrationPayload 구조체
type MigrationPayload struct {
	OtpParameters []OtpParameters
}

func main() {
	// otpauth-migration 데이터
	// uri := "otpauth-migration://offline?data=Ck8KKLr76ffcOEDfijzH8%2BMSHCi1VQ9l75G%2BOHLSajy0Auhr63bRS5sIBTkSCGF3cyByb290GhNBbWF6b24gV2ViIFNlcnZpY2VzIAEoATACEAEYASAAKOvOzrH7%2F%2F%2F%2F%2FwE%3D" // 실제 URI 입력
	// data := getDataFromURI(uri)

	// Base64 디코딩
	decodedData, err := base64.URLEncoding.DecodeString("Ck8KKLr76ffcOEDfijzH8%2BMSHCi1VQ9l75G%2BOHLSajy0Auhr63bRS5sIBTkSCGF3cyByb290GhNBbWF6b24gV2ViIFNlcnZpY2VzIAEoATACEAEYASAAKOvOzrH7%2F%2F%2F%2F%2FwE%3D")
	if err != nil {
		log.Fatalf("Base64 디코딩 실패: %v", err)
	}

	// 데이터 파싱
	payload := parseMigrationPayload(decodedData)

	// 결과 출력
	for _, otp := range payload.OtpParameters {
		secret := base32.StdEncoding.EncodeToString(otp.Secret)
		fmt.Printf("계정 이름: %s\n", otp.Name)
		fmt.Printf("발행자: %s\n", otp.Issuer)
		fmt.Printf("비밀키 (Base32): %s\n", secret)
	}
}

// getDataFromURI는 URI에서 data 값을 추출합니다.
func getDataFromURI(uri string) string {
	start := len("otpauth-migration://offline?data=")
	return uri[start:]
}

// parseMigrationPayload는 바이너리 데이터를 파싱하여 MigrationPayload를 반환합니다.
func parseMigrationPayload(data []byte) MigrationPayload {
	var payload MigrationPayload
	offset := 0

	// OtpParameters 반복 읽기
	for offset < len(data) {
		tag := data[offset]
		offset++

		if tag == 10 { // OtpParameters 시작
			length := int(data[offset])
			offset++

			end := offset + length
			otp := parseOtpParameters(data[offset:end])
			payload.OtpParameters = append(payload.OtpParameters, otp)

			offset = end
		}
	}
	return payload
}

// parseOtpParameters는 단일 OtpParameters를 파싱합니다.
func parseOtpParameters(data []byte) OtpParameters {
	var otp OtpParameters
	offset := 0

	for offset < len(data) {
		tag := data[offset]
		offset++

		field := tag >> 3
		switch field {
		case 1: // Secret
			length := int(data[offset])
			offset++
			otp.Secret = data[offset : offset+length]
			offset += length
		case 2: // Name
			length := int(data[offset])
			offset++
			otp.Name = string(data[offset : offset+length])
			offset += length
		case 3: // Issuer
			length := int(data[offset])
			offset++
			otp.Issuer = string(data[offset : offset+length])
			offset += length
		default:
			// 알 수 없는 필드는 건너뜀
			length := int(data[offset])
			offset += length + 1
		}
	}
	return otp
}
