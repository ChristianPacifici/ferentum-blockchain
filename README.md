# Ferentum Blockchain in Go

[![Go Version](https://img.shields.io/badge/go-1.20+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/ChristianPacifici/ferentum-blockchain)


A minimal blockchain implementation in Go, built for learning and experimentation.

---

## 📌 Overview

This project is a simple blockchain implementation written in Go. It demonstrates the core concepts of blockchain technology:
- **Blocks** with data, timestamps, and hashes.
- **Chaining** blocks using cryptographic hashes.
- **Immutability** and tamper detection.

---

## 🛠️ Features

- Create a blockchain with a genesis block.
- Add new blocks to the chain.
- Validate the integrity of the blockchain.
- Basic proof-of-work (mining) simulation.

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

2. Run the project
   ```shell
      go run main.go
   ```