SWEEP?=us-east-1,us-west-2
TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=stcloud

default: build

build: fmtcheck
	go install ./...

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4


fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

# Currently required by tf-deploy compile
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

depscheck:
	@echo "==> Checking source code with go mod tidy..."
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@echo "==> Checking source code with go mod vendor..."
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)


lint:
	@echo "==> Checking source code against linters..."
	@golangci-lint run ./$(PKG_NAME)/...
	@tfproviderlint \
		-c 1 \
		-AT001 \
		-AT002 \
		-S001 \
		-S002 \
		-S003 \
		-S004 \
		-S005 \
		-S007 \
		-S008 \
		-S009 \
		-S010 \
		-S011 \
		-S012 \
		-S013 \
		-S014 \
		-S015 \
		-S016 \
		-S017 \
		-S019 \
		./$(PKG_NAME)

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build test testacc fmt fmtcheck lint tools test-compile depscheck
