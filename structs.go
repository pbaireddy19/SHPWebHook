
package main


type WebHookResp struct {
	FulfillmentText     string `json:"fulfillmentText"`
	FulfillmentMessages []struct {
		Card struct {
			Title    string `json:"title"`
			Subtitle string `json:"subtitle"`
			ImageURI string `json:"imageUri"`
			Buttons  []struct {
				Text     string `json:"text"`
				Postback string `json:"postback"`
			} `json:"buttons"`
		} `json:"card"`
	} `json:"fulfillmentMessages"`
	Source  string `json:"source"`
	Payload Payload `json:"payload"`
	OutputContexts []struct {
		Name          string `json:"name"`
		LifespanCount int    `json:"lifespanCount"`
		Parameters    struct {
			Param string `json:"param"`
		} `json:"parameters"`
	} `json:"outputContexts"`
	FollowupEventInput struct {
		Name         string `json:"name"`
		LanguageCode string `json:"languageCode"`
		Parameters   struct {
			Param string `json:"param"`
		} `json:"parameters"`
	} `json:"followupEventInput"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
} 


type RichResponse struct {
	Items []SimpleResponse `json:"items"`
}

type Google struct {
	ExpectUserResponse bool `json:"expectUserResponse"`
	RichResponse  RichResponse  `json:"richResponse"`
}

type Payload struct {
	Google Google `json:"google"`
	Facebook struct {
		Text string `json:"text"`
	} `json:"facebook"`
	Slack struct {
		Text string `json:"text"`
	} `json:"slack"`
}

type WebHookRequest struct {
	ResponseID  string `json:"responseId"`
	Session     string `json:"session"`
	QueryResult struct {
		QueryText  string `json:"queryText"`
		Parameters struct {
			Param string `json:"param"`
		} `json:"parameters"`
		AllRequiredParamsPresent bool   `json:"allRequiredParamsPresent"`
		FulfillmentText          string `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		OutputContexts []struct {
			Name          string `json:"name"`
			LifespanCount int    `json:"lifespanCount"`
			Parameters    struct {
				Param string `json:"param"`
			} `json:"parameters"`
		} `json:"outputContexts"`
		Intent struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"intent"`
		IntentDetectionConfidence int `json:"intentDetectionConfidence"`
		DiagnosticInfo            struct {
		} `json:"diagnosticInfo"`
		LanguageCode string `json:"languageC	ode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
	} `json:"originalDetectIntentRequest"`
}
