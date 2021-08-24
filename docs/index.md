---
page_title: "Provider: cidrblock"
subcategory: ""
description: |-
  The cidrblock provider is used to get unallocated subnet prefixes from parent CIDR.
---

# cidrblock provider

The cidrblock provider is used to get unallocated subnet prefixes from parent CIDR.

## Example Usage

```terraform
terraform {
  required_providers {
    cidrblock = {
      source = "amilevskiy/cidrblock"
    }
  }
}

provider "cidrblock" {}
```