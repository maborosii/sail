package receiver

import (
	"fmt"
	"net/http"
	"sail/internal/receiver/routers"
	"time"
)

func Receiver(port int) {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
