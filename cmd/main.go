package main

import (
	"net"
	"net/http"
	"os"

	"github.com/ToirovSadi/http/cmd/app"
	"github.com/ToirovSadi/http/pkg/banners"
)

func main() {
	host := "0.0.0.0"
	port := "8080"
	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}
func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}
	server.Init()
	return srv.ListenAndServe()
}
