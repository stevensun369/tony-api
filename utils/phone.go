package utils

import (
	"backend/env"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, code string) (ok bool) {
  client := twilio.NewRestClientWithParams(
    twilio.ClientParams {
      Username: env.TwilioUsername,
      Password: env.TwilioPassword,
    },
  )

  params := &openapi.CreateMessageParams {}
  params.SetTo(to)
  params.SetFrom(env.TwilioNumber)
  params.SetBody(code[0:3] + "-" + code[3:6] + " este codul dumneavoastră de autentificare pentru elmtree.")

  _, err := client.Api.CreateMessage(params)
  if err != nil {
    return false
  } else {
    return true
  }
}