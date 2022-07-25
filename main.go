// package main

// func main() {

// }
package main

import (
	"fmt"
	"net/http"
	q "sail/pkg/queue"
	"time"

	"github.com/google/uuid"
)

func main() {
	flowControl := q.NewFlowControl()
	myHandler := MyHandler{
		flowControl: flowControl,
	}
	http.Handle("/", &myHandler)

	http.ListenAndServe(":8080", nil)
}

type MyHandler struct {
	flowControl *q.FlowControl
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieve http request")
	job := &q.Job{
		UUID:     uuid.NewString(),
		DoneChan: make(chan struct{}, 1),
		HandleJob: func() error {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Hello World"))
			time.Sleep(5 * time.Second)
			return nil
		},
	}

	h.flowControl.CommitJob(job)
	fmt.Println("commit job to job queue success")
	job.WaitDone()
}
