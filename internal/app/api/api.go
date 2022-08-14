package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"service_history/internal/app/contract"
	"service_history/internal/app/proto"
	"service_history/internal/app/requester"
	"service_history/internal/app/service"
	"service_history/pkg/config"
	"service_history/pkg/store/postgres"
	"time"
)

type Server struct {
	service   contract.IService
	repo      contract.IProvider
	requester contract.IRequester
}

func NewApiServer(c config.IConfig) *Server {
	return &Server{
		service:   NewServer(c.GetPort()),
		repo:      postgres.NewProvider(c.GetSQL()),
		requester: requester.New(c.GetPeriod()),
	}
}

func (s *Server) StartServ() error {
	var lis net.Listener
	err := s.repo.Open()
	if err != nil {
		<-time.After(1 * time.Second)
		err = s.repo.Open()
	}
	if err != nil {
		return err
	}

	err = s.requester.Start(s.repo.GetConn())
	if err != nil {
		return err
	}

	rpcServ := grpc.NewServer()
	proto.RegisterHistoryServer(rpcServ, service.NewHistoryServer(s.repo.GetConn()))

	go func() {
		if lis, err = net.Listen("tcp", os.Getenv("RPC_PORT")); err != nil {
			log.Println(err)
			return
		}
		log.Println("Start grpc server on", os.Getenv("RPC_PORT"))
		if err = rpcServ.Serve(lis); err != nil {
			log.Println(err)
			return
		}
	}()

	return s.service.Run(context.Background())
}
