package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/kyma-incubator/terraform-provider-gardener/client"
	"github.com/kyma-incubator/terraform-provider-gardener/shoot"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kube_file": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gardener_shoot": shoot.ResourceShoot(),
		},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &client.Config{
		KubeFile: d.Get("kube_file").(string),
	}
	return client.New(config)
}
