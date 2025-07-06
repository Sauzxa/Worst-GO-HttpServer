package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ConsumerRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumerResponse struct {
	Record Recored `json:"record"`
}

func (s *httpLogServer) handleConsumer(w http.ResponseWriter, r *http.Request) {
	// Get offset from query parameter
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		http.Error(w, "offset is required", http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(offset)
	if err != nil {
		if err == ErrorRecordNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	res := ConsumerResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
