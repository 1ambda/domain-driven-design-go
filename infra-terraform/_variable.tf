locals {
  project = "g-street"

  cluster_name = "${local.project}-cluster"

  region = "asia-northeast1"
  zones = [
    "asia-northeast1-a",
    "asia-northeast1-b",
    "asia-northeast1-c",
  ]


  min_master_version = "1.9.7-gke.3"
  node_version = "1.9.7-gke.3"

  initial_node_count = 1
  min_node_count = 1
  max_node_count = 2

  disk_size_gb = "100"
  disk_type = "pd-standard"
  machine_type = "n1-standard-1"

  network_name = "${local.project}-network"
  subnet_name = "${local.project}-subnet"
  ip_cidr_range = "10.127.0.0/20"
}

data "google_project" "current" { }

data "google_compute_default_service_account" "default" {}

data "google_client_config" "current" {}

