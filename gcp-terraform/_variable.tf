locals {
  service = "g-street"

  cluster_name = "${local.service}-cluster"

  region = "asia-northeast1"
  zones = [
    "asia-northeast1-a",
    "asia-northeast1-b",
    "asia-northeast1-c",
  ]


  network_name = "${local.service}-network"
  subnet_name = "${local.service}-subnet"
  ip_cidr_range = "10.127.0.0/20"
}

data "google_project" "current" { }

data "google_compute_default_service_account" "default" {}

data "google_client_config" "current" {}

