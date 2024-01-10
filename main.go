package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	// 기본값을 가진 포트 플래그를 정의
	var port string
	flag.StringVar(&port, "port", "1080", "The port number for the SOCKS5 server to listen on")
	flag.Parse() // 플래그를 파싱합니다.

	ip := "0.0.0.0" // 모든 인터페이스에서 들어오는 연결을 수락

	// SOCKS5 서버 설정 (username, password 인증)
	conf := &socks5.Config{}

	//cred := socks5.StaticCredentials{
	//	"foo": "bar",
	//}
	//cator := socks5.UserPassAuthenticator{Credentials: cred}
	//conf.AuthMethods = []socks5.Authenticator{cator}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// IP와 파싱된 포트를 이용하여 TCP 리스너 설정
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
