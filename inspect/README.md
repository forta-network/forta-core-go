# forta-inspector-bot

Inspects run environment capabilities and limits

Moved from [forta-network/forta-inspector-bot](forta-inspector-bot-repo)

# Inspection Details

## Scan API Inspection

### `scan-api.accessible`

Checks Scan APIs accessibility by establishing rpc connection and asking for current block height

`1` for accessible, `-1` for not

### `scan-api.module.eth`

The scan node should either declare `eth` module in `rpc_modules` method or it should return a valid response to
a `eth_chainId` query

`1` for supports, `-1` for not

### `scan-api.module.net`

The scan api should either declare `net` module in `rpc_modules` method or it should return a valid response to
a `net_version` query

`1` for supports, `-1` for not

### `scan-api.block-by-number.hash`

The hash of the block data retrieved from the scan API

## Trace API Inspection

### `trace-api.accessible`

Checks trace APIs accessibility by establishing rpc connection and asking for current its chain id.

`1` for accessible, `-1` for not

### `trace-api.supported`

Only necessary for Ethereum Mainnet.

Returns `1` by default on other chains

Checks trace APIs accessibility by establishing rpc connection and tracing the genesis block or the trace field in
provided gRPC message

`1` for supported, `-1` for not

### `trace-api.block-by-number.hash`

The hash of the block data retrieved from the Trace API

### `trace-api.trace-block.hash`

The hash of the block trace data retrieved from the Trace API

## Proxy Inspection

### `proxy.accessible`

Checks json-rpc proxy accessibility by establishing rpc connection and asking for current its chain id.

`1` for accessible, `-1` for not

### `proxy.chain-id`

Chain ID of the provided json-rpc

### `proxy.module.web3`

The json-rpc proxy should either declare `net` module in `rpc_modules` method or it should return a valid response to
a `net_version` query

`1` for supports, `-1` for not

### `proxy.module.eth`

The json-rpc proxy should either declare `eth` module in `rpc_modules` method or it should return a valid response to
a `eth_chainId` query

`1` for supports, `-1` for not

### `proxy.module.net`

The json-rpc proxy should either declare `net` module in `rpc_modules` method or it should return a valid response to
a `net_version` query

`1` for supports, `-1` for not

### `proxy.block-by-number.hash`

The hash of the block data retrieved from the json-rpc proxy.

### `proxy.history-support`

Returns the earliest block height that proxy supports. Lower is better.

## Network Inspection

### `network.access.outbound`

Tries to access https://api.forta.network

`1` for ok, `-1` for not accessible

### `network.speed.download`

Runs a download performance benchmark

Download speed in `mbps`

### `network.speed.upload`

Runs an upload performance benchmark

Upload speed in `mbps`

## System Resources Inspection

### `resources.memory.total`

Total memory in bytes

### `resources.memory.available`

Current available memory in bytes

### `resources.storage.total`

Total storage in bytes

### `resources.storage.available`

Current available storage in bytes

### `resources.cpu.benchmark`

The Time container needs to calculate first 1e5 prime numbers on average

Average time needed in nanoseconds

### `resources.cpu.usage`

Current average CPU usage percentage

## Inspection Validator

Crosschecks submitted `*.block-by-number.hash` data in inspection results with reference json-rpc providers

## Example Submission

```json
{
  "Inputs": {
    "blockNumber": 0,
    "scanApiUrl": "https://cloudflare-eth.com",
    "proxyApiUrl": "https://cloudflare-eth.com",
    "traceApiUrl": "https://cloudflare-eth.com",
    "checkTrace": false
  },
  "Metadata": {
    "proxy.block-by-number.hash": "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65",
    "scan-api.block-by-number.hash": "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65"
  },
  "Indicators": {
    "network.access.outbound": 1,
    "network.speed.download": 30.381701453262668,
    "network.speed.upload": 38.8611939336301,
    "proxy.accessible": 1,
    "proxy.chain-id": 1,
    "proxy.history-support": 5,
    "proxy.module.eth": 1,
    "proxy.module.net": 1,
    "proxy.module.web3": 1,
    "resources.cpu.benchmark": 1074533049,
    "resources.memory.available": 2090770432,
    "resources.memory.total": 16657076224,
    "resources.storage.available": 250357280768,
    "resources.storage.total": 10473631744,
    "scan-api.accessible": 1,
    "scan-api.chain-id": 1,
    "scan-api.module.eth": 1,
    "scan-api.module.net": 1,
    "trace-api.accessible": -1,
    "trace-api.supported": -1
  }
}
```

[forta-inspector-bot-repo]: https://github.com/forta-network/forta-inspector-bot/ 