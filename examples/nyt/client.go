package nyt

type (
	Client interface {
		GetMostPopular(string, string, uint) ([]*MostPopularResult, error)
		SemanticConceptSearch(string, string) ([]*SemanticConceptArticle, error)
	}
	ClientImpl struct {
		mostPopularToken string
		semanticToken    string
	}
)

func NewClient(mostPopToken, semanticToken string) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func (c *ClientImpl) GetMostPopular(resourceType string, section string, timePeriodDays uint) ([]*MostPopularResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientImpl) SemanticConceptSearch(conceptType, concept string) ([]*SemanticConceptArticle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientImpl) do(uri string) (body []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
