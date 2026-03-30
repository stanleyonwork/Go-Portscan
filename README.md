# Go-PortScan: High-Concurrency TCP Scanner

## Architectural Overview
A high-performance network reconnaissance tool engineered in Golang. Optimized for RAM efficiency and massive concurrent execution. The system utilizes a worker pool architecture to simultaneously scan 65,535 ports without inducing memory leaks or OS thread overhead.

## Technical Specifications
* **Optimal Concurrency:** Implementation of 1,000 goroutines as workers, managed via channel abstractions to minimize blocking operations.
* **I/O Management:** Utilizes `bufio` for dynamic target processing, seamlessly resolving IPv4, IPv6, and hostnames in real-time.
* **Data Integrity:** Asynchronous result aggregation deterministically ordered using the `sort` package for precise post-scan analysis.
* **Transport Layer Security:** TCP connection timeout regulation via `net.DialTimeout` to prevent file descriptor exhaustion on the host system.

## Installation & Execution

```bash
# Clone the repository
git clone https://github.com/stanleyonwork/Go-Portscan.git
cd go-portscan

# Compile binary with size optimization (strip debug symbols)
go build -ldflags="-s -w" -o scanner main.go

# Execute the binary
go run portscanner.go
```

## Security Protocol (Disclaimer)
This module is strictly developed for infrastructure auditing, vulnerability management, and data security research. Executing scans on unauthorized networks or infrastructure without explicit, written permission from the system owner is strictly prohibited. The developer assumes no liability for any misuse or damage caused by this utility.