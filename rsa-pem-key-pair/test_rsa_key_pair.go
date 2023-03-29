package main

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Load the private key from the file
    privateKeyFile, err := ioutil.ReadFile("private_key.pem")
    if err != nil {
        fmt.Println("Error loading private key file:", err)
        os.Exit(1)
    }
    privateKeyBlock, _ := pem.Decode(privateKeyFile)
    privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
    if err != nil {
        fmt.Println("Error parsing private key:", err)
        os.Exit(1)
    }

    // Sign a message using the private key
    message := []byte("Hello, world!")
    hash := sha256.Sum256(message)
    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
    if err != nil {
        fmt.Println("Error signing message:", err)
        os.Exit(1)
    }

    // Load the public key from the file
    publicKeyFile, err := ioutil.ReadFile("public_key.pem")
    if err != nil {
        fmt.Println("Error loading public key file:", err)
        os.Exit(1)
    }
    publicKeyBlock, _ := pem.Decode(publicKeyFile)
    publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
    if err != nil {
        fmt.Println("Error parsing public key:", err)
        os.Exit(1)
    }

    // Verify the signature using the public key
    hash = sha256.Sum256(message)
    err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
    if err != nil {
        fmt.Println("Error verifying signature:", err)
        os.Exit(1)
    }

    fmt.Println("RSA key pair tested successfully!")
}

