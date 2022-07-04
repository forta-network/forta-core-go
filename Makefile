export GOBIN = $(shell pwd)/toolbin

SWAGGER = $(GOBIN)/swagger

tools:
	@rm -rf toolbin
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29.0

proto:
	protoc -I=protocol --go_out=protocol/. protocol/metrics.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/agent.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/publisher.proto
	protoc -I=protocol --go_out=protocol/. protocol/batch.proto

mocks:
	mockgen -source ethereum/client.go -destination ethereum/mocks/mock_client.go
	mockgen -source feeds/interfaces.go -destination feeds/mocks/mock_feeds.go
	mockgen -source ethereum/contract_backend.go -destination ethereum/mocks/mock_ethclient.go
	mockgen -source registry/client.go -destination registry/mocks/mock_client.go
	mockgen -source ipfs/client.go -destination ipfs/mocks/mock_client.go
	mockgen -source release/client.go -destination release/mocks/mock_client.go

lint:
	golangci-lint run ./...

test:
	go test -v -count=1 ./...

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
swagger:
	@rm -rf clients/webhook/swagger
	@$(SWAGGER) generate client \
		-f protocol/webhook/swagger.yml \
		-c clients/webhook/client \
		-m clients/webhook/client/models
