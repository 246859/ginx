package ginx

import (
	"context"
	"errors"
	"fmt"
	"github.com/246859/ginx/middleware"
	"github.com/dstgo/size"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func defaultEngine() *gin.Engine {
	engine := gin.New()
	engine.NoRoute(middleware.NoRoute())
	engine.NoMethod(middleware.NoMethod())
	engine.Use(middleware.Logger(slog.Default(), "GinX"), middleware.Recovery(slog.Default(), nil))
	return engine
}

// New returns a new server instance
func New(options ...Option) *Server {

	server := new(Server)
	for _, option := range options {
		option(server)
	}

	if server.ctx == nil {
		server.ctx = context.Background()
	}

	if len(server.stopSignals) == 0 {
		server.stopSignals = []os.Signal{syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT}
	}

	if server.options.Mode == "" {
		server.options.Mode = gin.ReleaseMode
	}

	if server.options.Address == "" {
		server.options.Address = ":8080"
	}

	if server.options.MaxShutdownTimeout == 0 {
		server.options.MaxShutdownTimeout = time.Second * 5
	}

	if server.options.ReadTimeout == 0 {
		server.options.ReadTimeout = time.Second * 60
	}

	if server.options.ReadHeaderTimeout == 0 {
		server.options.ReadHeaderTimeout = time.Second * 60
	}

	if server.options.WriteTimeout == 0 {
		server.options.WriteTimeout = time.Second * 60
	}

	if server.options.IdleTimeout == 0 {
		server.options.IdleTimeout = time.Minute * 5
	}

	if server.options.MaxMultipartMemory == 0 {
		server.options.MaxMultipartMemory = int64(size.MB * 10)
	}

	if server.options.MaxHeaderBytes == 0 {
		server.options.MaxHeaderBytes = int(size.MB)
	}

	server.build()

	server.httpserver.Handler = server.Engine
	return server
}

// Server is a simple wrapper for http.Server and *gin.Engine, which is more convenient to use.
// It provides hooks can be executed at certain time, ability to graceful shutdown.
type Server struct {
	ctx context.Context

	httpserver *http.Server
	Engine     *gin.Engine

	BeforeStarting []HookFn
	AfterStarted   []HookFn
	OnShutdown     []HookFn

	stopSignals []os.Signal

	options Options
}

func (s *Server) Run() error {
	slog.InfoContext(s.ctx, fmt.Sprintf("server is listening on %v", s.options.Address))
	if s.options.TLS != nil {
		return s.httpserver.ListenAndServeTLS(s.options.TLS.Cert, s.options.TLS.Key)
	} else {
		return s.httpserver.ListenAndServe()
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}

// Spin runs the server in another go routine, and listening for os signals to graceful shutdown,
// you should use *Server.Spin() in most of time.
func (s *Server) Spin() error {
	notifyContext, signalCancel := signal.NotifyContext(s.ctx, s.stopSignals...)
	defer signalCancel()

	slog.Debug("before starting hooks")
	// execute before starting hooks
	err := s.executeHooks(notifyContext, s.BeforeStarting...)
	if err != nil {
		return err
	}

	runCh := make(chan error)

	go func() {
		runCh <- s.Run()
		close(runCh)
	}()

	slog.Debug("after starting hooks")
	// execute after started hooks
	err = s.executeHooks(notifyContext, s.AfterStarted...)
	if err != nil {
		return err
	}

	// wait for server closed or stop signal
	select {
	case <-notifyContext.Done():
		slog.InfoContext(s.ctx, "received stop signal, ready to shutdown")
	case err := <-runCh:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(s.ctx, "server run failed", slog.Any("error", err))
		} else {
			slog.InfoContext(s.ctx, "http server closed")
		}
	}

	// ready to server shutdown
	shutdownCh := make(chan error)
	timeoutCtx, shutdownCancel := context.WithTimeout(s.ctx, s.options.MaxShutdownTimeout)
	defer shutdownCancel()

	_ = s.Shutdown(timeoutCtx)

	go func() {
		slog.Debug("on shutdown hooks")
		shutdownCh <- s.executeHooks(timeoutCtx, s.OnShutdown...)
		close(shutdownCh)
	}()

	// wait timeout for execute shutdown hooks
	select {
	case <-timeoutCtx.Done():
		slog.ErrorContext(s.ctx, "shutdown timeout")
	case err := <-shutdownCh:
		if err != nil {
			slog.ErrorContext(s.ctx, "shutdown error", slog.Any("error", err))
			return err
		} else {
			slog.InfoContext(s.ctx, "server shutdown")
		}
	}

	// server finished
	return nil
}

func (s *Server) executeHooks(ctx context.Context, hooks ...HookFn) error {
	for _, hook := range hooks {
		if err := hook(ctx); err != nil {
			return err
		}
	}
	return nil
}

// build http server and engine
func (s *Server) build() {
	gin.SetMode(s.options.Mode)
	if s.httpserver == nil {
		s.httpserver = &http.Server{}
	}
	if s.Engine == nil {
		s.Engine = defaultEngine()
	}

	if s.httpserver.Addr == "" {
		s.httpserver.Addr = s.options.Address
	}

	if s.httpserver.ReadTimeout == 0 {
		s.httpserver.ReadTimeout = s.options.ReadTimeout
	}

	if s.httpserver.ReadHeaderTimeout == 0 {
		s.httpserver.ReadHeaderTimeout = s.options.ReadHeaderTimeout
	}

	if s.httpserver.WriteTimeout == 0 {
		s.httpserver.WriteTimeout = s.options.WriteTimeout
	}

	if s.httpserver.MaxHeaderBytes == 0 {
		s.httpserver.MaxHeaderBytes = s.options.MaxHeaderBytes
	}

	if s.httpserver.Handler != nil {
		if engine, ok := s.httpserver.Handler.(*gin.Engine); ok {
			// overlay engine
			s.Engine = engine
		} else {
			panic(fmt.Errorf("expected: *github.com/gin-gonic/gin.Engine, but got %T", s.httpserver.Handler))
		}
	} else {
		s.httpserver.Handler = s.Engine
	}

	if s.Engine.MaxMultipartMemory == 0 {
		s.Engine.MaxMultipartMemory = s.options.MaxMultipartMemory
	}
}
