package service

import "net/http"

func (s *JSONService) GetCats(r *http.Request) (int, interface{}, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
