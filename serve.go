package socks5

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func proxyServe() (*net.TCPAddr, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	go func() {
		err := http.Serve(l, &service{})
		if err != nil {
			fmt.Println("error serving", err)
		}
		log.Fatal("serving yo!")
	}()
	// reply is sent before server is started, ensure server is
	// actually running before sending response
	time.Sleep(50 * time.Millisecond)
	addr := l.Addr().(*net.TCPAddr)
	return addr, err
}

type service struct{}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got a request!")
	if r.Body == nil {
		fmt.Println("body is empty")
		return
	}
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("request!", string(bs))
}
