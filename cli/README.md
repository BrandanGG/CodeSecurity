# Tyr CLI

A CLI tool for securing your code through Static Analysis and Composition Analysis.

## Installation

```bash
go build -o tyr .
```

## Commands

### `tyr scan`

Scan code for vulnerabilities using Static Analysis and Composition Analysis.

**Flags:**
- `-p, --path string`: Path to scan (default: current directory)
- `-t, --types stringSlice`: Scan types to run (default: `["all"]`)

**Examples:**
```bash
tyr scan
tyr scan -p ./src -t sast
```

### `tyr version`

Print the version number.

**Examples:**
```bash
tyr version
tyr -v
```
