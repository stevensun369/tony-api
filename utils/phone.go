package utils

import (
	"backend/env"
	"errors"
	"fmt"
	"net/http"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, code string) error {
  return SendSMSO(to, code)
}

func SendTwilio(to string, code string) error {
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
  return err
}

func SendSMSO(to string, code string) error {
  body := fmt.Sprintf("%v este codul de autentificare pentru tony :)", code)
  client := &http.Client{}
  req, err := http.NewRequest("POST", fmt.Sprintf(`https://app.smso.ro/api/v1/send?to=%v&sender=%v&body="%v"`, to, env.SmsoSender, body), nil)
  if err != nil {
    return err
  }
  req.Header.Add("X-Authorization", env.SmsoKey)
  resp, err := client.Do(req)
  if err != nil {
    return err
  }
  if (resp.StatusCode != 200) {
    return errors.New("SMS-ul nu a putut fi transmis")
  }
  return nil
}