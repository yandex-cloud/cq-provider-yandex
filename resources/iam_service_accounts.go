// Code generated by yandex cloud generator; DO NOT EDIT.

package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

func IAMServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:         "yandex_iam_service_accounts",
		Resolver:     fetchIAMServiceAccounts,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Description: "ID of the resource.",
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Description: "ID of the folder that the resource belongs to.",
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
				Description: "Name of the service account.\n The name is unique within the cloud. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the service account. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
		},
	}

}

func fetchIAMServiceAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &iam.ListServiceAccountsRequest{FolderId: f}
		it := c.Services.IAM.ServiceAccount().ServiceAccountIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}
