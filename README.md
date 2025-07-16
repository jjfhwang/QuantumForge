# QuantumForge: On-Chain Intelligence, Zero-Knowledge Automation

QuantumForge is a decentralized, serverless, crypto-trading bot framework leveraging on-chain data aggregation and zk-SNARK powered position automation. It empowers users to develop and deploy sophisticated trading strategies directly on-chain, without relying on centralized exchanges or intermediaries. The framework prioritizes security, transparency, and decentralization, enabling trustless trading operations.

The primary goal of QuantumForge is to abstract away the complexities of interacting directly with blockchain data and smart contracts. It provides developers with a high-level API and a modular architecture to easily integrate custom trading logic. By aggregating data from multiple on-chain sources, including decentralized exchanges (DEXs) and oracles, QuantumForge provides accurate and real-time market insights. Furthermore, the integration of zk-SNARKs allows for the creation of automated trading positions that are verifiable on-chain but conceal sensitive strategy details, protecting intellectual property and preventing front-running. This promotes fair and efficient market participation for all users.

QuantumForge is designed to be serverless, meaning that the bot logic is executed within a decentralized environment. This eliminates the need for maintaining dedicated servers, reducing operational costs and enhancing scalability. The framework supports various execution environments, including decentralized compute platforms and secure enclave technologies. This flexibility allows users to choose the infrastructure that best suits their specific security and performance requirements. The modular design also enables easy extension and customization, allowing developers to tailor the framework to their unique trading needs and market conditions.

## Key Features

*   **Decentralized Data Aggregation:** The framework fetches and aggregates price data, liquidity information, and other relevant metrics from various on-chain sources like Uniswap, Chainlink, and custom data feeds. It handles inconsistencies and validates data integrity to ensure reliable market insights. For example, to configure a new data source, developers can implement the `DataSource` interface:
    `type DataSource interface {
        FetchData() (map[string]float64, error)
    }`
*   **zk-SNARK Position Automation:** Integrates zero-knowledge Succinct Non-Interactive ARguments of Knowledge (zk-SNARKs) to enable automated trading positions that are verifiable but conceal the underlying trading strategy. This prevents front-running and protects intellectual property. The framework provides utilities for generating and verifying proofs using a pre-defined circuit representing the trading logic.
*   **Serverless Execution Environment:** Supports deployment and execution within decentralized compute platforms like Functionland and secure enclaves, eliminating the need for centralized servers and enhancing scalability.
*   **Modular Architecture:** Designed with a modular architecture that allows developers to easily integrate custom trading strategies, data sources, and execution environments. New strategies can be added by implementing the `TradingStrategy` interface:
    `type TradingStrategy interface {
        Evaluate(data map[string]float64) (string, error) //Returns Buy, Sell, Hold
    }`
*   **On-Chain Order Management:** Provides tools for creating, managing, and executing orders directly on-chain through smart contract interactions. Handles gas optimization and transaction management for seamless order execution. Uses ABI encoding and signing for secure on-chain interactions.
*   **Risk Management Module:** Incorporates a risk management module that allows users to define constraints and limits on their trading strategies, mitigating potential losses. Includes features like stop-loss orders, take-profit levels, and position sizing controls.
*   **Event-Driven Architecture:** The framework utilizes an event-driven architecture to respond to market changes and trigger trading actions in real-time. This allows for dynamic and adaptive trading strategies. It utilizes channels for inter-module communication and asynchronous event handling.

## Technology Stack

*   **Go:** The primary programming language for the core framework logic due to its performance, concurrency features, and ease of use.
*   **Ethereum/Polygon/Arbitrum (Selectable):** The target blockchains for on-chain data aggregation and order execution. The framework supports multiple chains through a configurable chain ID.
*   **Solidity:** Used for developing and deploying smart contracts for on-chain order management and execution.
*   **zk-SNARKs (Groth16):** Implemented using a Go library, providing zero-knowledge proof generation and verification for position automation.
*   **gRPC:** Used for inter-process communication between modules and external data sources.
*   **Redis:** Used as a caching layer to improve performance and reduce on-chain data access frequency.

