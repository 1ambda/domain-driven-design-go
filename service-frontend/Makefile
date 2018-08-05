TAG 		= "Makefile"

VCS			= github.com
REPOSITORY	= 1ambda/domain-driven-design-go
MODULE		= service-frontend

VERSION 		= $(shell cat ./VERSION)

DOCKER			= docker

NPM		= npm

SWAGGER_CLIENT_DIR      = "./src/generated/swagger"
SWAGGER_FILE       = "../schema-swagger/gateway-rest.yaml"

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

.PHONY: docker.build
docker.build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Creating docker image: $(MODULE)"

	$(DOCKER) build --build-arg SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE)" -t "$(VCS)/$(REPOSITORY)/$(MODULE):$(VERSION)" .

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

