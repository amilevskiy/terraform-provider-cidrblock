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

# cidr_blocks = tolist(["10.223.193.0/24", "10.223.192.32/28", "10.223.196.0/22"])
output "cidr_blocks" {
  value = data.cidrblock_allocation.subnet.cidr_blocks
}
