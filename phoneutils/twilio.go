package phoneutils

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/twilio/twilio-go"
)

var TwilioInstance *Twilio

//Twilio struct contains connection information to a twilio service
type Twilio struct {
	Client           *twilio.RestClient
	VerifyServiceSid string
}

/*Loads TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN and VERIFY_SERVICE_SID environment variables and returns a pointer to `Twilio` struct. 

Warning: Does not throw an error incase of inaccurate environment variables. */ 
func NewTwilio() *Twilio {
	TWILIO_ACCOUNT_SID := os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN := os.Getenv("TWILIO_AUTH_TOKEN")
	VERIFY_SERVICE_SID := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	return &Twilio{
		Client:           client,
		VerifyServiceSid: VERIFY_SERVICE_SID,
	}
}

// Initializes the TwilioInstance global variable
func InitializeTwilio() {
	TwilioInstance = NewTwilio()
	log.Trace().Str("verifyservicesid", TwilioInstance.VerifyServiceSid).Msg("completed twilio initialization")
}
