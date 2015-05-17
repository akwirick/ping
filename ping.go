package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"net/http"
)

type PingClient struct {
	name string
	host string
	port int
	path string
	conn *websocket.Conn
}

var name = flag.String("clientName", "ping", "name for Pong client")
var host = flag.String("host", "localhost", "hostname for Pong server")
var port = flag.Int("port", 8080, "port for Pong server")
var path = flag.String("path", "/ws", "URL path to reach Pong server")

func GenericPingClient(name string, host string, port int, path string) *PingClient {
	pc := &PingClient{name, host, port, path, nil}
	return pc
}

func (pc *PingClient) connect() (conn *websocket.Conn, err error) {
	var hdr http.Header

	dialer := new(websocket.Dialer)
	url := fmt.Sprintf("ws://%s:%d%s", pc.host, pc.port, pc.path)
	glog.Info("Connecting to PongServer at %s", url)

	conn, _, err = dialer.Dial(url, hdr)
	return
}

func (s *PingClient) readLoop() {
	for {
		mt, r, err := s.conn.NextReader()
		_ = mt

		if err != nil {
			log.Printf("NextReader error: %v", err)
			s.conn.Close()
			s.started = false
			break
		}
		glog.Info("%v", mt)

		buf, err := ioutil.ReadAll(r)
		_ = buf
		glog.Info("read %d bytes: %v", len(buf), string(buf))
	}

	if !s.shutting_down {
		log.Printf("Reconnecting...")
		s.connect_and_start_loop()
	}
}

func (pc *PingClient) connectAndListenForever() {
	conn, err = pc.connect()
	if err == nil {

	}

}

func (pc *PingClient) String() {
	fmt.Sprintf("PingClient: %s @ %s:%d%s", pc.name, pc.host, pc.port, pc.path)
}

func main() {
	flag.parse()
	pc := GenericPingClient(name, host, port, path)
	pc.connectAndListenForever()
}
