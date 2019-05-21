package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os/exec"
	"strconv"
)

const version = "2019-05-21-A"

func main() {
	port := flag.Int("port", 8080, "web port")
	endpoint := flag.String("endpoint", "/", "endpoint")
	cmd := flag.String("cmd", "cat", "command")

	flag.Parse()

	http.HandleFunc(*endpoint, func(writer http.ResponseWriter, request *http.Request) {
		dump, err := httputil.DumpRequest(request, true)

		if err != nil {
			log.Printf("got error while dupmping request: %v", err)
		}

		log.Printf("-> serving incoming request: %v \n\n", string(dump))

		cmd := exec.Command("sh", "-c", *cmd)

		cmd.Stdin = bytes.NewBuffer(dump)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("exucution returned with error: %s\n", err)
			return
		}
		log.Printf("<- cmd:\n %v \n", string(out))
	})

	portStr := ":" + strconv.Itoa(*port)
	log.Printf("version %s, listening on port: %v", version, portStr)
	log.Panic(http.ListenAndServe(portStr, nil))
}
