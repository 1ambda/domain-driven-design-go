resource "google_container_cluster" "primary" {
  project = "${data.google_project.current.id}"

  name = "${local.project}"
  zone = "${element(local.zones, 0)}"
  initial_node_count = "${local.initial_node_count}"

  min_master_version = "${local.min_master_version}"
  node_version = "${local.node_version}"


  network = "${google_compute_network.primary.self_link}"
  subnetwork = "${google_compute_subnetwork.primary.self_link}"

  additional_zones = [
    "${element(local.zones, 1)}",
    "${element(local.zones, 2)}",
  ]

  addons_config {
    network_policy_config {
      disabled = false
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

  //  remove_default_node_pool = true

  lifecycle {
    ignore_changes = [
      "node_pool",
      "ip_allocation_policy",
      "min_master_version",
      "node_version",
    ]
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
      project = "${local.project}"
    }

    tags = [
      "terraform",
    ]
  }

  depends_on = [
  ]
}
