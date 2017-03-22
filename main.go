package main

import (
	"fmt"
	"bytes"
	"log"
	"net/http"
	"os/exec"
	"github.com/rs/cors"
)

// taken from https://github.com/rs/cors

func pissHandler(w http.ResponseWriter, r *http.Request) {

//		w.Header().Set("Access-Control-Allow-Origin", "http://10.0.0.26")

                cmd := exec.Command("strongswan","statusall")
                var out bytes.Buffer
                cmd.Stdout = &out
                err := cmd.Run()
                if err != nil {
                        log.Fatal(err)
                }

		fmt.Fprintln(w, out.String())
	}


func getServerTSHandler(w http.ResponseWriter, r *http.Request) {

		cmd := exec.Command("ip","a")
                var out bytes.Buffer
                cmd.Stdout = &out
                err := cmd.Run()
                if err != nil {
                        log.Fatal(err)
                }

		fmt.Fprintln(w, out.String())
	}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/piss", pissHandler)
	mux.HandleFunc("/getServerTS", getServerTSHandler)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9000", handler))

}
