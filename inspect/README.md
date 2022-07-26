# forta-inspector-bot

# Inspection Results

## Scan API Inspection

### `scan-api.accessible`

Scan API is accessible and responsive

`1` for accessible, `-1` for not

### `scan-api.chain-id`

Chain ID of the JSON-RPC API

### `scan-api.module.eth`

`eth` module is supported

The scan API should either declare `eth` module in `rpc_modules` method response or it should return a valid response to an `eth_chainId` request

`1` for supported, `-1` for not

### `scan-api.module.net`

`net` module is supported

The scan API should either declare `net` module in `rpc_modules` method response or it should return a valid response to a `net_version` request

`1` for supported, `-1` for not

### `scan-api.block-by-number.hash`

Keccak256 hash of the block data retrieved from the scan API

## Trace API Inspection

### `trace-api.accessible`

Trace API is accessible and responsive

`1` for accessible, `-1` for not

### `trace-api.supported`

`trace_block` method is supported

`1` for supported, `-1` for not

### `trace-api.block-by-number.hash`

Keccak256 hash of the block data retrieved from the Trace API

### `trace-api.trace-block.hash`

Keccak256 hash of the block trace data retrieved from the Trace API

## Proxy API Inspection

### `proxy-api.accessible`

The proxy API served within the node is accessible and responsive

`1` for accessible, `-1` for not

### `proxy-api.chain-id`

Chain ID of the JSON-RPC API

### `proxy-api.module.web3`

`web3` module is supported

The proxy API should either declare `net` module in `rpc_modules` method response or it should return a valid response to a `net_version` request

`1` for supported, `-1` for not

### `proxy-api.module.eth`

`eth` module is supported

The proxy API should either declare `eth` module in `rpc_modules` method or it should return a valid response to an `eth_chainId` request

`1` for supported, `-1` for not

### `proxy-api.module.net`

`net` module is supported

The proxy API should either declare `net` module in `rpc_modules` method or it should return a valid response to a `net_version` request

`1` for supported, `-1` for not

### `proxy-api.block-by-number.hash`

Keccak256 hash of the block data retrieved from the proxy API.

### `proxy-api.history-support`

The earliest block height that proxy supports. Lower is better.

## Network Inspection

### `network.access.outbound`

Can access https://api.forta.network

`1` for ok, `-1` for not accessible

### `network.speed.download`

Result of a download performance benchmark

Download speed in `mbps`

### `network.speed.upload`

Result of an upload performance benchmark

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

The time elapsed during the calculation of the first 1e5 prime numbers

Average time in nanoseconds

### `resources.cpu.usage`

Current average CPU usage percentage

## Inspection Validator

Crosschecks outputted `*.block-by-number.hash` and `trace-api.trace-block.hash` data in inspection results with reference json-rpc providers

## Example Output

```json
{
  "inputs": {
    "blockNumber": 123735915,
    "scanApiUrl": "https://cloudflare-eth.com",
    "proxyApiUrl": "https://cloudflare-eth.com",
    "traceApiUrl": "https://cloudflare-eth.com",
    "checkTrace": false
  },
  "metadata": {
    "proxy.block-by-number.hash": "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65",
    "scan-api.block-by-number.hash": "3abe2f22edf2b463cbc13343a947be9ebbf8c16c2b50b2b90e10a199a2344f65"
  },
  "indicators": {
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
