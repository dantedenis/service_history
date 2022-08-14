package api

import (
	"context"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"service_history/internal/app/contract"
	"syscall"
)

type server struct {
	serv *fasthttp.Server
	port string
}

func NewServer(port string) contract.IService {
	s := &server{
		serv: &fasthttp.Server{},
		port: port,
	}

	s.serv.Handler = s.configureRouter().Handler

	return s
}

func (s *server) Run(ctx context.Context) error {

	go func() {
		sigint := make(chan os.Signal, 1)
		defer close(sigint)

		signal.Notify(sigint, syscall.SIGINT, syscall.SIGQUIT)

		select {
		case signalHandle := <-sigint:
			log.Printf("Server is shutting down to: %+v\n\n", signalHandle)
		case <-ctx.Done():
			log.Println("Server is shutting down to: context.Done()")
		}

		if err := s.Stop(); err != nil {
			fmt.Println("Failed stop server, err:", err.Error())
		}
	}()

	log.Println("Run server:", s.port)
	return s.serv.ListenAndServe(s.port)
}

func (s *server) Stop() error {
	log.Println("server stop")
	return s.serv.Shutdown()
}

func (s *server) configureRouter() *router.Router {
	rout := router.New()

	rout.GET("/health", s.health)

	return rout
}
