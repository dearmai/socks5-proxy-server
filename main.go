package main

import (
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	// SOCKS5 서버를 설정합니다 (인증 없음)
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// TCP로 1080 포트를 열고 듣습니다
	listener, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		log.Fatalf("Failed to listen on 127.0.0.1:1080: %v", err)
	}

	// 서버 시작
	log.Println("SOCKS5 server listening on 127.0.0.1:1080")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
