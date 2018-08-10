resource "google_container_cluster" "primary" {
  project = "${data.google_project.current.id}"

  name = "${local.service}"
  region = "${local.region}"

  min_master_version = "${local.min_master_version}"
  node_version = "${local.node_version}"

  network = "${google_compute_network.primary.self_link}"
  subnetwork = "${google_compute_subnetwork.primary.self_link}"

  addons_config {
    network_policy_config {
      disabled = true
    }

    kubernetes_dashboard {
      disabled = false
    }

    http_load_balancing {
      disabled = false
    }

    horizontal_pod_autoscaling {
      disabled = false
    }
  }

  lifecycle {
    ignore_changes = [
      "node_pool",
      "network",
      "subnetwork",
    ]
  }

  # defaul node is not auto-upgraded, auto-repaired
  remove_default_node_pool = true
  node_pool {
    name = "default-pool"
  }

  depends_on = [
  ]
}


resource "google_container_node_pool" "primary" {
  name        = "primary"
  cluster     = "${google_container_cluster.primary.name}"
  region      = "${local.region}"
  node_count  = "${local.initial_node_count}"

  management {
    auto_repair   = true
    auto_upgrade  = true
  }

  node_config {
    disk_size_gb = "100"
    disk_type = "pd-standard"
    machine_type = "n1-standard-1"
    preemptible = false

    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/servicecontrol",
      "https://www.googleapis.com/auth/service.management.readonly",
      "https://www.googleapis.com/auth/trace.append"
    ]

    labels {
      project = "${local.service}"
    }

    tags = [
      "terraform",
    ]
  }

}
