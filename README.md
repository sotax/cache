


          
# High Performance Local Memory Cache

## Overview
FastCache is a high-performance local memory cache library built with Golang, designed for efficient in-memory data storage and retrieval with support for automatic data expiration. The library is built on top of VictoriaMetrics' fastcache and has been verified in high-concurrency production environments.

## Features
- **High Performance**: Optimized for fast data access and storage operations
- **Automatic Expiration**: Built-in support for time-based data expiration
- **Asynchronous Time Refresh**: Uses non-blocking time refresh mechanism to minimize system call overhead
- **Customizable Cache Size**: Allows configuration of cache memory allocation
- **Thread Safety**: Designed to be safe for concurrent use
- **Production-Proven**: Verified in high-concurrency production environments

## Technical Implementation
- Leverages VictoriaMetrics/fastcache for efficient memory management
- Implements asynchronous time refresh using vDSO technology to avoid unnecessary system calls
- Stores data with expiration metadata in a compact format
- Uses sync.Once for thread-safe initialization

## Installation
```bash
go get github.com/sotax/cache
```

## Basic Usage
```go
import "github.com/sotax/cache"

// Create and initialize a cache instance
c := cache.Cache{
    Expire: 3600000, // Data expiration time in milliseconds (1 hour)
    Size: 10 * 1024 * 1024, // Cache size in bytes (10MB)
}
c.Init()

// Set a key-value pair
c.Set("user:1001", []byte(`{"name":"John","age":30}`))

// Get a value by key
value := c.Get("user:1001")
if value != nil {
    // Use the retrieved value
}

// Clear all data in the cache
c.Clear()
```

## Performance Optimization
- Asynchronous time refresh reduces system call overhead
- Optimized data storage format minimizes memory usage
- Efficient expiration check mechanism

## Dependencies
- [VictoriaMetrics/fastcache](https://github.com/VictoriaMetrics/fastcache) - High-performance in-memory cache for Go

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements
- Built on top of VictoriaMetrics' fastcache library
- Verified in high-concurrency production environments

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
        
