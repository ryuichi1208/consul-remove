# Consul Node Manager

This Go application allows you to force a node to leave a Consul cluster and then prune the node from the cluster. It uses the HashiCorp Consul API to perform these operations.

## Features

- Force a node to leave the Consul cluster.
- Prune (deregister) a node from the Consul catalog.

## Prerequisites

- Go 1.16+
- Consul server

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/consul-node-manager.git
    cd consul-node-manager
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Build the application:

    ```sh
    go build -o consul-node-manager
    ```

2. Run the application with the required flags:

    ```sh
    ./consul-node-manager -node <node-name> -address <consul-address>
    ```

    Example:

    ```sh
    ./consul-node-manager -node node1 -address http://localhost:8500
    ```

    - `-node`: The name of the node to force leave and prune.
    - `-address`: The address of the Consul server (default: `http://localhost:8500`).

