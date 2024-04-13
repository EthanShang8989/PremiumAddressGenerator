# Ethereum Address Generator

## Overview

This Ethereum Address Generator allows users to generate Ethereum addresses that match specific prefixes and suffixes. It utilizes multiple CPU cores to efficiently create addresses.

## Features

- **Customizable Address Generation**: Users can specify desired prefixes and suffixes for the Ethereum addresses.
- **Multi-core Support**: Leverages multiple CPU cores for enhanced performance.
- **Simple and Audit-friendly Code**: The code is designed to be simple and easy to audit, ensuring minimal risk.

## Requirements

- [Go](https://golang.org/) (Version 1.14 or later recommended)
- [Ethereum Go-Ethereum](https://github.com/ethereum/go-ethereum) library

## Installation

To set up the Ethereum Address Generator on your system, follow these steps:

1. Ensure that Go is installed on your system. You can download it from https://golang.org/dl/.

2. Install the Ethereum Go-Ethereum library:

   ```
   go get -u github.com/ethereum/go-ethereum
   ```

3. Clone this repository:

   ```
   git clone  https://github.com/EthanShang8989/PremiumAddressGenerator.git
   ```

4. Navigate to the cloned repository:

   ```
   cd path/to/repo
   ```

## Usage

To run the program, use the following command:

```
go run main.go -prefix=<desired_prefix> -suffix=<desired_suffix> -cores=<number_of_cores>
```

### Flags

- `-prefix`: Specify the prefix that the Ethereum address should start with.
- `-suffix`: Specify the suffix that the Ethereum address should end with.
- `-cores`: Define how many CPU cores to use for generating addresses.

Example:

```
go run main.go -prefix="1a2b3" -suffix="4c5d6" -cores=4 
```

This will start the address generator using 4 CPU cores, trying to generate addresses starting with "1a2b3" and ending with "4c5d6".

## Contributing

Contributions to the Ethereum Address Generator are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.