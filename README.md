<!--
parent:
  order: false
-->

<div align="center">
  <h1> XOS Node </h1>
</div>

<div align="center">
  <a href="https://github.com/xos-labs/node/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/xos-labs/node.svg" />
  </a>
  <a href="https://github.com/xos-labs/node/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/xos-labs/node.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/xos-labs/node">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/xos-labs/node"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/xosnetwork">
    <img alt="Discord" src="https://img.shields.io/discord/809048090249134080.svg" />
  </a>
  <a href="https://x.com/xos_labs">
    <img alt="Twitter Follow XOS Labs" src="https://img.shields.io/twitter/follow/xos_labs"/>
  </a>
</div>

## About

XOS Network is a high-performance blockchain platform that combines the best features of Cosmos SDK and Ethereum Virtual Machine (EVM). It provides a robust infrastructure for building decentralized applications with enhanced scalability and interoperability.

## Quick Start

### Prerequisites

- [Golang 1.23](https://go.dev/doc/install)
- [jq](https://stedolan.github.io/jq/download/)
- [Make](https://www.gnu.org/software/make/)
- [GCC](https://gcc.gnu.org/install/)

### Building from Source

1. Clone the repository and checkout the latest release:
```bash
git clone https://github.com/xos-labs/node.git
```

2. Build the binaries:
```bash
cd node
make install
```

This will install the `xosd` binaries in your `$GOPATH/bin` directory.

3. Verify the installation:
```bash
xosd version
which xosd
```

## Documentation

For detailed documentation, please visit [docs.x.ink](https://docs.x.ink/).

## Community

Join our community to stay updated and get support:

- [XOS Twitter](https://x.com/xos_labs)
- [XOS Discord](https://discord.gg/xosnetwork)