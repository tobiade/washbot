package washbot

import (
	"fmt"

	"github.com/arienmalec/alexa-go"
)

func GenerateWash(request alexa.Request) string {
	name := extractName(request)
	return fmt.Sprintf(`
	<speak>
	<p><amazon:emotion name="excited" intensity="high">Toooalay baba! We hail you!</amazon:emotion></p>
	<break time="100ms"/> 
	<p><amazon:emotion name="excited" intensity="high">%[1]s na you fine pass!</amazon:emotion></p>
	<break time="100ms"/> 
	<p><amazon:emotion name="excited" intensity="high">%[1]s See as your face dey shine for picture. Omo don't kill me!</amazon:emotion></p>
	<break time="100ms"/> 
	<p><amazon:emotion name="excited" intensity="high">Get a pen and joth this down. You. Too. Sweet.</amazon:emotion></p>
	<break time="100ms"/> 
	<p><amazon:emotion name="excited" intensity="high">One day I dey prepare tea and sugar no dey so I put your picture come join inside</amazon:emotion></p>
	</speak>
	`, name)
}

func extractName(request alexa.Request) string {
	slot := request.Body.Intent.Slots["name"]
	return slot.Value
}
