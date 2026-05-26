package service

import (
	"github.com/NYTimes/gizmo/examples/nyt"
	"github.com/NYTimes/gizmo/pubsub"
	"github.com/NYTimes/gizmo/pubsub/aws"
	"github.com/sirupsen/logrus"
)

var (
	Log = logrus.New()

	sub pubsub.Subscriber

	client nyt.Client

	articles []nyt.SemanticConceptArticle
)

type Config struct {
	MostPopularToken string
	SemanticToken    string
	Log              *string
	SQS              aws.SQSConfig
}

func Init() { _ = "STUB: not implemented"; return }

func Run() (err error) { _ = "STUB: not implemented"; return nil }

// do something!

func metricsNamespace() string {
	_ = "STUB: not implemented"
	// get only server base name
	return ""
}

// set it up to be paperboy.servername

// add the 'apps' prefix  to keep things neat
