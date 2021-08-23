module github.com/amilevskiy/terraform-provider-cidrblock

go 1.16

require (
	github.com/hashicorp/terraform-plugin-docs v0.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	inet.af/netaddr v0.0.0-20210729200904-31d5ee66059c
	github.com/amilevskiy/cidrblock v0.0.2
)

replace github.com/amilevskiy/terraform-provider-cidrblock/internal/provider => ./internal/provider
