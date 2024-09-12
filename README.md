# ecies-cli

Command line utility to encrypt / decrypt string messages with [ECIES](https://en.wikipedia.org/wiki/Integrated_Encryption_Scheme) (Ethereum-used cryptography scheme).

# Usage

```text
Usage: ecies-cli [command] [arg]


Command Line Options
Usage: ./ecies-cli [
    privatekey						creates a private key
    publickey PRIVATE_KEY			show public key
    decrypt PRIVATE_KEY MESSAGE		decrypts a message encrypted by public key
    decrypt PRIVATE_KEY MESSAGE		encrypts a message with public key

```

# Build

```text
go build
```