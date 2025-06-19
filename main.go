package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/derricw/soulmsg/msg"
)

const conjunctionRate = 0.2

func randBool(trueProportion float32) bool {
	return rand.Float32() < trueProportion
}

func response(w http.ResponseWriter, req *http.Request) {
	doConjuction := randBool(conjunctionRate)
	msg := fmt.Sprintf("%s\n", msg.RandomMessage(doConjuction))
	log.Debug().Str("soul_message", msg).Str("source_ip", req.RemoteAddr).Str("user_agent", req.UserAgent()).Msg("received request")
	fmt.Fprintf(w, msg)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	http.HandleFunc("/", response)
	err := http.ListenAndServe(":"+httpPort, nil)
	fmt.Printf("%s", err)
}
