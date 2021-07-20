// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/kms/v1"
)

func KmsSymmetricKeys() *schema.Table {
	return &schema.Table{
		Name:         "yandex_kms_symmetric_keys",
		Resolver:     fetchKmsSymmetricKeys,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "symmetric_key_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    client.ResolveFolderID,
			},
			{
				Name:        "created_at",
				Type:        schema.TypeTimestamp,
				Description: "",
				Resolver:    client.ResolveAsTime,
			},
			{
				Name:        "name",
				Type:        schema.TypeString,
				Description: "Name of the key.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the key.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "status",
				Type:        schema.TypeString,
				Description: "Current status of the key.",
				Resolver:    client.EnumPathResolver("Status"),
			},
			{
				Name:        "primary_version_id",
				Type:        schema.TypeString,
				Description: "ID of the key version.",
				Resolver:    schema.PathResolver("PrimaryVersion.Id"),
			},
			{
				Name:        "primary_version_key_id",
				Type:        schema.TypeString,
				Description: "ID of the symmetric KMS key that the version belongs to.",
				Resolver:    schema.PathResolver("PrimaryVersion.KeyId"),
			},
			{
				Name:        "primary_version_status",
				Type:        schema.TypeString,
				Description: "Status of the key version.",
				Resolver:    client.EnumPathResolver("PrimaryVersion.Status"),
			},
			{
				Name:        "primary_version_algorithm",
				Type:        schema.TypeString,
				Description: "Encryption algorithm that should be used when using the key version to encrypt plaintext.",
				Resolver:    client.EnumPathResolver("PrimaryVersion.Algorithm"),
			},
			{
				Name:        "primary_version_created_at_seconds",
				Type:        schema.TypeBigInt,
				Description: "",
				Resolver:    schema.PathResolver("PrimaryVersion.CreatedAt.Seconds"),
			},
			{
				Name:        "primary_version_created_at_nanos",
				Type:        schema.TypeInt,
				Description: "",
				Resolver:    schema.PathResolver("PrimaryVersion.CreatedAt.Nanos"),
			},
			{
				Name:        "primary_version_primary",
				Type:        schema.TypeBool,
				Description: "Indication of a primary version, that is to be used by default for all cryptographic\n operations that don't have a key version explicitly specified.",
				Resolver:    schema.PathResolver("PrimaryVersion.Primary"),
			},
			{
				Name:        "primary_version_destroy_at_seconds",
				Type:        schema.TypeBigInt,
				Description: "",
				Resolver:    schema.PathResolver("PrimaryVersion.DestroyAt.Seconds"),
			},
			{
				Name:        "primary_version_destroy_at_nanos",
				Type:        schema.TypeInt,
				Description: "",
				Resolver:    schema.PathResolver("PrimaryVersion.DestroyAt.Nanos"),
			},
			{
				Name:        "primary_version_hosted_by_hsm",
				Type:        schema.TypeBool,
				Description: "Indication of the version that is hosted by HSM.",
				Resolver:    schema.PathResolver("PrimaryVersion.HostedByHsm"),
			},
			{
				Name:        "default_algorithm",
				Type:        schema.TypeString,
				Description: "Default encryption algorithm to be used with new versions of the key.",
				Resolver:    client.EnumPathResolver("DefaultAlgorithm"),
			},
			{
				Name:        "rotated_at_seconds",
				Type:        schema.TypeBigInt,
				Description: "",
				Resolver:    schema.PathResolver("RotatedAt.Seconds"),
			},
			{
				Name:        "rotated_at_nanos",
				Type:        schema.TypeInt,
				Description: "",
				Resolver:    schema.PathResolver("RotatedAt.Nanos"),
			},
			{
				Name:        "rotation_period_seconds",
				Type:        schema.TypeBigInt,
				Description: "",
				Resolver:    schema.PathResolver("RotationPeriod.Seconds"),
			},
			{
				Name:        "rotation_period_nanos",
				Type:        schema.TypeInt,
				Description: "",
				Resolver:    schema.PathResolver("RotationPeriod.Nanos"),
			},
			{
				Name:        "deletion_protection",
				Type:        schema.TypeBool,
				Description: "Flag that inhibits deletion of the key",
				Resolver:    schema.PathResolver("DeletionProtection"),
			},
		},
	}
}

func fetchKmsSymmetricKeys(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &kms.ListSymmetricKeysRequest{FolderId: f}
		it := c.Services.Kms.SymmetricKey().SymmetricKeyIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}