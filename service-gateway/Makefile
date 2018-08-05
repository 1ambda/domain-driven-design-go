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
LEFTPAREN:=(
RIGHTPAREN:=)
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./pkg/*" -not -path "./internal/mock*")
GODIRS = $(shell go list -f '{{.Dir}}' ./...  | grep -vFf <$(LEFTPAREN)go list -f '{{.Dir}}' ./vendor/...$(RIGHTPAREN) | grep -vFf <$(LEFTPAREN)go list -f '{{.Dir}}' ./pkg/...$(RIGHTPAREN))

GIT_COMMIT		= $(shell git rev-parse HEAD)
GIT_DIRTY		= $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
VERSION 		= $(shell cat ./VERSION)

APP				= server

.PHONY: install.tool
install.tool:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing prerequisites"

	@ echo "\t golang:dep \t\t\t(dependency)"
	@ $(GOGET) github.com/golang/dep/cmd/dep

	@ echo "\t golang:go-swagger \t\t(build)"
	@ $(GOGET) github.com/go-swagger/go-swagger/cmd/swagger

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

.PHONY: install.dependency
install.dependency:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Installing dependencies"
	@ $(GODEP) ensure
	@ echo ""

.PHONY: install
install: install.tool install.dependency

SWAGGER_SPEC 	= "../schema-swagger/gateway-rest.yaml"
.PHONY: swagger.version
swagger.version:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Updating swagger spec version: $(VERSION)"
	@VERSION=$(VERSION) SWAGGER_FILE=$(SWAGGER_SPEC) ../script/update-swagger-version.sh
	@ echo ""

.PHONY: swagger.server
swagger-server: swagger.version
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Cleaning swagger files"
	@ rm -rf pkg/generated/swagger/* || true
	@ mkdir -p pkg/generated/swagger || true
	@ echo ""

	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Validating swagger spec"
	@ swagger validate $(SWAGGER_SPEC)
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Generating swagger code"
	@ swagger generate server --spec=$(SWAGGER_SPEC) --exclude-main \
		 --target=pkg/generated/swagger --model-package=swagmodel \
		 --server-package=swagserver --api-package=swagapi
	@ echo ""

.PHONY: swagger
swagger: swagger.version swagger.server

.PHONY: build
build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Building $(CMD_DIR)/$(APP)/$(MAIN)"
	@ $(eval LDFLAG := $(shell $(GOVVV) -flags -pkg $(shell go list ./internal/config)))
	@ GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -race -i -ldflags="$(LDFLAG)" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN)
	@ echo ""

.PHONY: docker.build
docker.build:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Creating docker image: $(MODULE)"

	@ rm -rf $(DOCKER_BUILD_SH) || true
	@ $(eval LDFLAG := $(shell $(GOVVV) -flags -pkg $(shell go list ./internal/config)))
	@ echo CGO_ENABLED=1 GOOS=$(DOCKER_OS) GOARCH=$(DOCKER_ARCH) $(GOBUILD) -i -ldflags=\"$(LDFLAG)\" -o $(BIN_DIR)/$(APP) -v $(CMD_DIR)/$(APP)/$(MAIN) > $(DOCKER_BUILD_SH)
	@ chmod +x $(DOCKER_BUILD_SH)

	$(DOCKER) build --build-arg SRC_ROOT="$(VCS)/$(REPOSITORY)/$(MODULE)" -t "$(VCS)/$(REPOSITORY)/$(MODULE):$(VERSION)" .


.PHONY: run
run:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Running app: $(APP)"
	@ $(BIN_DIR)/$(APP)

.PHONY: clean
clean:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Clean"
	@rm -rf ./bin
	@ echo ""


.PHONY: test
test:
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Test"
	@ ginkgo -progress -race -cover -r -randomizeAllSpecs -randomizeSuites
	@ echo ""

.PHONY: coverage.convert
coverage.convert: test
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Coverage Convert"
	@ ../script/ginkgo-coverage.sh
	@ echo ""

.PHONY: coverage.push
coverage.push: test
	@ echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Coverage Push"
	@ ../script/ginkgo-coverage.sh --codecov
	@ echo ""

.PHONY: coverage
coverage: coverage.push

.PHONY: check.lint
check.lint:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u +'%H:%M:%S')) - Check Lint"
	@$(GOLINT) --vendor --errors ./...  \
		--exclude=vendor --exclude=pkg --exclude=internal/*mock_ \
		--enable=unparam --enable=nakedret --enable=safesql \
		--deadline=600s

.PHONY: check.format
check.format:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Check Format"
	@$(GOFMT) -l -w $(GOFILES)

.PHONY: check.simplify
check.simplify:
	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Check Simplify"
	@gofmt -s -l -w $(GOFILES)

.PHONY: check
check: check.format check.simplify check.lint


MOCKGEN			= mockgen
MOCK_PREFIX		= mock
MOCK_PKG_DIRS	:= "internal/domain/product" "internal/domain/user" "internal/domain/order"

$(MOCK_PKG_DIRS):
	@$(eval TARGET := $@)

	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Cleaning mock files under $(TARGET)"
	@rm -rf $(TARGET)/$(MOCK_PREFIX)*

	@echo "[$(TAG)] ($(shell TZ=UTC date -u '+%H:%M:%S')) - Generating mock files for $(TARGET)"
	@PKG_DIR=$(TARGET) ../script/generate-gomock.sh
	@echo ""

.PHONY: mock.generate
mock.generate: $(MOCK_PKG_DIRS)


.PHONY: godoc
godoc:
	godoc -http=:8080 -index


.PHONY: release
release: install clean swagger build
