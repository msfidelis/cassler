
![logo](.github/assets/logo.jpeg)

# Cassler - SSL Validator Tool

> If your read fast, it's sounds like "Cassia Eller"

Tooling to validate HTTPS Certificates and Connections Around Web :spider: 

## Running Tests

```bash
go test -v -race
```

## Installation

### Using Go tools

```bash
go get github.com/msfidelis/cassler
```

### On MacOSX

```bash
wget https://github.com/msfidelis/cassler/releases/download/v0.0.7/cassler_0.0.7_darwin_amd64 -O /usr/local/bin/cassler

chmod +x /usr/local/bin/cassler
```

### On Linux x64

```bash
wget https://github.com/msfidelis/cassler/releases/download/v0.0.7/cassler_0.0.7_linux_amd64 -O /usr/local/bin/cassler

chmod +x /usr/local/bin/cassler
```

## Usage

```bash
cassler -h
```

### Check Certificates

```bash
cassler --url google.com.br

Checking Certificates: google.com.br on port 443

Server Certificate:
Common Name: *.google.com.br
Issuer: CN=GTS CA 1O1,O=Google Trust Services,C=US
Subject: CN=*.google.com.br,O=Google LLC,L=Mountain View,ST=California,C=US
Signature Algorithm: SHA256-RSA
Created: 2020-09-22 15:29:04 +0000 UTC
Expires: 2020-12-15 15:29:04 +0000 UTC
Expiration time: 64 days
Certificate Version: 3

DNS Names:
- *.google.com.br
- google.com.br

Issuing Certificate URL's:
- http://pki.goog/gsr2/GTS1O1.crt

Server IP's: 
* 2800:3f0:4001:81b::2003
* 172.217.173.99

Certificate Authority:

GTS CA 1O1
Issuer: CN=GlobalSign,OU=GlobalSign Root CA - R2,O=GlobalSign
Subject: CN=GTS CA 1O1,O=Google Trust Services,C=US
Signature Algorithm: SHA256-RSA
Created: 2017-06-15 00:00:42 +0000 UTC
Expires: 2021-12-15 00:00:42 +0000 UTC
Expiration time: 429 days
Certificate Version: 3
```

### Check TLS Versions Enabled on Servers

```bash
cassler --url https://google.com --mode tls

Testing TLS Versions: google.com on port 443

TLS Versions Enabled on 2800:3f0:4001:813::200e:
- tls1.0: true
- tls1.1: true
- tls1.2: true
- tls1.3: true

TLS Versions Enabled on 172.217.162.142:
- tls1.0: true
- tls1.1: true
- tls1.2: true
- tls1.3: true
```

### Full Scan 

```bash
cassler --url https://tls-v1-2.badssl.com --port 1012 --mode scan

Checking Certificates: tls-v1-2.badssl.com on port 1012 

Server Certificate: 
Common Name: *.badssl.com
Issuer: CN=DigiCert SHA2 Secure Server CA,O=DigiCert Inc,C=US
Subject: CN=*.badssl.com,O=Lucas Garron Torres,L=Walnut Creek,ST=California,C=US
Signature Algorithm: SHA256-RSA
Created: 2020-03-23 00:00:00 +0000 UTC
Expires: 2022-05-17 12:00:00 +0000 UTC
Expiration time: 582 days
Certificate Version: 3

DNS Names: 
- *.badssl.com
- badssl.com

Issuing Certificate URL's: 
- http://cacerts.digicert.com/DigiCertSHA2SecureServerCA.crt

Server IP's: 
* 104.154.89.105 

Certificate Authority: 

DigiCert SHA2 Secure Server CA
Issuer: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
Subject: CN=DigiCert SHA2 Secure Server CA,O=DigiCert Inc,C=US
Signature Algorithm: SHA256-RSA
Created: 2013-03-08 12:00:00 +0000 UTC
Expires: 2023-03-08 12:00:00 +0000 UTC
Expiration time: 877 days
Certificate Version: 3



Testing TLS Versions: tls-v1-2.badssl.com on port 1012 

TLS Versions Enabled on 104.154.89.105: 
- tls1.0: false 
- tls1.1: false 
- tls1.2: true 
- tls1.3: false
```

### Specify a DNS Server

```bash
cassler --url raj.ninja --mode scan --dns 1.1.1.1
Checking Certificates: raj.ninja on port 443


DNS Lookup on: 1.1.1.1

Server Certificate:
Common Name: raj.ninja
Issuer: CN=Let's Encrypt Authority X3,O=Let's Encrypt,C=US
Subject: CN=raj.ninja
Signature Algorithm: SHA256-RSA
Created: 2020-11-26 20:46:27 +0000 UTC
Expires: 2021-02-24 20:46:27 +0000 UTC
Expiration time: 78 days
Certificate Version: 3

DNS Names:
- raj.ninja

Issuing Certificate URL's:
- http://cert.int-x3.letsencrypt.org/

Server IP's:
* 185.199.110.153
* 185.199.111.153
* 185.199.109.153
* 185.199.108.153
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
