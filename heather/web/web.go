package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type Op struct {
	say   string
	sleep int
}

var opCh chan Op

func sayHandler(w http.ResponseWriter, r *http.Request) {
	qstr := r.URL.Query()
	text := qstr.Get("text")
	sleep := qstr.Get("sleep")
	if text != "" {
		opCh <- Op{say: text}
	}
	if sleep != "" {
		t, err := strconv.Atoi(sleep)
		if err != nil {
			fmt.Printf("Ignoring bad sleep request with value %q\n", sleep)
			return
		}
		opCh <- Op{sleep: t}
	}
}

func main() {
	opCh = make(chan Op, 100)

	// Start processing routine.
	go func() {
		for {
			op := <-opCh
			if op.say != "" {
				fmt.Printf("Saying %q\n", op.say)
				cmd := exec.Command("say", "-v", "Victoria", op.say)
				if err := cmd.Run(); err != nil {
					fmt.Printf("Couldn't say %q: %v\n", op.say, err)
				}
			}
			if op.sleep != 0 {
				fmt.Printf("Sleeping for %d\n", op.sleep)
				time.Sleep(time.Duration(op.sleep) * time.Second)
			}
		}
	}()

	http.HandleFunc("/say", sayHandler)
	http.ListenAndServe(":8080", nil)
}
