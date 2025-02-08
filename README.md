
# Probex - A Simple Probe Tool for TCP, HTTP, and gRPC 🧪

Probex is a simple command-line tool to check the availability of TCP, HTTP, and gRPC services. It allows you to probe specific ports or URLs to see if they are accessible within a given timeout. The tool supports both verbose and silent modes, allowing you to customize the output based on your needs.

## Features 🌟

- **TCP Probe**: Check if a TCP port is open.
- **HTTP Probe**: Check if an HTTP endpoint is reachable.
- **gRPC Probe**: Check if a gRPC server is reachable.
- **Silence Mode**: Suppress all output when the `-s` flag is provided.
- **Custom Timeout**: Set the timeout duration for each probe.

## Installation ⚙️

To install Probex, clone the repository and build the project:

```bash
git clone https://github.com/yourusername/probex.git
cd probex
go build -o probex .
```

## Usage 📚

### TCP Probe 🌐

Check if a TCP port is open:

```bash
./probex tcp <address> <port>
```

Example:

```bash
./probex tcp 127.0.0.1 8080
```

### HTTP Probe 🌍

Check if an HTTP endpoint is reachable:

```bash
./probex http <url>
```

Example:

```bash
./probex http http://example.com
```

Use the `-k` flag to skip TLS verification for HTTP probes:

```bash
./probex http https://example.com -k
```

### gRPC Probe 💬

Check if a gRPC server is reachable:

```bash
./probex grpc <address> <port>
```

Example:

```bash
./probex grpc 127.0.0.1 50051
```

### Silence Mode 🔕 

Suppress all output by using the `-s` flag:

```bash
./probex -s tcp 127.0.0.1 8080
```

### Timeout 🕒

Set a custom timeout with the `-t` flag:

```bash
./probex -t 5s tcp 127.0.0.1 8080
```

### Flags 🏷️

- `-s`: Silence mode (no output)
- `-t <timeout>`: Custom timeout (default: 3 seconds)
- `-k`: Skip TLS verification for HTTP probes

## Contributing 🤝

Feel free to open an issue or submit a pull request if you'd like to contribute! We'd love to have your help in improving Probex.

## License 📝

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
