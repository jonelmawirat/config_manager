# Config Manager in Go

This project implements the Singleton design pattern in Go, creating a configuration manager that ensures only one instance of the configuration object exists throughout the application's lifecycle. This pattern is particularly useful for managing shared resources, like configuration settings, that need to be consistently accessible from various parts of an application.

## Overview

The `configmanager` package provides a `Config` interface and a `config` struct that together implement the Singleton pattern. The `config` struct holds configuration data in a map, and the package provides a global `NewConfig` function to instantiate or retrieve the singleton instance.

## Features

- **Singleton Instance:** Ensures that only one instance of the `config` struct is created, providing a single, global point of access to the configuration data.
- **Thread-Safe Initialization:** Uses the `sync.Once` mechanism to ensure that the singleton instance is created only once, even in a concurrent environment.
- **Basic Configuration Management:** Allows setting and retrieving configuration values through `set` and `get` methods.

## Code Structure

- `configmanager.go`: Contains the `Config` interface, the `config` struct, and the `NewConfig` function to instantiate the singleton. It includes methods to set and retrieve configuration values.
- `test.go`: Provides tests for the `configmanager` package, ensuring that the singleton behavior is maintained and the configuration values are correctly set and retrieved.

## Usage

To use the configuration manager, you can call `NewConfig` to get the singleton instance and then use the `set` and `get` methods to manage configuration data.

```go
package main

import (
    "fmt"
    "config_manager"
)

func main() {
    // Retrieve or create the singleton instance
    conf := config_manager.NewConfig()

    // Set configuration values
    conf.set("Name", "John Doe")
    conf.set("Age", "31")

    // Retrieve configuration values
    name, _ := conf.get("Name")
    age, _ := conf.get("Age")

    fmt.Printf("Name: %s, Age: %s\n", name, age)
}
```

## Testing

To run the tests, use the following command:

```bash
go test
```

The tests in `test.go` verify the singleton behavior, ensuring that multiple calls to `NewConfig` return the same instance and that configuration values are shared across instances.

## Notes

- The `set` and `get` methods are not exported in the current implementation. If you intend to use them outside the `config_manager` package, you'll need to capitalize these methods to make them public (e.g., `Set` and `Get`).

## Conclusion

This project is a simple demonstration of the Singleton design pattern in Go, aimed at providing a consistent and accessible configuration management solution.
