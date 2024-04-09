export GOBIN = $(shell pwd)/toolbin

LINT = $(GOBIN)/golangci-lint
FORMAT = $(GOBIN)/goimports

SWAGGER = $(GOBIN)/swagger
PROTOC_GEN_GO = $(GOBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC = $(GOBIN)/protoc-gen-go-grpc
GOMERGETYPES = $(GOBIN)/gomergetypes
MOCKGEN = $(GOBIN)/mockgen

PROTOC = ./toolbin/protoc --plugin=protoc-gen-go=$(PROTOC_GEN_GO) --plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC)

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

	@go install github.com/golang/mock/mockgen@5b455625bd2c8ffbcc0de6a0873f864ba3820904

	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0

	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

	@go install github.com/forta-network/go-merge-types/cmd/gomergetypes@ddc3807

	@wget https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip
	@mv protoc-21.12-linux-x86_64.zip toolbin
	@unzip toolbin/protoc-21* -d toolbin/protoc-bin
	@mv toolbin/protoc-bin/bin/protoc toolbin

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
	$(PROTOC) -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/dispatcher.proto

.PHONY: mocks
mocks:
	$(MOCKGEN) -source ethereum/client.go -destination ethereum/mocks/mock_client.go
	$(MOCKGEN) -source feeds/interfaces.go -destination feeds/mocks/mock_feeds.go
	$(MOCKGEN) -source ethereum/contract_backend.go -destination ethereum/mocks/mock_ethclient.go
	$(MOCKGEN) -source registry/client.go -destination registry/mocks/mock_client.go
	$(MOCKGEN) -source registry/version.go -destination registry/mocks/mock_version.go
	$(MOCKGEN) -source domain/registry/regmsg/regmsg.go -destination domain/registry/regmsg/mocks/mock_regmsg.go
	$(MOCKGEN) -source ipfs/client.go -destination ipfs/mocks/mock_client.go
	$(MOCKGEN) -source release/client.go -destination release/mocks/mock_client.go
	$(MOCKGEN) -source domain/ethereum.go -destination domain/mocks/mock_ethereum.go
	$(MOCKGEN) -source manifest/client.go -destination manifest/mocks/mock_client.go
	$(MOCKGEN) -source clients/graphql/client.go -destination clients/mocks/mock_graphql_client.go
	$(MOCKGEN) -source utils/ethutils/iterator.go -destination utils/ethutils/mocks/mock_iterator.go

.PHONY: test
test:
	go test -v -count=1 -short -coverprofile=coverage.out ./...

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
	$(GOMERGETYPES) --config contracts/merged/contract_agent_registry/transactor.yml

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

	abigen --out "contracts/external/contract_public_lock/public_lock.go" \
		--pkg "contract_public_lock" \
		--type "PublicLock" \
		--abi "contracts/external/contract_public_lock/PublicLock.json"


.PHONY: pull-contracts
pull-contracts:
	./scripts/pull-contracts.sh forta-contracts 59ae488

.PHONY: swagger
swagger: require-tools
	@rm -rf clients/webhook/swagger
	@$(SWAGGER) generate client \
		-f protocol/webhook/swagger.yml \
		-c clients/webhook/client \
		-m clients/webhook/client/models
