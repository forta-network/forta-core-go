export GOBIN = $(shell pwd)/toolbin

LINT = $(GOBIN)/golangci-lint
FORMAT = $(GOBIN)/goimports

SWAGGER = $(GOBIN)/swagger
PROTOC_GEN_GO = $(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GOBIN)/protoc-gen-go-grpc
GOMERGETYPES = $(GOBIN)/gomergetypes

PROTOC = protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC)

.PHONY: require-tools
require-tools: tools
	@echo 'Checking installed tools...'
	@file $(LINT) > /dev/null
	@file $(FORMAT) > /dev/null

	@file $(SWAGGER) > /dev/null

	@file $(PROTOC_GEN_GO) > /dev/null
	@file $(PROTOC_GEN_GO_GRPC) > /dev/null

	@file $(GOMERGETYPES) > /dev/null

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

	@go install github.com/forta-network/go-merge-types/cmd/gomergetypes@89a84e7

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
	mockgen -source registry/version.go -destination registry/mocks/mock_version.go
	mockgen -source ipfs/client.go -destination ipfs/mocks/mock_client.go
	mockgen -source release/client.go -destination release/mocks/mock_client.go
	mockgen -source domain/ethereum.go -destination domain/mocks/mock_ethereum.go

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: abigen
abigen: pull-contracts
	rm -rf contracts/generated
	./scripts/abigen.sh forta-contracts components/staking/FortaStaking.sol FortaStaking FortaStaking forta_staking_0_1_2
	./scripts/abigen.sh forta-contracts components/_old/staking/FortaStaking_0_1_1.sol FortaStaking_0_1_1 FortaStaking forta_staking_0_1_1
	$(GOMERGETYPES) --config contracts/merged/contract_forta_staking/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_forta_staking/filterer.yml

	./scripts/abigen.sh forta-contracts components/staking/allocation/StakeAllocator.sol StakeAllocator StakeAllocator stake_allocator_0_1_0
	$(GOMERGETYPES) --config contracts/merged/contract_stake_allocator/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_stake_allocator/filterer.yml

	./scripts/abigen.sh forta-contracts components/staking/rewards/RewardsDistributor.sol RewardsDistributor RewardsDistributor rewards_distributor_0_1_0
	$(GOMERGETYPES) --config contracts/merged/contract_rewards_distributor/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_rewards_distributor/filterer.yml

	./scripts/abigen.sh forta-contracts components/agents/AgentRegistry.sol AgentRegistry AgentRegistry agent_registry_0_1_6
	./scripts/abigen.sh forta-contracts components/_old/agents/AgentRegistry_0_1_4.sol AgentRegistry_0_1_4 AgentRegistry agent_registry_0_1_4
	$(GOMERGETYPES) --config contracts/merged/contract_agent_registry/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_agent_registry/filterer.yml

	./scripts/abigen.sh forta-contracts components/scanners/ScannerNodeVersion.sol ScannerNodeVersion ScannerNodeVersion scanner_node_version_0_1_1
	./scripts/abigen.sh forta-contracts components/_old/scanners/ScannerNodeVersion_0_1_0.sol ScannerNodeVersion_0_1_0 ScannerNodeVersion scanner_node_version_0_1_0
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_node_version/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_node_version/filterer.yml

	./scripts/abigen.sh forta-contracts components/scanner_pools/ScannerPoolRegistry.sol ScannerPoolRegistry ScannerPoolRegistry scanner_pool_registry_0_1_0 "--alias _register=underscoreRegister"
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_pool_registry/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_pool_registry/filterer.yml

	./scripts/abigen.sh forta-contracts token/Forta.sol Forta Forta forta_0_2_0
	$(GOMERGETYPES) --config contracts/merged/contract_forta/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_forta/filterer.yml

	./scripts/abigen.sh forta-contracts components/scanners/ScannerRegistry.sol ScannerRegistry ScannerRegistry scanner_registry_0_1_4 "--alias _register=underscoreRegister"
	./scripts/abigen.sh forta-contracts components/_old/scanners/ScannerRegistry_0_1_3.sol ScannerRegistry_0_1_3 ScannerRegistry scanner_registry_0_1_3 "--alias _register=underscoreRegister"
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_registry/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_scanner_registry/filterer.yml

	./scripts/abigen.sh forta-contracts components/dispatch/Dispatch.sol Dispatch Dispatch dispatch_0_1_5
	./scripts/abigen.sh forta-contracts components/_old/dispatch/Dispatch_0_1_4.sol Dispatch_0_1_4 Dispatch dispatch_0_1_4
	$(GOMERGETYPES) --config contracts/merged/contract_dispatch/caller.yml
	$(GOMERGETYPES) --config contracts/merged/contract_dispatch/filterer.yml


.PHONY: pull-contracts
pull-contracts:
	./scripts/pull-contracts.sh forta-contracts caner/include-old-in-generated-abis

.PHONY: swagger
swagger: require-tools
	@rm -rf clients/webhook/swagger
	@$(SWAGGER) generate client \
		-f protocol/webhook/swagger.yml \
		-c clients/webhook/client \
		-m clients/webhook/client/models
