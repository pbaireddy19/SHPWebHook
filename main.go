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

	log.Fatal(http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), r))
}

func getDeviceDetails(w http.ResponseWriter, r *http.Request) {

	speech := "Sorry, this feature is not supported yet."
	displayText := "Sorry, this feature is not supported yet."

	fmt.Println("Received request to get device details")
	fmt.Println(r)
	var request WebHookRequest 
	_ = json.NewDecoder(r.Body).Decode(&request)
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(request)
	fmt.Println(body)
	fmt.Println(err)
	
	
	intentName := request.QueryResult.Intent.DisplayName
		

		switch intentName {
		case "OnlineDeviceCountIntent":
			speech = "I can help you with that.There are 5 devices online"
			displayText = "I can help you with that.There are 5 devices online"
		case "SummaryOfNetwork":
			speech = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
			displayText = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
		default:
			speech = "Sorry, this feature is not supported yet."
			displayText = "Sorry, this feature is not supported yet."
		}
	
	/*if intentName == "network-summary-intent" {
		speech = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
		displayText = "Okay, there are 5 online devices on your network. 1 device is currently blocked. We are trying to determine the status of McAfee Antivirus protection for your devices. You also have unread notifications. Would you like me to read out a few?"
	} else if intentName == "devices-intent" {
		
	} else if intentName == "network-summary-intent - yes" {
		speech = "Okay. Here are your recent notifications.Family room light is trying to access a malicious website and we blocked it"
		displayText = "Okay. Here are your recent notifications.Family room light is trying to access a malicious website and we blocked it"
	}*/
		fmt.Println(speech)
	hookResp := WebHookResp{
		FulfillmentText : displayText,
		//Payload : Payload{Google{RichResponse{Items{{SimpleResponse{speech}}}}}}
	}
	hookResp.Payload.Google.RichResponse.Items = append(hookResp.Payload.Google.RichResponse.Items,SimpleResponse{speech})

	

	json.NewEncoder(w).Encode(hookResp)
}


