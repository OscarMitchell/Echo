package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/OscarMitchell/echo/backend/src/bridge"
	"github.com/OscarMitchell/echo/backend/src/lib"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Server struct {
	ctx         context.Context
	server      net.Listener
	connections *lib.Set[net.Conn]
	mtx         sync.Mutex
}

func NewServer(ctx context.Context, port string) (*Server, error) {
	server, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("!!!!! Error creating a new server: %w", err)
	}
	return &Server{
		ctx:         ctx,
		server:      server,
		connections: lib.NewSet[net.Conn](),
	}, nil
}

func (s *Server) Run() {
	errCount := 0
	for {
		conn, err := s.server.Accept()
		if errors.Is(err, net.ErrClosed) {
			rt.LogInfo(s.ctx, "Server closed")
			break
		}

		// TODO: Add a debouncer instead of just killing the server when
		// encountering error spam. We can continue to attempt to keep alive the
		// server regardless of error state, it doesn't really effect us.
		if err != nil {
			rt.LogErrorf(s.ctx, "Error when attempting to accept incoming connection to server: %v", err)
			errCount++
			if errCount > 50 {
				rt.LogError(s.ctx, "Too many errors - closing server")
				s.Shutdown()
			}
			continue
		}
		errCount = 0

		go s.handleIncomingConnection(conn)
	}
}

func (s *Server) Shutdown() {
	rt.LogInfo(s.ctx, "Shutdown triggered, closing connections and server")
	s.mtx.Lock()
	for conn := range s.connections.All() {
		_ = conn.Close()
	}
	s.mtx.Unlock()
	_ = s.server.Close()
}

func (s *Server) handleIncomingConnection(conn net.Conn) {
	newConnMsg := fmt.Sprintln("New connection from:", conn.RemoteAddr().String())
	rt.LogInfo(s.ctx, newConnMsg)
	// bridge.WriteToTcpConsole(s.ctx, newConnMsg) // TODO: Convert to alert when implemented

	s.mtx.Lock()
	s.connections.Add(conn)
	s.mtx.Unlock()

	buffer := bufio.NewReader(conn)
	for {
		msg, err := buffer.ReadString('\n')
		if errors.Is(err, io.EOF) {
			s.handleDisconnect(conn)
			break
		}
		if errors.Is(err, net.ErrClosed) {
			rt.LogInfof(s.ctx, "Connection with %s closed", conn.RemoteAddr().String())
			break
		}
		if err != nil {
			rt.LogErrorf(s.ctx, "Failed to read from the socket: %v", err)
		}
		bridge.PresentMessage(s.ctx, msg, conn.RemoteAddr().String())
	}
}

func (s *Server) handleDisconnect(conn net.Conn) {
	rt.LogInfof(s.ctx, "Connection closed by client: %s", conn.RemoteAddr().String())
	s.mtx.Lock()
	_ = conn.Close()
	s.connections.Remove(conn)
	s.mtx.Unlock()
}
