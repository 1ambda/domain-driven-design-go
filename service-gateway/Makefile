GOARCH  = amd64
GOOS    = $(OS)

ifeq ($(GOOS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
endif

TAG 		= "Makefile"

GOCMD		= go
GODEP		= dep
GOVVV		= govvv
GOLINT		= gometalinter
GOBUILD		= $(GOCMD) build
GOFMT		= $(GOCMD)fmt
GOGET		= go get -u

VCS			= github.com
REPOSITORY	= 1ambda/domain-driven-design-go
MODULE		= service-gateway

DOCKER			= docker
DOCKER_BUILD_SH = .BUILD.sh
DOCKER_OS		= linux
DOCKER_ARCH		= amd64

MAIN		= main.go
BIN_DIR		= bin
CMD_DIR		= cmd
VENDOR_DIR	= vendor

# Pure go sources (not vendored and not generated)
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./pkg/*" -not -path "./internal/mock*")

GIT_COMMIT		= $(shell git rev-parse HEAD)
GIT_DIRTY		= $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
VERSION 		= $(shell cat ./VERSION)

APP				= server

.PHONY: default
default:
	@ mmake help

# Install required tools for development
.PHONY: install.tool
install.tool:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Installing tools"

	@ echo "\t golang:dep \t\t\t(dependency)"
	@ $(GOGET) github.com/golang/dep/cmd/dep

	@ echo "\t golang:go-swagger \t\t(build)"
	@ $(GOGET) github.com/go-swagger/go-swagger/cmd/swagger

	@ echo "\t golang:mmake \t\t\t(build)"
	@ $(GOGET) github.com/tj/mmake/cmd/mmake

	@ echo "\t golang:govvv \t\t\t(build)"
	@ $(GOGET) github.com/ahmetb/govvv

	@ echo "\t golang:ginkgo \t\t\t(testing)"
	@ $(GOGET) github.com/onsi/ginkgo/ginkgo

	@ echo "\t golang:gomega \t\t\t(testing)"
	@ $(GOGET) github.com/onsi/gomega/...

	@ echo "\t golang:mockgen \t\t(testing)"
	@ $(GOGET) github.com/golang/mock/mockgen

	@ echo "\t golang:gometalinter \t\t(check)"
	@ $(GOGET) github.com/alecthomas/gometalinter
	@ $(GOLINT) --install

	@ echo ""

# Install golang dependencies using dep
.PHONY: install.dependency
install.dep:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Installing golang dependencies"
	@ $(GODEP) ensure

	@ echo ""

.PHONY: install
install: install.tool install.dep

SWAGGER_SPEC 	= "../schema-swagger/gateway-rest.yaml"
.PHONY: swagger.version
swagger.version:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Updating swagger version"
	@VERSION=$(VERSION) SWAGGER_FILE=$(SWAGGER_SPEC) ../script/update-swagger-version.sh
	@ echo ""

.PHONY: swagger.server
swagger.server: swagger.version
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Swagger"
	@ echo "-----------------------------------------\n"

	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Cleaning swagger files"

	@ rm -rf pkg/generated/swagger/* || true
	@ mkdir -p pkg/generated/swagger || true
	@ echo ""

	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Validating swagger schemas"
	@ swagger validate $(SWAGGER_SPEC)

	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Generating swagger files"
	@ swagger generate server --spec=$(SWAGGER_SPEC) --exclude-main \
		 --target=pkg/generated/swagger --model-package=swagmodel \
		 --server-package=swagserver --api-package=swagapi \
		 --with-flatten=full

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Swagger"
	@ echo ""

# Generate swagger models for server based on the schemas
.PHONY: swagger
swagger: swagger.version swagger.server


# Build application and write the binary into the `bin` directory
.PHONY: build
build:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Build"
	@ echo "-----------------------------------------\n"

	@ $(eval LDFLAG := $(shell $(GOVVV) -flags -pkg $(shell go list ./internal/config)))
	@ GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -race -i -ldflags="$(LDFLAG)" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN)
	@ echo 'GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -race -i -ldflags="$(LDFLAG)" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN)'

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Build"
	@ echo ""

.PHONY: docker.prepare
docker.prepare:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Preparing docker build"

	@ rm -rf $(DOCKER_BUILD_SH) || true
	@ $(eval LDFLAG := $(shell $(GOVVV) -flags -pkg $(shell go list ./internal/config)))
	@ echo CGO_ENABLED=0 GOOS=$(DOCKER_OS) GOARCH=$(DOCKER_ARCH) $(GOBUILD) -i -ldflags=\"$(LDFLAG)\" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN) > $(DOCKER_BUILD_SH)
	@ chmod +x $(DOCKER_BUILD_SH)

.PHONY: docker.build
docker.build:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Creating docker image $(MODULE)"
	@ echo "-----------------------------------------\n"

	$(DOCKER) build --build-arg SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE)" -t "$(VCS)/$(REPOSITORY)/$(MODULE):$(VERSION)" .

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Creating docker image $(MODULE)"
	@ echo ""

# Build docker image
.PHONY: docker
docker: docker.prepare docker.build

# Run the built application
.PHONY: run
run:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Running Application: $(APP)"
	@ echo "-----------------------------------------\n"
	@ $(BIN_DIR)/$(APP)

# Remove all outputs
.PHONY: clean
clean:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Clean"
	@ echo "-----------------------------------------\n"

	@ echo "rm -rf ./bin"
	@ rm -rf ./bin

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Clean"
	@ echo ""


# Run unit tests
.PHONY: test
test:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Unit Testing"
	@ echo "-----------------------------------------\n"

	@ ginkgo -progress -race -cover -r -randomizeAllSpecs -randomizeSuites

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Unit Testing"
	@ echo ""

.PHONY: coverage.convert
coverage.convert: test
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Coverage convert"
	@ ../script/ginkgo-coverage.sh
	@ echo ""

.PHONY: coverage.push
coverage.push: test
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Coverage push"
	@ ../script/ginkgo-coverage.sh --codecov
	@ echo ""

.PHONY: coverage
coverage: coverage.push

# Run gometaliter
.PHONY: check.lint
check.lint:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Check lint"
	@ echo "-----------------------------------------\n"

	@ echo "Executing gometalinter..."
	@$(GOLINT) --vendor --errors ./...  \
		--exclude=vendor --exclude=pkg --exclude=internal/*mock_ \
		--enable=unparam --enable=nakedret --enable=safesql \
		--deadline=600s


	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Check lint"
	@ echo ""

# Run gofmt
.PHONY: check.format
check.format:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Check format"
	@ echo "-----------------------------------------\n"

	@ echo "Executing gofmt..."
	@$(GOFMT) -l -w $(GOFILES)

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Check format"
	@ echo ""

# Run gofmt with `-s` option.
.PHONY: check.simplify
check.simplify:
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Started: Check simplify"
	@ echo "-----------------------------------------\n"

	@ echo "Executing gofmt..."
	@gofmt -s -l -w $(GOFILES)

	@ echo "\n-----------------------------------------"
	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Finished: Check simplify"
	@ echo ""

# Run all check targets (format, simplify and lint)
.PHONY: check
check: check.format check.simplify check.lint


MOCKGEN			= mockgen
MOCK_PREFIX		= mock
MOCK_PKG_DIRS	:= "internal/domain/product" "internal/domain/user" "internal/domain/order"

$(MOCK_PKG_DIRS):
	@$(eval TARGET := $@)

	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Cleaning mock files under $(TARGET)"
	@rm -rf $(TARGET)/$(MOCK_PREFIX)*

	@ echo "[$(TAG)] ($$(date -u '+%H:%M:%S')) - Generating mock files for $(TARGET)"
	@PKG_DIR=$(TARGET) ../script/generate-gomock.sh

	@echo ""

# Generate mock files
.PHONY: mock.generate
mock.generate: $(MOCK_PKG_DIRS)

# Run godoc on localhost:8000
.PHONY: godoc
godoc:
	godoc -http=:8000 -index


.PHONY: release
release: install clean swagger build
