package cipher_value

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	rsaPrivateKeyPath = "keys/rsakey.pem"     // openssl genrsa -out rsakey.pem 2048
	rsaPublicKeyPath  = "keys/rsakey.pem.pub" // openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub
)

var (
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
)

func init() {
	fmt.Println("Reading and parse key files...")

	fmt.Println("Reading public key.")
	// Leer llave pública.
	pubPEMData, err := ioutil.ReadFile(rsaPublicKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	block, rest := pem.Decode([]byte(pubPEMData))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		fmt.Println("pub is of type RSA:", pub)
	case *dsa.PublicKey:
		fmt.Println("pub is of type DSA:", pub)
	case *ecdsa.PublicKey:
		fmt.Println("pub is of type ECDSA:", pub)
	default:
		panic("unknown type of public key")
	}

	fmt.Printf("Got a %T, with remaining data: %q\n", pub, rest)

	rsaPublicKey = pub.(*rsa.PublicKey)

	// Leer llave privada.
	fmt.Println("Reading private key.")

	privPEMData, err := ioutil.ReadFile(rsaPrivateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	block, rest = pem.Decode([]byte(privPEMData))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("failed to parse PKCS1 private key: " + err.Error())
	}

	switch priv := interface{}(priv).(type) {
	case *rsa.PrivateKey:
		fmt.Println("priv is of type RSA:", priv)
	case *dsa.PrivateKey:
		fmt.Println("priv is of type DSA:", priv)
	case *ecdsa.PrivateKey:
		fmt.Println("priv is of type ECDSA:", priv)
	default:
		panic("unknown type of private key")
	}

	fmt.Printf("Got a %T, with remaining data: %q\n", priv, rest)
	rsaPrivateKey = priv

	fmt.Println("Finish Reading and parse key files!!!")
}

func EncryptValue(secretMessage string) (string, error) {
	secretMessageByte := []byte(secretMessage)
	label := []byte("") // It can be empty https://golang.org/pkg/crypto/rsa/#EncryptOAEP

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, rsaPublicKey, secretMessageByte, label)
	if err != nil {
		log.Fatal("Error from encryption: %s", err)
		return "", err
	}
	ciphertextHex := hex.EncodeToString(ciphertext)
	return ciphertextHex, nil
}

func DecryptValue(ciphertextHex string) (string, error) {
	ciphertext, _ := hex.DecodeString(ciphertextHex)

	label := []byte("") // It can be empty https://golang.org/pkg/crypto/rsa/#EncryptOAEP

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, rsaPrivateKey, ciphertext, label)
	if err != nil {
		log.Fatal("Error from decryption: %s\n", err)
		return "", err
	}

	return string(plaintext), nil
}
