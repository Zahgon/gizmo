package service

import (
	"net/http"
)

func (s *MixedService) GetMostPopular(r *http.Request) (int, interface{}, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string { _ = "STUB: not implemented"; return "" }
