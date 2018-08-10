locals {
  # https://cloud.google.com/kubernetes-engine/versioning-and-upgrades
  gke_master_version = "1.10.5-gke.3"
  gke_node_version = "1.10.5-gke.3"

  gke_disk_size     = "100"
  gke_disk_type     = "pd-standard"
  gke_machine_type  = "n1-standard-1"

  gke_initial_node_count = 1
}

resource "google_container_cluster" "primary" {
  project = "${data.google_project.current.id}"

  name = "${local.service}"
  region = "${local.region}"

  min_master_version = "${local.gke_master_version}"
  node_version = "${local.gke_node_version}"

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
  node_count  = "${local.gke_initial_node_count}"

  management {
    auto_repair   = true
    auto_upgrade  = true
  }

  node_config {
    disk_size_gb = "${local.gke_disk_size}"
    disk_type = "${local.gke_disk_type}"
    machine_type = "${local.gke_machine_type}"
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

resource "google_compute_global_address" "g-street" {
  name = "g-street"
}

output "global_address_g-street" {
  value = "${google_compute_global_address.g-street.address}"
}
