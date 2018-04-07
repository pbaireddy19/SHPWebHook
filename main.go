package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get-device-details", getDeviceDetails).Methods("POST")
	r.HandleFunc("/get-vuln-details", getVulnDetails).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), r))
}

func getDeviceDetails(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to get device details")
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	speech := "There are 5 devices online"
	displayText:= "There are 5 devices online"

	hookResp := WebHookResp {
		speech,
		displayText,
		
	}

	json.NewEncoder(w).Encode(hookResp)
}

func getVulnDetails(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to get Vuln details")
	speech := "One vulnerability found. Family room light is trying to access a malicious website and we blocked it"
	displayText:= "One vulnerability found. Family room light is trying to access a malicious website and we blocked it"

	hookResp := WebHookResp {
		speech,
		displayText,
		
	}

	json.NewEncoder(w).Encode(hookResp)
}
