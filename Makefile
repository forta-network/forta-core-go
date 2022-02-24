proto:
	protoc -I=protocol --go_out=protocol/. protocol/metrics.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/agent.proto
	protoc -I=protocol --go-grpc_out=protocol/. --go_out=protocol/. protocol/publisher.proto
	protoc -I=protocol --go_out=protocol/. protocol/batch.proto

mocks:
	mockgen -source ethereum/client.go -destination ethereum/mocks/mock_client.go
	mockgen -source clients/interfaces.go -destination clients/mocks/mock_clients.go
	mockgen -source feeds/interfaces.go -destination feeds/mocks/mock_feeds.go
	mockgen -source ethereum/contract_backend.go -destination ethereum/mocks/mock_ethclient.go

test:
	go test -v -count=1 ./...

abigen:
	abigen --abi ./contracts/agent_registry.json --out ./contracts/agent_registry.go --pkg contracts --type AgentRegistry
	abigen --abi ./contracts/dispatch.json --out ./contracts/dispatch.go --pkg contracts --type Dispatch
	abigen --abi ./contracts/alerts.json --out ./contracts/alerts.go --pkg contracts --type Alerts
	abigen --abi ./contracts/scanner_registry.json --out ./contracts/scanner_registry.go --pkg contracts --type ScannerRegistry
	abigen --abi ./contracts/scanner_node_version.json --out ./contracts/scanner_node_version.go --pkg contracts --type ScannerNodeVersion

