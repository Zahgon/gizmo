package api

import (
	"context"
	"net/http"

	ocontext "golang.org/x/net/context"

	"github.com/NYTimes/gizmo/examples/nyt"
)

// GRPC LAYER, add the middleware layer ourselves
func (s service) GetMostPopularResourceTypeSectionTimeframe(ctx ocontext.Context, req *GetMostPopularResourceTypeSectionTimeframeRequest) (*MostPopularResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SHARED BIZ LAYER
func (s service) getMostPopular(ctx context.Context, r interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CUSTOM HTTP REQUEST DECODER
func decodeMostPopularRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BIZ LOGIC THAT SHOULD/COULD LIVE SOMEWHERE ELSE?
func mpToMP(res []*nyt.MostPopularResult) *MostPopularResponse {
	_ = "STUB: not implemented"
	return nil
}
