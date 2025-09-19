# Ferentum Blockchain

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/ChristianPacifici/ferentum-blockchain)

A simple, persistent blockchain implementation with a CLI interface, written in Go.

---

## 📌 Overview

**Ferentum Blockchain** is a minimal blockchain project designed for learning and experimentation. It includes:
- **Blockchain core** with blocks, hashing, and proof-of-work.
- **CLI interface** for adding blocks, printing the chain, and validation.
- **Persistent storage** using Go's `gob` encoding.

---

## 🛠 Features

- Create and manage a blockchain with a genesis block.
- Add new blocks with proof-of-work.
- Validate blockchain integrity.
- Persistent storage in `ferentum-blockchain.dat`.

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (1.20 or later)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ChristianPacifici/ferentum-blockchain.git
   cd ferentum-blockchain
   ```

2. Build the CLI:
   ```bash
   go build -o ferentum-blockchain ./cmd/cli
   ```

---

## 📂 Project Structure

```
ferentum-blockchain/
├── go.mod
├── go.sum
├── main.go
├── blockchain/
│   ├── blockchain.go
│   └── pow.go
└── cmd/
    └── cli/
        └── cli.go
```

---

## 💻 Usage

### Build the CLI

```bash
go build -o ferentum-blockchain ./cmd/cli
```

### Install the CLI (Optional)

```bash
go install tech.pacifici/blockchain/cmd/cli
```

---

## 📜 Commands

### Add a Block

```bash
./ferentum-blockchain add "My Data"
```
Adds a new block with the specified data to the blockchain.

---

### Print the Blockchain

```bash
./ferentum-blockchain print
```
Prints all blocks in the blockchain, including their index, timestamp, data, previous hash, hash, and nonce.

**Example Output:**
```
Index: 0
Timestamp: [current time]
Data: Genesis Block
PrevHash:
Hash: [hash value]
Nonce: [nonce value]
---
Index: 1
Timestamp: [current time]
Data: First Block
PrevHash: [hash of genesis block]
Hash: [hash value]
Nonce: [nonce value]
---
```

---

### Validate the Blockchain

```bash
./ferentum-blockchain validate
```
Validates the integrity of the blockchain.

**Example Output:**
```
Is blockchain valid? true
```

---

## 🔧 Example Workflow

1. **Initialize the Blockchain**
   Run any command to create the genesis block and `ferentum-blockchain.dat`:
   ```bash
   ./ferentum-blockchain print
   ```

2. **Add Blocks**
   ```bash
   ./ferentum-blockchain add "First Block"
   ./ferentum-blockchain add "Second Block"
   ```

3. **Print the Blockchain**
   ```bash
   ./ferentum-blockchain print
   ```

4. **Validate the Blockchain**
   ```bash
   ./ferentum-blockchain validate
   ```

---

## 🔧 Extending the Project

- **Add More Commands**: Implement commands like `reset`, `info`, or `mine`.
- **Improve Error Handling**: Add user-friendly error messages.
- **Add Networking**: Implement peer-to-peer communication.
- **Add Transactions**: Extend the blockchain to support transactions.
- **Build some APIS**: Add some go APIs.

---

## 🤝 Contributing

Contributions are welcome! Open an issue or submit a pull request.

---

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).