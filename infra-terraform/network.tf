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
