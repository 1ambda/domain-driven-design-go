TAG 		= "Makefile"

DOCKER 		= docker
MYSQLCLIENT = mycli
PIP 		= pip

VCS				= github.com
REPOSITORY		= 1ambda/domain-driven-design-go
MODULE_GATEWAY	= service-gateway
MODULE_FRONTEND = service-frontend

.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"
	@ $(PIP) install -U mycli

.PHONY: compose
compose:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running docker-compose"
	@ docker stop $(docker ps -a -q) || true
	@ docker rm -f $(docker ps -a -q) || true
	@ docker volume rm $(docker volume ls -f dangling=true -q) || true
	@ docker-compose -f docker-compose.storage.yml rm -fsv || true
	@ docker-compose -f docker-compose.storage.yml up

.PHONY: compose.all
compose.all:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running docker-compose"
	@ docker stop $(docker ps -a -q) || true
	@ docker rm -f $(docker ps -a -q) || true
	@ docker volume rm $(docker volume ls -f dangling=true -q) || true
	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml build
	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml rm -fsv || true
	@ GATEWAY_SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE_GATEWAY)" docker-compose -f docker-compose.storage.yml -f docker-compose.application.yml up

.PHONY: mycli
mycli:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Connecting to mysql"
	@ $(MYSQLCLIENT) -u root -h localhost application -p root

