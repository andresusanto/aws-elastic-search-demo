package es

import (
	"github.com/andresusanto/aws-elastic-search-demo/ingester/internal/config"
	"github.com/aws/aws-sdk-go/aws/defaults"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic/v7"
	"github.com/sha1sum/aws_signing_client"
)

// NewClient returns new elastic search client based from environments
func NewClient(c *config.Config) (*elastic.Client, error) {
	if !c.SignESClient {
		return elastic.NewClient(
			elastic.SetURL(c.ESEndpoint),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false),
		)
	}

	signer := v4.NewSigner(
		defaults.CredChain(
			defaults.Config(),
			defaults.Handlers(),
		),
	)
	awsClient, err := aws_signing_client.New(signer, nil, "es", c.Region)
	if err != nil {
		return nil, err
	}

	return elastic.NewClient(
		elastic.SetURL(c.ESEndpoint),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
}
