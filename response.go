package washbot

import (
	"github.com/arienmalec/alexa-go"
)

func WelcomeResponse() alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "SSML",
				SSML: "<speak><amazon:emotion name=\"excited\" intensity=\"high\">Hello! Who am I gassing up today?</amazon:emotion></speak>",
			},
			ShouldEndSession: false,
		},
	}
	return r
}

func SSMLResponse(response string) alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "SSML",
				SSML: response,
			},
			ShouldEndSession: false,
		},
	}
	return r
}

func SimpleResponse(response string) alexa.Response {
	return alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "PlainText",
				Text: response,
			},
			ShouldEndSession: false,
		},
	}
}
