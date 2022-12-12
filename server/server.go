package server

import "net/http"

type Country struct {
	Name      string
	Languague string
}

var countries []*Country = []*Country{}

func new(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
