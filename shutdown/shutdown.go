package shutdown

import (
	"fmt"
	"time"
	"github.com/rs/zerolog/log"
)

var IsShutdown *bool

/*Starts system shutdown in case an error which the system cannot recover from ocuurs.

Changes the IsShutdown global variable to true which blocks all mutation and query types from new incoming requests. 

Then allows for a grace period of ten minutes for all pending requests to be processed and completed.

Then finally shutsdown the system with a fatal error message.

Takes the error message, the resolver name and the id of the record associated with the error
*/
func InitiateShutdown(err error, path string, id int){
	*IsShutdown = true
    ticker := time.NewTicker(time.Second)
	go func(){
	now := time.Now()
	futureTime := now.Add(10 * time.Minute) // Add 10 minutes
	log.Warn().Str("path", path).Int("record_id", id).Msg(err.Error())
	log.Trace().Str("path", path).Msg("system shutdown initiated")
       for{
		select{
		case t:= <-ticker.C:
			fmt.Printf("\rsystem shutdown in: %v", futureTime.Sub(t))
		}
	   }
	}()
	time.Sleep(10*time.Minute)
	fmt.Println("system shutdown completed at : ", time.Now())
	log.Fatal().Msg("system shutdown due to neo4j, postgres synchronization error")
}