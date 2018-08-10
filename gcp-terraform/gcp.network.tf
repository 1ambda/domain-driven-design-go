resource "google_compute_network" "primary" {
  name                    = "${local.network_name}"
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "primary" {
  name                     = "${local.subnet_name}"
  ip_cidr_range            = "${local.ip_cidr_range}"
  network                  = "${google_compute_network.primary.self_link}"
  region                   = "${local.region}"
  private_ip_google_access = true
}

//resource "google_compute_firewall" "firewalli-int" {
//  name    = "${terraform.workspace}-firewall-internal"
//  network = "${google_compute_network.primary.name}"
//
//  allow {
//    protocol = "icmp"
//  }
//
//  allow {
//    protocol = "tcp"
//  }
//
//  allow {
//    protocol = "udp"
//  }
//
//  source_ranges = ["${var.ip_cidr_range}"]
//}
//
//resource "google_compute_firewall" "firewalli-ext" {
//  name    = "${terraform.workspace}-firewall-external"
//  network = "${google_compute_network.primary.name}"
//
//  allow {
//    protocol = "icmp"
//  }
//
//  allow {
//    protocol = "tcp"
//    ports    = ["22", "6443"]
//  }
//
//  source_ranges = ["0.0.0.0/0"]
//}
