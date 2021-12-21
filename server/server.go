package server

import (
	"context"
	"fmt"
	"go_web/config"
	"net/http"
)

var c = config.Conf

func NewServer(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	server := &http.Server{
		Addr:    c.System.Addr,
		Handler: newRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("web server shutdown completely")
			} else {
				fmt.Println(err.Error())
				fmt.Println("Web server closed with exceptions")
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("http: shutting down web server")
	err := server.Close()
	if err != nil {
		fmt.Println("Fail to shut down web server")
		fmt.Println(err)
	}
}
