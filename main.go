// package main

// func main() {

// }
package main

import (
	"fmt"
	"net/http"
)

func main() {
	flowControl := NewFlowControl()
	myHandler := MyHandler{
		flowControl: flowControl,
	}
	http.Handle("/", &myHandler)

	http.ListenAndServe(":8080", nil)
}

type MyHandler struct {
	flowControl *FlowControl
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieve http request")
	job := &Job{
		DoneChan: make(chan struct{}, 1),
		handleJob: func(job *Job) error {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Hello World"))
			return nil
		},
	}

	h.flowControl.CommitJob(job)
	fmt.Println("commit job to job queue success")
	job.WaitDone()
}
