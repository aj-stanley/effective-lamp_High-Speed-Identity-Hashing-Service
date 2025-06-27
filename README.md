# High-Speed Identity Hashing Service

A high-performance HTTP service for generating secure identity hashes using the BLAKE3 cryptographic hash function. This project combines the speed of Rust's BLAKE3 implementation with Go's excellent HTTP server capabilities through FFI (Foreign Function Interface).

## 🚀 Features

- **High Performance**: Leverages Rust's optimized BLAKE3 implementation for maximum speed
- **Secure Hashing**: Uses BLAKE3, a modern cryptographic hash function that's faster than SHA-3, SHA-2, and BLAKE2
- **HTTP API**: Simple REST API built with Go's standard library
- **Cross-Language Integration**: Demonstrates efficient Rust-Go interoperability via C FFI
- **Timestamp Support**: Includes timestamp-based hashing for time-sensitive identity generation
- **Memory Safe**: Proper memory management across language boundaries

## 🏗️ Architecture

The service consists of two main components:

### Rust Library (`rustlib/`)
- **Purpose**: High-performance cryptographic operations
- **Key Functions**:
  - `hash_identity()`: Generates BLAKE3 hash from email and timestamp
  - `free_str()`: Memory management for C-compatible strings
- **Dependencies**: BLAKE3 crate for cryptographic hashing

### Go HTTP Server (`gobin/`)
- **Purpose**: HTTP API and request handling
- **Features**:
  - REST endpoint at `/hash`
  - JSON request/response handling
  - Automatic timestamp generation if not provided
  - FFI integration with Rust library

## 📋 Prerequisites

- **Rust**: 1.70+ with Cargo
- **Go**: 1.19+ 
- **Make**: For build automation

## 🛠️ Installation & Build

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd effective-lamp_High-Speed-Identity-Hashing-Service
   ```

2. **Build the project**:
   ```bash
   make all
   ```
   This will:
   - Build the Rust library as a static library
   - Copy the library to the Go binary directory
   - Compile the Go HTTP server

3. **Run the service**:
   ```bash
   make run
   ```
   The server will start on port 8080.

## 🔧 Usage

### API Endpoint

**POST** `/hash`

**Request Body**:
```json
{
  "email": "user@example.com",
  "timestamp": 1640995200
}
```

**Response**:
```json
{
  "hash": "a1b2c3d4e5f6789..."
}
```

### Example Usage

```bash
# Hash with specific timestamp
curl -X POST http://localhost:8080/hash \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "timestamp": 1640995200}'

# Hash with current timestamp (timestamp field optional)
curl -X POST http://localhost:8080/hash \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com"}'
```

## 🧪 Hash Generation Process

1. **Input Preparation**: Combines email and timestamp as `"email:timestamp"`
2. **BLAKE3 Hashing**: Processes the input through BLAKE3 algorithm
3. **Hex Encoding**: Converts the hash to hexadecimal string format
4. **Response**: Returns the hash via JSON API

## 🏃‍♂️ Performance

- **BLAKE3 Algorithm**: Significantly faster than SHA-2 and SHA-3
- **Rust Implementation**: Leverages SIMD instructions and optimized assembly
- **Static Linking**: No runtime dependencies for the hash library
- **Minimal Overhead**: Efficient FFI with proper memory management

## 🗂️ Project Structure

```
.
├── Makefile              # Build automation
├── README.md            # This file
├── gobin/               # Go HTTP server
│   ├── main.go         # HTTP server implementation
│   ├── fastid          # Compiled binary (after build)
│   └── librustlib.a    # Rust static library (after build)
└── rustlib/            # Rust hashing library
    ├── Cargo.toml      # Rust dependencies
    ├── Cargo.lock      # Dependency lock file
    ├── src/
    │   └── lib.rs      # Core hashing implementation
    └── target/         # Rust build artifacts
```

## 🧹 Development

### Building Components Separately

```bash
# Build only Rust library
cd rustlib && cargo build --release

# Build only Go server (after Rust library is built)
cd gobin && go build -o fastid main.go
```

### Cleaning Build Artifacts

```bash
make clean
```

## 🔒 Security Considerations

- **BLAKE3**: Cryptographically secure hash function
- **Input Validation**: Server validates JSON input format
- **Memory Safety**: Proper cleanup of C strings to prevent leaks
- **No Secrets**: Service doesn't store or log sensitive data

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

[Add your license information here]

## 🆘 Troubleshooting

### Common Issues

1. **Build Failures**: Ensure Rust and Go are properly installed
2. **Port Conflicts**: Change port in `main.go` if 8080 is occupied
3. **Library Not Found**: Run `make clean && make all` to rebuild

### Getting Help

- Check build logs for specific error messages
- Ensure all prerequisites are installed
- Verify file permissions for build artifacts
