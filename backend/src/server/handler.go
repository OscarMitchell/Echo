package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ServerHandler struct {
	ctx    context.Context
	port   string
	server *Server
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{port: ""}
}

// ***************
// Wails Callbacks
// ***************

func (h *ServerHandler) OnStartup(ctx context.Context) {
	h.ctx = ctx
}

func (h *ServerHandler) OnShutdown(ctx context.Context) {
	h.StopServer()
}

// **********
// Logic Code
// **********

func (h *ServerHandler) StartServer() error {
	if h.port == "" {
		errMsg := "No port has been set, cannot start server without a port"
		runtime.LogError(h.ctx, errMsg)
		return fmt.Errorf(errMsg)
	}

	serv, err := NewServer(h.ctx, h.port)
	if err != nil {
		runtime.LogError(h.ctx, err.Error())
		return err
	}

	go serv.Run()
	runtime.LogInfo(h.ctx, "Successfully setup server, waiting for connections...")
	h.server = serv
	return nil
}

func (h *ServerHandler) SetPort(port int) error {
	if h.server != nil {
		runtime.LogWarning(h.ctx, "Tried to change port on running server")
		return fmt.Errorf("Cannot change the port of a running server")
	}
	portString := ":" + strconv.Itoa(port)
	h.port = portString
	return nil
}

func (h *ServerHandler) StopServer() {
	if h.server == nil {
		runtime.LogWarning(h.ctx, "Attempt to stop the server but no server was runnning")
		return
	}
	h.server.Shutdown()
	h.server = nil
}

func (h *ServerHandler) IsUp() bool {
	return h.server != nil
}
