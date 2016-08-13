package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/Aorioli/procedural/concerns/version"
)

type descriptionServer struct {
	Intro   string          `json:"intro"`
	Version version.Version `json:"version"`
}

func (m descriptionServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(ContentType, "application/json")
	json.NewEncoder(w).Encode(m)
}

func Description(intro string, v version.Version) http.Handler {
	return descriptionServer{
		Intro:   intro,
		Version: v,
	}
}

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

// EncodeResponse encodes a generic response
func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	switch v := response.(type) {
	case Error:
		w.Header().Add(ContentType, ApplicationJSON)
		w.WriteHeader(v.Status)
		return json.NewEncoder(w).Encode(v)
	case error:
		w.Header().Add(ContentType, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(w).Encode(v)
	}

	w.Header().Add(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
