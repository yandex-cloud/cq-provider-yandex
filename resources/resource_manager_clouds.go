package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

func ResourceManagerClouds() *schema.Table {
	return &schema.Table{
		Name:        "yandex_resource_manager_clouds",
		Resolver:    fetchResourceManagerClouds,
		Multiplex:   client.IdentityMultiplex,
		IgnoreError: client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Description: "ID of the cloud.",
				Resolver:    client.ResolveResourceId,
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
				Description: "Name of the cloud. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Description of the cloud. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "organization_id",
				Type:        schema.TypeString,
				Description: "ID of the organization that the cloud belongs to.",
				Resolver:    schema.PathResolver("OrganizationId"),
			},
		},
	}

}

func fetchResourceManagerClouds(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	req := &resourcemanager.ListCloudsRequest{}
	it := c.Services.ResourceManager.Cloud().CloudIterator(ctx, req)
	for it.Next() {
		res <- it.Value()
	}

	return nil
}