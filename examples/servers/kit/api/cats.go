package api

import (
	"context"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	ocontext "golang.org/x/net/context"

	"github.com/NYTimes/gizmo/examples/nyt"
)

// GRPC layer, add the service-wide middleware ourselves
func (s service) GetCats(ctx ocontext.Context, r *google_protobuf.Empty) (*CatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SHARED BUSINESS LAYER
func (s service) getCats(ctx context.Context, _ interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BIZ LOGIC (SHOULD/COULD BE IN SOME BIZ PACKAGE)
func semToCat(res []*nyt.SemanticConceptArticle) *CatsResponse {
	_ = "STUB: not implemented"

	// translate Semantic to CatResponse
	return nil
}
