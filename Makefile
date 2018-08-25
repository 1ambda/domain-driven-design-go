TAG		= "Makefile"

DOCKER		= docker
MYSQLCLIENT	= mycli
PIP			= pip

VCS				= github.com
REPOSITORY		= 1ambda/domain-driven-design-go
MODULE_GATEWAY	= service-gateway
MODULE_FRONTEND = service-frontend

.PHONY: default
default:
	@ mmake help

# Install required tools for development
.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Installing prerequisites"
	@ echo "-----------------------------------------\n"

	@ $(PIP) install -U mycli

# Run docker-compose for storages
.PHONY: compose
compose:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Running docker-compose (storage only)"
	@ echo "-----------------------------------------\n"

	@ docker-compose -f docker-compose.storage.yml rm -fsv || true
	@ docker-compose -f docker-compose.storage.yml up

# Run docker-compose for storages + applications
.PHONY: compose.all
compose.all:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Running docker-compose (storage + application)"
	@ echo "-----------------------------------------\n"

	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml build
	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml rm -fsv || true
	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml up

# Clean docker resources (image, volume, container and network)
.PHONY: compose.clean
compose.clean:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Starting: Cleaning docker resources"
	@ echo "-----------------------------------------\n"

	@ docker stop `docker ps -a -q` || true
	@ docker rm -f `docker ps -a -q` || true
	@ docker rmi -f `docker images --quiet --filter "dangling=true"` || true
	@ docker volume rm `docker volume ls -f dangling=true -q` || true
	@ docker network rm `docker network ls -q` || true

	@ echo ""
	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Cleaning docker resources"

# Connect to docker-composed MySQL through mycli
.PHONY: mycli
mycli:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Connecting to MySQL through mycli"
	@ echo "-----------------------------------------\n"

	@ $(MYSQLCLIENT) -u root -h localhost application -p root

