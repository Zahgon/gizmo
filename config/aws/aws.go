package aws // import "github.com/NYTimes/gizmo/config/aws"

import (
	"github.com/bradfitz/gomemcache/memcache"
)

const (
	// RegionUSEast1 is a helper constant for AWS configs.
	RegionUSEast1 = "us-east-1"
	// RegionUSWest is a helper constant for AWS configs.
	RegionUSWest = "us-west-1"
)

type (
	// Config holds common AWS credentials and keys.
	Config struct {
		AccessKey       string `envconfig:"AWS_ACCESS_KEY"`
		MFASerialNumber string `envconfig:"AWS_MFA_SERIAL_NUMBER"`
		Region          string `envconfig:"AWS_REGION"`
		RoleARN         string `envconfig:"AWS_ROLE_ARN"`
		SecretKey       string `envconfig:"AWS_SECRET_KEY"`
		SessionToken    string `envconfig:"AWS_SESSION_TOKEN"`
		// Endpoint is an optional endpoint URL (hostname only or fully qualified URI)
		// that overrides the default endpoint for a client. Leave the value as "nil"
		// to use the default endpoint. Currently, only the gizmo SNS Publisher is
		// using this value. Note that AWS emulators (such as localstack) often have
		// different endpoint URL for each emulated service.
		EndpointURL *string `envconfig:"AWS_ENDPOINT_URL"`
	}

	// S3 holds the info required to work with Amazon S3.
	S3 struct {
		Config
		Bucket string `envconfig:"AWS_S3_BUCKET_NAME"`
	}

	// DynamoDB holds some basic info required to work with Amazon DynamoDB.
	DynamoDB struct {
		Config
		TableName string `envconfig:"AWS_DYNAMODB_TABLE_NAME"`
	}

	// ElastiCache holds the basic info required to work with
	// Amazon ElastiCache.
	ElastiCache struct {
		Config
		ClusterID string `envconfig:"AWS_ELASTICACHE_CLUSTER_ID"`
	}
)

// MustClient will use the cache cluster ID to describe
// the cache cluster and instantiate a memcache.Client
// with the cache nodes returned from AWS.
func (e *ElastiCache) MustClient() *memcache.Client { _ = "STUB: not implemented"; return nil }

// LoadConfigFromEnv will attempt to load the Config struct
// from environment variables.
func LoadConfigFromEnv() Config { _ = "STUB: not implemented"; return *new(Config) }

// LoadDynamoDBFromEnv will attempt to load the DynamoDB struct
// from environment variables. If not populated, nil
// is returned.
func LoadDynamoDBFromEnv() DynamoDB { _ = "STUB: not implemented"; return *new(DynamoDB) }

// LoadS3FromEnv will attempt to load the S3 struct
// from environment variables.
func LoadS3FromEnv() S3 { _ = "STUB: not implemented"; return *new(S3) }

// LoadElastiCacheFromEnv will attempt to load the ElasiCache struct
// from environment variables.
func LoadElastiCacheFromEnv() ElastiCache { _ = "STUB: not implemented"; return *new(ElastiCache) }
