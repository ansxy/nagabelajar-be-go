package http

import (
	"log"
	"net/http"

	"github.com/ansxy/nagabelajar-be-go/config"
)

func NewHTTPServer(addr string, handler http.Handler, conf *config.Config) (err error) {
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// run server
	log.Printf("[http-server-online] %v\n", addr)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("[http-server-failed] \n%v\n", err.Error())
		return err
	}

	return
}
