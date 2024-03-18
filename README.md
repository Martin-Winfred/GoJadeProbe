# GoJadeProbe

GoJadeProbe is a system probe written in Go, designed to collect and report various information about the host.

## Features

- Collects host information such as architecture, operating system, hostname, kernel version, version, platform, family, CPU load, memory usage, used memory, total memory, network name, received bytes, sent bytes, and local IP.
- Stores the collected data in a SQLite database.
- Provides an API to query and export the data.

## Installation

First, you need to install Go. Then, you can use the following command to get and install GoJadeProbe:

```bash
git clone https://github.com/Martin-Winfred/GoJadeProbe.git
```

## Usage

To run GoJadeProbe, you can use the following command:

```bash
go run main.go
```

Then, you can visit `http://localhost:8080` in your browser to view and query the data.

## Contributing

Contributions of any kind are welcome, including reporting issues, suggesting new features, improving the code, etc.

## License

GoJadeProbe is licensed under the *AGPL* License. See the `LICENSE` file for details.
