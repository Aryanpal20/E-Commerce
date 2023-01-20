package messages

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SMS(setto string, body string) {
	client := twilio.NewRestClient()
	params := &openapi.CreateMessageParams{}
	params.SetTo(setto)
	params.SetFrom("+15017274199")
	params.SetBody(body)
	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
