TAG 		= "Makefile"

VCS			= github.com
REPOSITORY	= 1ambda/domain-driven-design-go
MODULE		= service-frontend

DOCKER	= docker
NPM		= npm

ENV_FILE = .env
ENDPOINT_GATEWAY := $(or ${ENDPOINT_GATEWAY},${ENDPOINT_GATEWAY},http://localhost:30001/api)

SWAGGER_CLIENT_DIR      = "./src/generated/swagger"
SWAGGER_FILE       = "../schema-swagger/gateway-rest.yaml"

GIT_COMMIT		= $(shell git rev-parse HEAD)
GIT_DIRTY		= $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
VERSION 		= $(shell cat ./VERSION)

.PHONY: prepare
prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"
	@

	@ echo "\t brew:swagger-codegen \t\t\t(tool)"
	@ brew install swagger-codegen

.PHONY: install
install:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing dependencies"
	@ rm -rf node_modules
	@ $(NPM) install

.PHONY: build
build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Building frontend distribution"
	@ $(NPM) run build

.PHONY: docker.prepare
docker.prepare:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Preparing docker build: $(MODULE)"

	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - GATEWAY_ENDPOINT: $(GATEWAY_ENDPOINT)"
	@ echo VUE_APP_GATEWAY_ENDPOINT=$(ENDPOINT_GATEWAY) > $(ENV_FILE)
	@ echo VUE_APP_DEBUG=false >> $(ENV_FILE)
	@ cat $(ENV_FILE)


.PHONY: docker.build
docker.build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Creating docker image: $(MODULE)"

	$(DOCKER) build -t "$(VCS)/$(REPOSITORY)/$(MODULE):$(VERSION)" .

.PHONY: docker
docker: docker.prepare docker.build


.PHONY: run
run:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running frontend application"
	@ $(NPM) run serve

.PHONY: swagger-client
swagger-client:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Generating swagger client files"

	@ rm -rf $(SWAGGER_CLIENT_DIR)		|| true
	@ mkdir -p $(SWAGGER_CLIENT_DIR)	|| true
	@ swagger-codegen generate -i $(SWAGGER_FILE) -l typescript-fetch -o $(SWAGGER_CLIENT_DIR)

.PHONY: swagger
swagger: swagger-client

