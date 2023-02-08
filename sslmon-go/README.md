<!-- @format -->

# SSL Monitor in Go 🔒💻

A simple Go program to monitor an SSL website and retrieve information about the SSL certificate and its chain.

## Features 🎉

- Monitors an SSL website and retrieves the status code of the response.
- Retrieves information about the SSL certificate, including:
  - Version
  - Cipher Suite
  - Subject
  - Issuer
  - Validity Period
  - Signature Algorithm
  - Public Key Algorithm
  - Converts the information to a JSON string for easy parsing.

## Usage 💻

1. Clone the repository to your local machine:

2. Navigate to the directory:

```
cd sslmon-go
```

3. Replace the URL in main.go with the desired SSL website to monitor.
4. Run the program:

```
go run main.go
```

## Output 📈

The program outputs the status code of the response and the information about the SSL certificate in JSON format:

```
{
  "Version": 771,
  "HandshakeComplete": true,
  "DidResume": false,
  "CipherSuite": 49199,
  "PeerCertificates": [
    {
      "Subject": "CN=www.example.com, O=Example Inc, L=San Francisco, ST=California, C=US",
      "Issuer": "CN=Example Root CA, O=Example Inc, L=San Francisco, ST=California, C=US",
      "NotBefore": "2021-01-01T00:00:00Z",
      "NotAfter": "2024-01-01T00:00:00Z",
      "SignatureAlgorithm": "sha256WithRSAEncryption",
      "PublicKeyAlgorithm": "rsa"
    }
  ]
}
```

## Note 📝

For testing purposes, the script skips SSL certificate verification by setting InsecureSkipVerify to true. In a production environment, it's recommended to set InsecureSkipVerify to false and use a valid SSL certificate.

## Contribution 🤝

Contributions are welcome! If you find a bug or have an idea for a feature, please open an issue or submit a pull request.

## License 📄

This project is licensed under the MIT License. See the LICENSE file for details.
