package main

import (
	"fmt"
	"bytes"
	"log"
	"net/http"
	"os/exec"
	"github.com/rs/cors"
	"encoding/json"
)


func pissHandler(w http.ResponseWriter, r *http.Request) {

// taken from https://github.com/rs/cors
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

		cmd := exec.Command("ip","r")
		cmd2 := exec.Command("awk","{print $1}")

		cmd2.Stdin, _ = cmd.StdoutPipe()

                var out bytes.Buffer
                cmd2.Stdout = &out

		_ = cmd2.Start()
		_ = cmd.Run()
		_ = cmd2.Wait()

/*
                err := cmd.Run()
                if err != nil {
                        log.Fatal(err)
                }
*/
		fmt.Fprintln(w, out.String())
	}

type configForm struct {

	leftsubnet string `json:"serverTS"`

/*
        Name string `json:"nameForm"`
        Email string `json:"emailForm"`
        Phone string `json:"phoneForm"`
        Check string `json:"checkForm"`
        Message string `json:"messageForm"`
*/
}

func sendConfigHandler(rw http.ResponseWriter, req *http.Request) {

        c := &configForm{}
        json.NewDecoder(req.Body).Decode(c)

	var out = c.leftsubnet

	fmt.Fprintln(rw,out)
	fmt.Println(rw,out)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/piss", pissHandler)
	mux.HandleFunc("/getServerTS", getServerTSHandler)
	mux.HandleFunc("/sendConfig", sendConfigHandler)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9000", handler))

}
