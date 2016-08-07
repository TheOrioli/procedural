package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageServer struct {
	Message string
}

func (m MessageServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(ContentType, "application/json")
	fmt.Fprintf(w, `{"message":"%s"}`, m.Message)
}

const ContentType = "Content-Type"

// EncodeResponse encodes a generic response
func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	switch v := response.(type) {
	case Error:
		w.Header().Add(ContentType, "application/json")
		w.WriteHeader(v.Status)
		return json.NewEncoder(w).Encode(v)
	case error:
		w.Header().Add(ContentType, "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(v)
	}

	w.Header().Add(ContentType, "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
