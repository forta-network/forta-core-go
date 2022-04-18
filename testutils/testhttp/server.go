package testhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/forta-network/forta-core-go/utils/apiutils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type server struct {
	cfg MockHttpConfig
}

type MockHttpConfig struct {
	MockAPIs []MockApi
}

type MockApi struct {
	Method       string
	Path         string
	ResponseBody interface{}
	Status       int
}

func (s *server) addRoute(r *mux.Router, api MockApi) {
	b, err := json.Marshal(api.ResponseBody)
	if err != nil {
		log.Panic(err)
	}
	r.HandleFunc(api.Path, func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write(b); err != nil {
			log.Panic(err)
		}
		if api.Status != 0 {
			writer.WriteHeader(api.Status)
		}
	}).Methods(api.Method)
}

func (s *server) createRoutes() *mux.Router {
	r := mux.NewRouter()
	for _, api := range s.cfg.MockAPIs {
		s.addRoute(r, api)
	}
	return r
}

func (s *server) ServerHost() string {
	return fmt.Sprintf("localhost:%d", 7777)
}

func (s *server) ServerURL() string {
	return fmt.Sprintf("http://%s", s.ServerHost())
}

func (s *server) Start(ctx context.Context) error {
	go func() {
		_ = apiutils.ListenAndServe(ctx, &http.Server{
			Handler:      s.createRoutes(),
			Addr:         s.ServerHost(),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}, "started alerts api")
	}()
	return s.WaitForStart(ctx)
}

func (s *server) WaitForStart(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := net.DialTimeout("tcp", s.ServerHost(), 500*time.Millisecond)
			if err != nil {
				continue
			}
			defer conn.Close()
			return nil
		}
	}
}

func NewHttpServer(cfg MockHttpConfig) *server {
	return &server{
		cfg: cfg,
	}
}
