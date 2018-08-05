TAG 		= "Makefile"

DOCKER 		= docker
MYSQLCLIENT = mycli
PIP 		= pip

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
	@ docker-compose rm -fsv || true
	@ docker-compose up

.PHONY: mycli
mycli:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Connecting to mysql"
	@ $(MYSQLCLIENT) -u root -h localhost application -p root


.PHONY: app-gateway
mycli:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running GATEWAY"

.PHONY: app-frontend
mycli:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running FRONTEND"
