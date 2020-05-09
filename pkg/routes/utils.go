package routes

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

// mustJSONMarshal returns the JSON encoded value or panics on error
func mustJSONMarshal(value interface{}) []byte {
	json, err := json.Marshal(value)
	if err != nil {
		log.Fatal().Err(err).Msg("didn't expect an error during json marshaling")
	}

	return json
}

// generateErrorResponse puts the given message in errorResponse struct and returns
// it as JSON.
func generateErrorResponse(msg string) []byte {
	response := &errorResponse{}
	response.Body.Error = msg
	return mustJSONMarshal(response.Body)
}

// parseRequestBody unmarshals the body into given params from a http.Request
func parseRequestBody(bodyReader io.ReadCloser, decodedBody interface{}) error {
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, decodedBody)
}