## Installation

1.  **Install Go:** Ensure you have Go 1.18 or later installed. Download and install from [https://go.dev/dl/](https://go.dev/dl/).
2.  **Clone the Repository:**
    `git clone https://github.com/jjfhwang/QuantumForge.git`
    `cd QuantumForge`
3.  **Install Dependencies:**
    `go mod download`
    `go mod tidy`
4.  **Install SnarkJS and Circom:**
    These are required for zk-SNARK circuit compilation and proof generation. Follow the instructions on the Circom website to install nodejs and then run:
    `npm install -g circom circomlib snarkjs`
5.  **Compile the zk-SNARK Circuit (example):**
    `cd circuits/example_circuit`
    `circom circuit.circom --r1cs --wasm --sym`
    `snarkjs powersoftau new bn128 12 example_circuit_0000.ptau.log`
    `snarkjs powersoftau contribute example_circuit_0000.ptau.log example_circuit_0001.ptau.log --name="First contribution"`
    `snarkjs powersoftau prepare phase2 example_circuit_0001.ptau.log example_circuit_final.ptau`
    `snarkjs groth16 setup circuit.r1cs example_circuit_final.ptau circuit_key.zkey`

## Configuration

The framework relies on environment variables for configuration. Create a `.env` file in the root directory and set the following variables:

*   `ETHEREUM_RPC_URL`: The URL of your Ethereum/Polygon/Arbitrum RPC provider (e.g., Infura, Alchemy).
*   `PRIVATE_KEY`: Your private key for signing transactions. **Important: Never commit your private key to a public repository.**
*   `CHAIN_ID`: The ID of the blockchain you are targeting (e.g., 1 for Ethereum Mainnet, 137 for Polygon).
*   `REDIS_ADDRESS`: The address of your Redis server (e.g., `localhost:6379`).
*   `CONTRACT_ADDRESS`: The address of the deployed smart contract for order execution.
*   `ZKEY_PATH`: The path to the generated `circuit_key.zkey` file.
*   `WASM_PATH`: The path to the generated `circuit.wasm` file.

You can load these environment variables using a library like `godotenv`:

    `import "github.com/joho/godotenv"`

    `err := godotenv.Load(".env")`
    `if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }`

## Usage

The framework provides a comprehensive API for interacting with on-chain data, managing orders, and generating zk-SNARK proofs.

*   **Fetching On-Chain Data:** Use the data aggregation module to fetch price data from DEXs:

    `price, err := dataAggregator.GetPrice("ETH/USDT")`
    `if err != nil {
        log.Fatalf("Error fetching price: %v", err)
    }`
*   **Creating and Executing Orders:** Use the order management module to create and execute orders on-chain:

    `order := orderManagement.NewOrder("ETH/USDT", "BUY", 1.0, 3000.0)`
    `txHash, err := orderManagement.ExecuteOrder(order, privateKey)`
    `if err != nil {
        log.Fatalf("Error executing order: %v", err)`
    }`

*   **Generating zk-SNARK Proofs:** The zk-SNARK module allows you to generate proofs for your trading strategy:

    `proof, publicSignals, err := zkSNARK.GenerateProof(inputData, zkeyPath, wasmPath)`
    `if err != nil {
        log.Fatalf("Error generating proof: %v", err)`
    }`

Complete API documentation, including detailed examples and code snippets, is available in the `docs` directory.

## Contributing

We welcome contributions to QuantumForge! Please follow these guidelines:

*   Fork the repository and create a branch for your changes.
*   Write clear and concise commit messages.
*   Submit a pull request with a detailed description of your changes.
*   Ensure your code adheres to the Go coding standards.
*   Include unit tests for any new functionality.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/QuantumForge/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the following projects and communities for their contributions and support:

*   Ethereum Foundation
*   Circom and SnarkJS developers
*   The Go programming language community