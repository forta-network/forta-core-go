package health

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Constants
const (
	DefaultServerPort = "8090"
)

// HealthChecker checks service health and generates reports.
type HealthChecker func() Reports

// Summarizer checks all reports and creates a summary report.
type Summarizer func(Reports) *Report

// Reporter is a health reporter interface.
type Reporter interface {
	Name() string
	Health() Reports
}

// CheckerFrom makes a health checker handler from Reporter implementations.
func CheckerFrom(summarizer Summarizer, reporters ...Reporter) HealthChecker {
	return func() (allReports Reports) {
		for _, reporter := range reporters {
			reports := reporter.Health()
			reports.ObfuscateDetails()
			for _, report := range reports {
				if len(report.Name) == 0 {
					report.Name = fmt.Sprintf("service.%s", reporter.Name())
				} else {
					report.Name = fmt.Sprintf("service.%s.%s", reporter.Name(), report.Name)
				}
			}
			allReports = append(allReports, reports...)
		}
		if summarizer != nil {
			allReports = append(allReports, summarizer(allReports))
		}
		return
	}
}

// ServerErrorHandler lets the caller do custom stuff when ListenAndServe fails.
type ServerErrorHandler func(err error)

// StartServer starts the health check server to receive and handle incoming health check requests.
func StartServer(ctx context.Context, port string, serverErrHandler ServerErrorHandler, healthChecker HealthChecker) {
	port = strings.ReplaceAll(port, ":", "")
	if len(port) == 0 {
		port = DefaultServerPort
	}
	mux := http.NewServeMux()
	Handle(mux, healthChecker)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if serverErrHandler != nil {
				serverErrHandler(err)
			} else {
				log.WithError(err).Error("health server failed")
			}
		}
	}()
	go func() {
		<-ctx.Done()
		server.Close()
	}()
}

// MakeHandler transforms a health checker and makes it an HTTP handler.
func MakeHandler(healthChecker HealthChecker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		reports := healthChecker()
		if reports == nil {
			reports = Reports{}
		}
		if err := json.NewEncoder(w).Encode(reports); err != nil {
			log.WithError(err).Warn("failed to encode health check reports")
		}
	})
}

// Handle transforms and registers health checker to http.DefaultServeMux.
func Handle(mux *http.ServeMux, healthChecker HealthChecker) {
	mux.Handle("/health", MakeHandler(healthChecker))
}

// Service is a service implementation of a health server, to make things easier.
type Service struct {
	ctx              context.Context
	port             string
	serverErrHandler ServerErrorHandler
	healthChecker    HealthChecker
}

// NewService creates a new service.
func NewService(ctx context.Context, port string, serverErrHandler ServerErrorHandler, healthChecker HealthChecker) *Service {
	return &Service{ctx: ctx, port: port, serverErrHandler: serverErrHandler, healthChecker: healthChecker}
}

// Start starts a service.
func (service *Service) Start() error {
	StartServer(service.ctx, service.port, service.serverErrHandler, service.healthChecker)
	return nil
}

// Stop stops the service.
func (service *Service) Stop() error {
	return nil
}

// Name returns the name of the service.
func (service *Service) Name() string {
	return "health"
}
