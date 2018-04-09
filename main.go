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
	fmt.Println(r)
	var request WebHookRequest 
	_ = json.NewDecoder(r.Body).Decode(&request)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	fmt.Println(err)
	var speech = ""
	var displayText = ""
	
	intentName := request.Result.Metadata.IntentName
		speech = "Sorry, this feature is not supported yet."
		displayText = "Sorry, this feature is not supported yet."
	
	if intentName == "network-summary-intent" {
		speech = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
		displayText = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
	} else if intentName == "devices-intent" {
		speech = "I can help you with that.There are 5 devices online"
		displayText = "I can help you with that.There are 5 devices online"
	}

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
