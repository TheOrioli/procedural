package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error is a wrapper struct that provides HTTP Status code
// to an error
type Error struct {
	Cause  error
	Status int
}

func (e Error) Error() string {
	return fmt.Sprintf("%s", e.Cause.Error())
}

// Err wraps a cause with a status code
func Err(cause error, status int) error {
	return Error{
		Cause:  cause,
		Status: status,
	}
}

// MarshalJSON satisfies json.Marshaler
func (e Error) MarshalJSON() ([]byte, error) {
	r := map[string]string{
		"error": e.Cause.Error(),
	}

	return json.Marshal(r)
}

func CheckError(w http.ResponseWriter, response interface{}) (bool, error) {
	switch v := response.(type) {
	case Error:
		w.Header().Add(ContentType, ApplicationJSON)
		w.WriteHeader(v.Status)
		return true, json.NewEncoder(w).Encode(v)
	case error:
		w.Header().Add(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		return true, json.NewEncoder(w).Encode(v)
	}

	return false, nil
}
