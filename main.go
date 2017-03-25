package main

import (
	"fmt"
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"github.com/rs/cors"
	"encoding/json"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//universal error checker from https://gobyexample.com/writing-files

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

//	serverTS string
//	clientTS string


	Leftsubnet string `json:"serverTS"`
	Rightsubnet string `json:"clientTS"`

}

func sendConfigHandler(rw http.ResponseWriter, req *http.Request) {
/*
        c := &configForm{}
	g := req.Body
	err := json.NewDecoder(g).Decode(c)

//	err := p.Decode(c)

	check(err)
*/
//	var out = c.leftsubnet

//	log.SetOutput(os.Stderr)
//	log.Println(c.leftsubnet)

//	fmt.Fprintln(rw,out)
//	fmt.Println(rw,out)
//	fmt.Println(rw,req)
//	fmt.Fprintln(rw,c.leftsubnet)

	decoder := json.NewDecoder(req.Body)

	var t configForm
	err := decoder.Decode(&t)

	check(err)
	
	fmt.Println(t.Leftsubnet)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/piss", pissHandler)
	mux.HandleFunc("/getServerTS", getServerTSHandler)
	mux.HandleFunc("/sendConfig", sendConfigHandler)

	log.SetOutput(os.Stderr)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":9000", handler))

}
