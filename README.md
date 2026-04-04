# Netutils Go

Concurrent network utility tool for checking URL health.

## Installation

```bash
go install github.com/user/netutils-go@latest
```

## Usage

```bash
# Basic check
netutils https://example.com https://google.com

# JSON output with custom timeout
netutils --json --timeout 5 https://example.com

# Limit concurrent workers
netutils -w 3 url1 url2 url3 url4
```

## Docker

```bash
docker build -t netutils .
docker run netutils https://example.com
```

## Development

```bash
make build
make test
make lint
```

## License

MIT