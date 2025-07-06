package server

import (
	"encoding/json"
	"net/http"
)

// getting the post req nabnilha struct n appendiha ll records  ou nreturni offset as responce (json encouding)

type ProducerRequest struct {
	Record Recored
}
type ProducerResponse struct {
	Offset uint64 `json:"offset"`
}

func (s *httpLogServer) handleProducer(w http.ResponseWriter, r *http.Request) {
	var req ProducerRequest
	err := json.NewDecoder(r.Body).Decode(&req) // req body n decodeha l struct
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	off, err := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ProducerResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
