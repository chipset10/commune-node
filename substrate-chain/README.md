# Commune Node Engine

## Getting Started

### Build

Use the following command to build the node without launching it:

```sh
cargo build --release
```

### Embedded Docs

```sh
./target/release/node-template -h
```

### Single-Node Mode

```sh
./target/release/node-template --dev
```

To purge the development chain's state, run the following command:

```sh
./target/release/node-template purge-chain --dev
```

To start the development chain with detailed logging, run the following command:

```sh
RUST_BACKTRACE=1 ./target/release/node-template -ldebug --dev
```