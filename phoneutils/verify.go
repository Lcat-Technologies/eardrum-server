package phoneutils

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

//sends an OTP code to the phone number passed as an argument
func SendOtp(phone_number string) error {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phone_number)
	params.SetChannel("sms")

	resp, err := TwilioInstance.Client.VerifyV2.CreateVerification(TwilioInstance.VerifyServiceSid, params)

	if err != nil {
		log.Error().Str("phone_number", phone_number).Msg(err.Error())
		return errors.New("error sending verification code. Please verify format")
	} else {
		log.Info().Str("phone_number", phone_number).Str("verification_resource", *resp.Sid).Msg(fmt.Sprintf("Sent verification code"))
	}
	return nil
}

//verifies the validity of an OTP code with regard to the phone number. Takes both the OTP and the phone number as arguments
func CheckOtp(phone_number string, code string) error {

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(phone_number)
	params.SetCode(code)

	resp, err := TwilioInstance.Client.VerifyV2.CreateVerificationCheck(TwilioInstance.VerifyServiceSid, params)

	if err != nil {
		log.Error().Str("phone_number", phone_number).Str("code", code).Msg(err.Error())
		return errors.New("error vefifying code")
	} else if *resp.Status == "approved" {
		log.Info().Str("phone_number", phone_number).Str("code", code).Msg("Entered correct verification code!")
	} else {
		log.Info().Str("phone_number", phone_number).Str("code", code).Msg("Entered incorrect verification code!")
		return errors.New("entered incorrect code")
	}
	return nil
}
