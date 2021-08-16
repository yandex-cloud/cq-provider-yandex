package resources

import (
	"context"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func StorageBuckets() *schema.Table {
	return &schema.Table{
		Name:        "yandex_object_buckets",
		Resolver:    fetchStorageBuckets,
		Multiplex:   client.EmptyMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "yandex_object_bucket_serv_side_encrypt_config_rules",
				Resolver:    fetchStorageBucketServerSideEncryptionRules,
				Multiplex:   client.EmptyMultiplex,
				IgnoreError: client.IgnoreErrorHandler,
				Columns: []schema.Column{
					{
						Name:     "storage_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentPathResolver("Name"),
					},
					{
						Name:     "bucket_key_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("BucketKeyEnabled"),
					},
					{
						Name:     "apply_server_side_encryption_by_default_kms_master_key_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ApplyServerSideEncryptionByDefault.KMSMasterKeyID"),
					},
					{
						Name:     "apply_server_side_encryption_by_default_sse_algorithm ",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ApplyServerSideEncryptionByDefault.SSEAlgorithm"),
					},
				},
			},
		},
	}
}

type ObjectsStorage struct {
	Name  string
	Rules []*s3.ServerSideEncryptionRule
}

func fetchStorageBuckets(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client).S3Client
	listResp, err := c.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for _, value := range listResp.Buckets {
		encryptResp, err := c.GetBucketEncryption(&s3.GetBucketEncryptionInput{
			Bucket: value.Name,
		})
		if err != nil {
			continue
		}
		if encryptResp != nil && encryptResp.ServerSideEncryptionConfiguration != nil {
			res <- ObjectsStorage{*value.Name, encryptResp.ServerSideEncryptionConfiguration.Rules}
		}
	}
	return nil
}

func fetchStorageBucketServerSideEncryptionRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	storageObject := parent.Item.(ObjectsStorage)
	for _, rule := range storageObject.Rules {
		res <- rule
	}
	return nil
}