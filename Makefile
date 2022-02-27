proto:
	protoc -I=protocol --go_out=protocol/. protocol/metrics.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/agent.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/publisher.proto
	protoc -I=protocol --go_out=protocol/. protocol/batch.proto

mocks:
	mockgen -source ethereum/client.go -destination ethereum/mocks/mock_client.go
	mockgen -source feeds/interfaces.go -destination feeds/mocks/mock_feeds.go
	mockgen -source ethereum/contract_backend.go -destination ethereum/mocks/mock_ethclient.go
	mockgen -source clients/interfaces.go -destination clients/mocks/mock_clients.go

test:
	go test -v -count=1 ./...

abigen:
	./scripts/abigen.sh alerts
	./scripts/abigen.sh agent_registry
	./scripts/abigen.sh dispatch
	./scripts/abigen.sh scanner_node_version
	./scripts/abigen.sh scanner_registry "--alias _register=underscoreRegister"
