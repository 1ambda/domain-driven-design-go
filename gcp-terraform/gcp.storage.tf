locals {
  database_instance_prefix = "${local.service}"

  database_version = "MYSQL_5_7"
  database_name = "g-street"
  database_charset   = "utf8mb4"
  database_collation = "utf8mb4_general_ci"

  database_instance_tier = "db-g1-small"
  database_master_instance = "g-street"

  database_disk_size        = "10"
  database_disk_type        = "PD_SSD"
  database_disk_autoresize  = true

  database_master_zone  = "${element(local.zones, 0)}"
  database_replica_zone = "${element(local.zones, 1)}"
}

resource "google_sql_database_instance" "master" {
  project              = "${data.google_project.current.id}"

  name                 = "${local.database_instance_prefix}-master"
  region               = "${local.region}"
  database_version     = "${local.database_version}"

  settings {
    tier                        = "${local.database_instance_tier}"

    disk_autoresize             = "${local.database_disk_autoresize}"
    disk_size                   = "${local.database_disk_size}"
    disk_type                   = "${local.database_disk_type}"

    availability_type           = "ZONAL"
    activation_policy           = "ALWAYS"

    replication_type            = "SYNCHRONOUS"

    ip_configuration {
      require_ssl  = false
      ipv4_enabled = true

      authorized_networks = [
        {
          value = "0.0.0.0/0"
        }
      ]
    }

    location_preference {
      zone = "${local.database_master_zone}"
    }

    backup_configuration {
      binary_log_enabled  = true
      enabled             = true
      start_time          = "02:30" # 2:30 AM
    }

    maintenance_window {
      day                 = 1 # Monday
      hour                = 2 # 2 AM
      update_track        = "stable"
    }
  }
}

resource "google_sql_database_instance" "replica" {
  project              = "${data.google_project.current.id}"

  name                 = "${local.database_instance_prefix}-replica"
  region               = "${local.region}"
  database_version     = "${local.database_version}"

  master_instance_name = "${google_sql_database_instance.master.name}"

  replica_configuration {
    # connect_retry_interval = "${lookup(var.replica, "retry_interval", "60")}"
    failover_target = true
  }

  settings {
    tier                        = "${local.database_instance_tier}"

    disk_autoresize             = "${local.database_disk_autoresize}"
    disk_size                   = "${local.database_disk_size}"
    disk_type                   = "${local.database_disk_type}"

    availability_type           = "ZONAL"
    activation_policy           = "ALWAYS"

    crash_safe_replication      = true

    location_preference {
      zone = "${local.database_replica_zone}"
    }

    maintenance_window {
      day          = 3 # Wednesday
      hour         = 2 # 2 AM
      update_track = "stable"
    }
  }
}

resource "google_sql_database" "default" {
  project   = "${data.google_project.current.id}"
  instance  = "${google_sql_database_instance.master.name}"

  name      = "${local.database_name}"

  charset   = "${local.database_charset}"
  collation = "${local.database_collation}"
}

variable "database_username" {
  default = ""
  description = "Username of the host to access the database"
}

variable "database_password" {
  default = ""
  description = "Password of the host to access the database"
}

resource "google_sql_user" "default" {
  project  = "${data.google_project.current.id}"

  name     = "${var.database_username}"
  password = "${var.database_password}"
  instance = "${google_sql_database_instance.master.name}"

  depends_on = [
    "google_sql_database_instance.master",
  ]

  lifecycle {
    ignore_changes = [
      "name",
      "password",
    ]
  }
}

