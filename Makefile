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

test:
	go test -v -count=1 ./...

abigen: pull-contracts
	./scripts/abigen.sh forta-token components/agents agent_registry
	./scripts/abigen.sh forta-token components/dispatch dispatch
	./scripts/abigen.sh forta-token components/scanners scanner_node_version
	./scripts/abigen.sh forta-token components/scanners scanner_registry "--alias _register=underscoreRegister"

.PHONY: pull-contracts
pull-contracts:
	./scripts/pull-contracts.sh forta-token master
