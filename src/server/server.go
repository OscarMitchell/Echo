package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	FIFTY_MS       = time.Millisecond * 50
	ONE_HUNDRED_MS = time.Millisecond * 100
)

type Server struct {
	ctx            context.Context
	server         net.Listener
	connections    map[net.Conn]bool
	connectionsMtx sync.Mutex
}

func NewServer(ctx context.Context, port string) (*Server, error) {
	server, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("!!!!! Error creating a new server: %w", err)
	}
	return &Server{
		ctx:         ctx,
		server:      server,
		connections: make(map[net.Conn]bool),
	}, nil
}

func (s *Server) Run() {
	errCount := 0
	for {
		conn, err := s.server.Accept()
		if errors.Is(err, net.ErrClosed) {
			log.Println("Server closed")
			break
		}
		if err != nil {
			log.Printf("!!! Error when attempting to accept incoming connection to server: %v", err)
			errCount++
			if errCount > 50 {
				log.Println("!!!!! Too many errors closing socket")
				s.Shutdown()
			}

			time.Sleep(ONE_HUNDRED_MS)
			continue
		}
		errCount = 0

		go s.handleIncomingConnection(conn)
	}
}

func (s *Server) Shutdown() {
	log.Println("Shutdown triggered, closing connections and server")
	s.connectionsMtx.Lock()
	for conn := range s.connections {
		_ = conn.Close()
	}
	s.connectionsMtx.Unlock()
	_ = s.server.Close()
}

func (s *Server) handleIncomingConnection(conn net.Conn) {
	newConnMsg := fmt.Sprintln("New connection from:", conn.RemoteAddr().String())
	runtime.LogInfo(s.ctx, newConnMsg)
	runtime.EventsEmit(s.ctx, "message-rx", newConnMsg)

	s.connectionsMtx.Lock()
	s.connections[conn] = true
	s.connectionsMtx.Unlock()

	buffer := bufio.NewReader(conn)
	for {
		str, err := buffer.ReadString('\n')
		if errors.Is(err, io.EOF) {
			s.handleDisconnect(conn)
			break
		}
		if errors.Is(err, net.ErrClosed) {
			log.Printf("Connection with %s closed", conn.RemoteAddr().String())
			break
		}
		if err != nil {
			log.Printf("!!!!! Failed to read from the socket: %v", err)
		}
		msg := fmt.Sprintf("%s > %s", conn.RemoteAddr().String(), str)
		runtime.EventsEmit(s.ctx, "message-rx", msg)
		runtime.LogInfo(s.ctx, msg)
	}
}

func (s *Server) handleDisconnect(conn net.Conn) {
	log.Printf("Connection closed by client: %s", conn.RemoteAddr().String())
	s.connectionsMtx.Lock()
	_ = conn.Close()
	delete(s.connections, conn)
	s.connectionsMtx.Unlock()
}
