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
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func VpcSubnets() *schema.Table {
	return &schema.Table{
		Name:         "yandex_vpc_subnets",
		Resolver:     fetchVpcSubnets,
		Multiplex:    client.FolderMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteFolderFilter,
		Columns: []schema.Column{
			{
				Name:        "subnet_id",
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
				Description: "Name of the subnet. The name is unique within the project. 3-63 characters long.",
				Resolver:    schema.PathResolver("Name"),
			},
			{
				Name:        "description",
				Type:        schema.TypeString,
				Description: "Optional description of the subnet. 0-256 characters long.",
				Resolver:    schema.PathResolver("Description"),
			},
			{
				Name:        "labels",
				Type:        schema.TypeJSON,
				Description: "",
				Resolver:    client.ResolveLabels,
			},
			{
				Name:        "network_id",
				Type:        schema.TypeString,
				Description: "ID of the network the subnet belongs to.",
				Resolver:    schema.PathResolver("NetworkId"),
			},
			{
				Name:        "zone_id",
				Type:        schema.TypeString,
				Description: "ID of the availability zone where the subnet resides.",
				Resolver:    schema.PathResolver("ZoneId"),
			},
			{
				Name:        "v_4_cidr_blocks",
				Type:        schema.TypeStringArray,
				Description: "CIDR block.\n The range of internal addresses that are defined for this subnet.\n This field can be set only at Subnet resource creation time and cannot be changed.\n For example,\u00a010.0.0.0/22\u00a0or\u00a0192.168.0.0/24.\n Minimum subnet size is /28, maximum subnet size is /16.",
				Resolver:    schema.PathResolver("V4CidrBlocks"),
			},
			{
				Name:        "v_6_cidr_blocks",
				Type:        schema.TypeStringArray,
				Description: "IPv6 not available yet.",
				Resolver:    schema.PathResolver("V6CidrBlocks"),
			},
			{
				Name:        "route_table_id",
				Type:        schema.TypeString,
				Description: "ID of route table the subnet is linked to.",
				Resolver:    schema.PathResolver("RouteTableId"),
			},
			{
				Name:        "dhcp_options_domain_name_servers",
				Type:        schema.TypeStringArray,
				Description: "",
				Resolver:    schema.PathResolver("DhcpOptions.DomainNameServers"),
			},
			{
				Name:        "dhcp_options_domain_name",
				Type:        schema.TypeString,
				Description: "",
				Resolver:    schema.PathResolver("DhcpOptions.DomainName"),
			},
			{
				Name:        "dhcp_options_ntp_servers",
				Type:        schema.TypeStringArray,
				Description: "",
				Resolver:    schema.PathResolver("DhcpOptions.NtpServers"),
			},
		},
	}
}

func fetchVpcSubnets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	locations := []string{c.FolderId}

	for _, f := range locations {
		req := &vpc.ListSubnetsRequest{FolderId: f}
		it := c.Services.Vpc.Subnet().SubnetIterator(ctx, req)
		for it.Next() {
			res <- it.Value()
		}
	}

	return nil
}