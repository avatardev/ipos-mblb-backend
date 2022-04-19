package httputil

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func ListenAndServe(addr string, m http.Handler) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		log.Printf("[ListenAndServe] listening to %v", addr)
		err := http.ListenAndServe(addr, m)
		if err != nil {
			log.Fatalf("[ListenAndServe] error listening to http: %v", err)
		}
	}()

	<-c
}
