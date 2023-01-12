export GOBIN = $(shell pwd)/toolbin

LINT = $(GOBIN)/golangci-lint
FORMAT = $(GOBIN)/goimports

SWAGGER = $(GOBIN)/swagger
PROTOC_GEN_GO = $(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GOBIN)/protoc-gen-go-grpc

PROTOC = protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC)

.PHONY: require-tools
require-tools: tools
	@echo 'Checking installed tools...'
	@file $(LINT) > /dev/null
	@file $(FORMAT) > /dev/null

	@file $(SWAGGER) > /dev/null

	@file $(PROTOC_GEN_GO) > /dev/null
	@file $(PROTOC_GEN_GO_GRPC) > /dev/null

	@echo "All tools found in $(GOBIN)!"

.PHONY: tools
tools:
	@echo 'Installing tools...'
	@rm -rf toolbin
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.0
	@go install golang.org/x/tools/cmd/goimports@v0.1.11

	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0

	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: fmt
fmt: require-tools
	@go mod tidy
	@$(FORMAT) -w $$(go list -f {{.Dir}} ./...)

.PHONY: lint
lint: require-tools fmt
	@$(LINT) run ./...

.PHONY: proto
proto: protogen fmt

.PHONY: protogen
protogen: require-tools
	$(PROTOC) -I=protocol --go_out=protocol/. protocol/bot_metrics.proto
	$(PROTOC) -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/alert.proto
	$(PROTOC) -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/agent.proto
	$(PROTOC) -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/publisher.proto
	$(PROTOC) -I=protocol --go_out=protocol/. protocol/batch.proto
	$(PROTOC) -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/storage.proto

.PHONY: mocks
mocks:
	mockgen -source ethereum/client.go -destination ethereum/mocks/mock_client.go
	mockgen -source feeds/interfaces.go -destination feeds/mocks/mock_feeds.go
	mockgen -source ethereum/contract_backend.go -destination ethereum/mocks/mock_ethclient.go
	mockgen -source registry/client.go -destination registry/mocks/mock_client.go
	mockgen -source ipfs/client.go -destination ipfs/mocks/mock_client.go
	mockgen -source release/client.go -destination release/mocks/mock_client.go
	mockgen -source domain/ethereum.go -destination domain/mocks/mock_ethereum.go

.PHONY: test
test:
	go test -v -count=1 -coverprofile=coverage.out ./...

.PHONY: coverage
coverage:
	go tool cover -func=coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'

.PHONY: coverage-func
coverage-func:
	go tool cover -func=coverage.out

.PHONY: coverage-html
coverage-html:
	go tool cover -html=coverage.out -o=coverage.html

.PHONY: abigen
abigen: pull-contracts
	rm -rf contracts
	./scripts/abigen.sh forta-contracts components/staking/FortaStaking.sol FortaStaking forta_staking
	./scripts/abigen.sh forta-contracts components/agents/AgentRegistry.sol AgentRegistry agent_registry
	./scripts/abigen.sh forta-contracts components/dispatch/Dispatch.sol Dispatch dispatch
	./scripts/abigen.sh forta-contracts components/scanners/ScannerNodeVersion.sol ScannerNodeVersion scanner_node_version
	./scripts/abigen.sh forta-contracts components/scanners/ScannerRegistry.sol ScannerRegistry scanner_registry "--alias _register=underscoreRegister"
	./scripts/abigen.sh forta-contracts components/access/AccessManager.sol AccessManager access_manager
	./scripts/abigen.sh forta-contracts components/router/Router.sol Router router
	./scripts/abigen.sh forta-contracts token/Forta.sol Forta forta

.PHONY: pull-contracts
pull-contracts:
	./scripts/pull-contracts.sh forta-contracts master

.PHONY: swagger
swagger: require-tools
	@rm -rf clients/webhook/swagger
	@$(SWAGGER) generate client \
		-f protocol/webhook/swagger.yml \
		-c clients/webhook/client \
		-m clients/webhook/client/models
