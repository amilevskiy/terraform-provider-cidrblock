# Terraform provider cidrblock

Available in the [Terraform Registry](https://registry.terraform.io/providers/amilevskiy/cidrblock/latest).

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `build` command: 
```sh
$ go build
```

## Using the provider

```terraform
data "cidrblock_allocation" "subnet" {
  cidr_block = "10.223.192.0/18"
  exclude_cidr_blocks = [
    "10.223.192.0/28",
    "10.223.192.16/28",
    "10.223.194.0/28",
    "10.223.200.0/28"
  ]
  prefix_lengths = [24, 28, 22]
}

data "cidrblock_summarization" "vpc" {
  cidr_blocks = [
    "10.192.0.0/22",
    "10.192.4.0/22",
    "10.192.8.0/22",
    "10.192.12.0/22",
    "10.192.16.0/22",
    "10.192.20.0/22",
  ]
}
```
