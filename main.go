package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	ecies "github.com/ecies/go/v2"
)

func encrypt(publicKey string, message string) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		fatal(err)
	}
	messageBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		messageNew, err := strconv.Unquote(`"` + message + `"`)
		if err == nil {
			message = messageNew
		}
		messageBytes = []byte(message)
	}

	pubKey, err := ecies.NewPublicKeyFromBytes(publicKeyBytes)
	if err != nil {
		fatal(err)
	}
	ciphertext, err := ecies.Encrypt(pubKey, messageBytes)
	if err != nil {
		fatal(err)
	}

	b64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println(b64)
}

func decrypt(privateKey string, message string) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		fatal(err)
	}
	messageBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fatal(err)
	}

	privkey := ecies.NewPrivateKeyFromBytes(privateKeyBytes)
	plaintext, err := ecies.Decrypt(privkey, messageBytes)
	if err != nil {
		fatal(err)
	}

	fmt.Println(string(plaintext))
}

func privateKey() {
	privateKey, err := ecies.GenerateKey()
	if err != nil {
		fatal(err)
	}

	privateKeyBytes := privateKey.Bytes()
	b64 := base64.StdEncoding.EncodeToString(privateKeyBytes)

	fmt.Println(b64)
}

func publicKey(privateKeyB64 string) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyB64)
	if err != nil {
		fatal(err)
	}
	privkey := ecies.NewPrivateKeyFromBytes(privateKeyBytes)

	publicKeyBytes := privkey.PublicKey.Bytes(false)
	b64 := base64.StdEncoding.EncodeToString(publicKeyBytes)

	fmt.Println(b64)
}

func usage() {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("    privatekey\n"))
	builder.WriteString(fmt.Sprintf("    publickey PRIVATE_KEY\n"))
	builder.WriteString(fmt.Sprintf("    decrypt PRIVATE_KEY MESSAGE\n"))
	builder.WriteString(fmt.Sprintf("    decrypt PRIVATE_KEY MESSAGE\n"))

	fmt.Printf("Command Line Options\nUsage: %s [\n%s]", os.Args[0], builder.String())
	os.Exit(0)
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	switch os.Args[1] {
	case "privatekey":
		if len(os.Args) != 2 {
			usage()
		}

		privateKey()
		break
	case "publickey":
		if len(os.Args) != 3 {
			usage()
		}

		publicKey(os.Args[2])
		break
	case "encrypt":
		if len(os.Args) != 4 {
			usage()
		}

		encrypt(os.Args[2], os.Args[3])
		break
	case "decrypt":
		if len(os.Args) != 4 {
			usage()
		}

		decrypt(os.Args[2], os.Args[3])
		break
	}
}
