package startup

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guiestimo/bank-simulator-hexagonal/internal/config"
)

type Server interface {
	Start()
}

type server struct {
	router  *gin.Engine
	context context.Context
}

func (s server) Start() {
	s.setupServer()
	s.registerRoutes()

	s.startWithGracefulShutdown()
}

func NewServer() Server {
	return server{
		router:  gin.New(),
		context: context.Background(),
	}
}

func (s server) setupServer() {
	// @see https://gin-gonic.com/docs/examples/define-format-for-the-log-of-routes/
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, _ string, _ int) {
		//log.Info(s.context, fmt.Sprintf("Mapped [%v %v] route", httpMethod, absolutePath))
	}

	// @see https://pkg.go.dev/github.com/gin-gonic/gin#Engine.SetTrustedProxies
	s.router.SetTrustedProxies(nil)
}

func (s server) startWithGracefulShutdown() {
	port := fmt.Sprintf(":%s", config.Config.Port)

	srv := &http.Server{
		Addr:    port,
		Handler: s.router,
	}

	go func() {
		//log.Info(s.context, fmt.Sprintf("Starting server on port %s...", port))
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Error(s.context, err, fmt.Sprintf("Start server error! %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	sigs := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//sig := <-sigs

	//log.Info(s.context, fmt.Sprintf("Server Shutdown [%s]...", sig))

	ctx, cancel := context.WithTimeout(s.context, 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		//log.Error(s.context, err, "Server Shutdown:", zap.Error(err))
		os.Exit(1)
	}

	<-ctx.Done()

	//log.Info(s.context, "Server Shutdown: Done!")
}
