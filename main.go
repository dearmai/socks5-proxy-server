package main

import (
	"fmt"
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	// 서버 IP와 포트를 변수로 설정
	ip := "0.0.0.0" // 모든 인터페이스에서 들어오는 연결을 수락
	port := "1080"  // 기본 SOCKS5 포트

	// SOCKS5 서버 설정 (인증 없음)
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// IP와 포트를 이용하여 TCP 리스너 설정
	address := fmt.Sprintf("%s:%s", ip, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	// 서버 시작 정보 로깅
	log.Printf("SOCKS5 server listening on %s\n", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
