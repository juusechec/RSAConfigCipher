package main

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

func main() {

	secretMessage := []byte("Ejemplo de mensaje para ser encriptado!")
	label := []byte("") // It can be empty https://golang.org/pkg/crypto/rsa/#EncryptOAEP

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, rsaPublicKey, secretMessage, label)
	if err != nil {
		log.Fatal("Error from encryption: %s", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// desencriptar
	ciphertext, _ = hex.DecodeString("60dcd0742e8f637bd76dc17763698852bd30c92b183393278a6d7c7f5fc6b1b7a958110fb3ccb851273dd2d0b292653be95a026ea7e48d809aa16486383c80101ccd352e4521ea8b19284e121ea9ffa2ae45b1497356a67d1e3ac74759fb352ee862c41df8cce51c36ab5f78d6fab32e13831f279dff740450c2787560328d244c36adba467d26a1efbffb0c473a928e2e1e516f8f8126d26d0138a2f9b805e906370ef3f5b333b4a1fa2b054915d19a33d726a2e580647e0b0b9677118ed34892e6b7b9a02d334ba0c6c662bb9940ac640ccdcb6d558cbbebe11339aabfe2312c222aab01196cef16aad229b1fff87c11d24a735da499328b00726f77cbcc8b")
	label = []byte("") // It can be empty https://golang.org/pkg/crypto/rsa/#EncryptOAEP

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng2 := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng2, rsaPrivateKey, ciphertext, label)
	if err != nil {
		log.Fatal("Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))
}
