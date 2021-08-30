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

# summarized_cidr_blocks = tolist(["10.192.0.0/20", "10.192.16.0/21"])
output "summarized_cidr_blocks" {
  value = data.cidrblock_summarization.vpc.summarized_cidr_blocks
}
