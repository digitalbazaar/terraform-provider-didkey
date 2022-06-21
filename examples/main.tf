terraform {
  required_providers {
    didkey = {
      version = "0.2"
      source  = "hashicorp.com/edu/didkey"
    }
  }
}

provider "didkey" {}

resource "didkey" "example" {
  keepers = {
    "key" = "abc"
  }
}

output "did_key_seed" {
  value = didkey.example.secret_seed_multibase
}

output "did_key_id" {
  value = didkey.example.public_did
}


resource "didkey" "example_two" {
  keepers = {
    "key" = "123"
  }
}

output "did_key_seed_two" {
  value = didkey.example_two.secret_seed_multibase
}

output "did_key_id_two" {
  value = didkey.example_two.public_did
}
