# cache-redis-config
======================

## Description
------------

cache-redis-config is a lightweight, high-performance Redis configuration management tool designed to simplify the process of managing Redis configurations in a multi-node Redis cluster. It provides a simple and intuitive API for setting, getting, and deleting Redis configuration keys across multiple nodes.

## Features
------------

* **Multi-Node Support**: Manage Redis configurations across multiple nodes with ease
* **Configuration Management**: Easily set, get, and delete Redis configuration keys
* **High-Performance**: Optimized for high-performance and low-latency operations
* **Simple API**: Intuitive and easy-to-use API for managing Redis configurations
* **Scalability**: Designed to handle large-scale operations and high-traffic Redis clusters

## Technologies Used
-------------------

* **Redis**: As the underlying database for storing and retrieving configuration data
* **Go**: As the primary language for building the cache-redis-config tool
* **GoKit**: A Go framework for building web services and APIs

## Installation
------------

### Prerequisites

* Go (1.13 or later) installed on your system
* Redis (6.0 or later) installed on your system
* A Redis server up and running on your system

### Building and Running

1. Clone the repository using Git: `git clone https://github.com/your-username/cache-redis-config.git`
2. Navigate to the project directory: `cd cache-redis-config`
3. Run the following command to build the project: `go build .`
4. Run the following command to start the server: `./cache-redis-config`

### Configuration

* Set the `REDIS_HOST` and `REDIS_PORT` environment variables to the hostname and port of your Redis server
* Set the `REDIS_DB` environment variable to the database number of your Redis server

### Example Use Cases
--------------------

* Set a Redis configuration key: `localhost:8080/config/set foo bar`
* Get a Redis configuration key: `localhost:8080/config/get foo`
* Delete a Redis configuration key: `localhost:8080/config/del foo`

### API Documentation
--------------------

For detailed API documentation, please refer to the [API Documentation page](API.md)

### Contributing
--------------

Contributions are welcome! Please fork the repository, make changes, and submit a pull request.

### License
----------

The cache-redis-config tool is released under the MIT License.