package washbot

import (
	"github.com/arienmalec/alexa-go"
)

func Handler(request alexa.Request) (alexa.Response, error) {
	return handleRequest(request), nil
}
func handleRequest(request alexa.Request) alexa.Response {
	requestType := request.Body.Type
	switch requestType {
	case "LaunchRequest":
		return handleLaunchRequest(request)
	case "IntentRequest":
		return handleIntentRequest(request)
	default:
		return handleUnknownRequest(request)
	}
}

func handleLaunchRequest(request alexa.Request) alexa.Response {
	return WelcomeResponse()
}

func handleIntentRequest(request alexa.Request) alexa.Response {
	switch request.Body.Intent.Name {
	case "WashIntent":
		return SSMLResponse(GenerateWash(request))
	default:
		return SimpleResponse("Sorry, I can't do what you asked of me.")
	}
}

func handleUnknownRequest(request alexa.Request) alexa.Response {
	return SimpleResponse("Sorry, I can't handle your request")
}
