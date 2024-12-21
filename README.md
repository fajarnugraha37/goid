# GOID

goid is a Go library that provides implementations for ULID and various UUID versions, including UUID v1, v2, v3, v4, v5, v6, and v7. The library is designed to be modular, efficient, and thread-safe, using only the Go standard library.

## Features
- ULID: Universally Unique Lexicographically Sortable Identifier.
    consists of 26 characters, which includes:
    - A 48-bit timestamp (milliseconds since Unix epoch).
    - A 80-bit random component.
- UUID v1: Time-based UUID.
    Combines the current timestamp with the MAC address of the generating machine: 
    - 60 bits for timestamp
    - 48 bits for MAC address
    - 14 bits for a sequence number.
- UUID v2: DCE Security UUID. 
    Similar to v1 but includes a local domain identifier (e.g., POSIX UID/GID).
- UUID v3: Name-based UUID using MD5:
    - Generated from a namespace identifier and a name using the MD5 hashing algorithm.
    - Deterministic: the same input will always produce the same UUID.
- UUID v4: Random UUID.
    - Generated using random numbers.
    - Most commonly used due to its simplicity and randomness.
- UUID v5: Name-based UUID using SHA-1.
    - Similar to v3 but uses SHA-1 instead of MD5 for hashing.
    - Also deterministic.
- UUID v6: Time-based UUID (reordered).
    - A variant of v1 that reorders the timestamp to improve sorting.
- UUID v7: Time-based UUID (Unix epoch).
    - A new version that uses a Unix timestamp in milliseconds and random bits.
    - Designed for better performance and sorting.
  
## Installation

To install the goid library, use the following command:
```bash
go get github.com/fajarnugraha37/goid
```

## Usage

### Generating a ULID

### Generating a UUID v1

### Generating a UUID v2

### Generating a UUID v3

### Generating a UUID v4

### Generating a UUID v5

### Generating a UUID v6

### Generating a UUID v7

## Contributing

Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
