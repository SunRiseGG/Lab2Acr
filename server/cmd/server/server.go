package main

import (
	"context"
	"fmt"
	"github.com/SunRiseGG/lab2/server/scenarios"
	"net/http"
)

type HttpPortNumber int

type BankingApiServer struct {
  Port HttpPortNumber
  BankingHandler scenarios.HttpHandlerFunc
  Server *http.Server
}

func (s *BankingApiServer) Start() error {
    if s.BankingHandler == nil {
      return fmt.Errorf("Banking HTTP Handler is not defined - cannot start!")
    }
    if s.Port == 0 {
      return fmt.Errorf("Port is not defined!")
    }
    handler := new(http.ServeMux)
    handler.HandleFunc("/banking", s.BankingHandler)
    s.Server = &http.Server{
      Addr: fmt.Sprintf(":%d", s.Port),
      Handler: handler,
    }

    return s.Server.ListenAndServe()
}

func (s *BankingApiServer) Stop() error {
  if s.Server == nil {
    return fmt.Errorf("Server wasn't started!")
  }
  return s.Server.Shutdown(context.Background())
}
