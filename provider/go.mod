module github.com/pulumi/pulumi-gardener/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/kyma-incubator/terraform-provider-gardener => github.com/w9n/terraform-provider-gardener v0.0.11
	k8s.io/client-go => k8s.io/client-go v0.17.17 // hotfix TODO
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/kyma-incubator/terraform-provider-gardener v0.0.11
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
)
