package nyttest

import "github.com/NYTimes/gizmo/examples/nyt"

type Client struct {
	TestGetMostPopular        func(string, string, uint) ([]*nyt.MostPopularResult, error)
	TestSemanticConceptSearch func(string, string) ([]*nyt.SemanticConceptArticle, error)
}

func (t *Client) GetMostPopular(resourceType, section string, timeframe uint) ([]*nyt.MostPopularResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Client) SemanticConceptSearch(conceptType, concept string) ([]*nyt.SemanticConceptArticle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
