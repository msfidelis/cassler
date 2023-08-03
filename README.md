![logo](.github/assets/logo.jpeg)

<p>
  <a href="README.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/fidelissauro" target="_blank">
    <img alt="Twitter: fidelissauro" src="https://img.shields.io/twitter/follow/fidelissauro.svg?style=social" />
  </a>
  <a href="/" target="_blank">
    <img alt="Build CI" src="https://github.com/msfidelis/cassler/workflows/cassler%20ci/badge.svg" />
  </a>  
  <a href="/" target="_blank">
    <img alt="Release" src="https://github.com/msfidelis/cassler/workflows/release%20packages/badge.svg" />
  </a>    
</p>

# Cassler - SSL Validator Tool

> Cassler is an CA's and SSL certificates analyzer. But if your read fast, it's sounds like "Cassia Eller"

Tooling to validate HTTPS Certificates and Connections Around Web :spider: 

## Running Tests

```bash
go test -v -race
```

## Running Linter

```bash
golint -set_exit_status ./...
```

## Installation

### Using Go tools

```bash
go get github.com/msfidelis/cassler
```

### On MacOSX amd64

```bash
wget https://github.com/msfidelis/cassler/releases/download/v0.0.12/cassler_0.0.12_darwin_amd64 -O /usr/local/bin/cassler

chmod +x /usr/local/bin/cassler
```

### On MacOSX arm64

```bash
wget https://github.com/msfidelis/cassler/releases/download/v0.0.12/cassler_0.0.12_darwin_arm64 -O /usr/local/bin/cassler

chmod +x /usr/local/bin/cassler
```

### On Linux x64

```bash
wget https://github.com/msfidelis/cassler/releases/download/v0.0.12/cassler_0.0.12_linux_amd64 -O /usr/local/bin/cassler

chmod +x /usr/local/bin/cassler
```

### Running on Docker

```bash
docker run -it fidelissauro/cassler:latest --url google.com
```

## Usage

```bash
cassler -h

  -dns string
    	DNS Server, default 8.8.8.8 (default "8.8.8.8")
  -lookup
    	Check reverse DNS Lookup for hosts IP's
  -mode check
    	Actions; Default: check; Available options `check` for check certificates, `tls` to test TLS connection, `scan` for complete checks on hosts (default "check")
  -port int
    	Server port, default: 443 (default 443)
  -url string
    	URL to validate SSL certificate,ex: https://google.com
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

### Check TLS Versions Enabled on Servers, with Reverse DNS Lookup 

```bash
cassler --url https://google.com --mode tls --lookup

Testing TLS Versions: google.com on port 443

DNS Lookup on: 8.8.8.8

TLS Versions Enabled on 2800:3f0:4001:824::200e:
- tls1.0: true
- tls1.1: true
- tls1.2: true
- tls1.3: true

TLS Versions Enabled on 142.250.219.174:
- tls1.0: true
- tls1.1: true
- tls1.2: true
- tls1.3: true


Starting reverse DNS Lookup on:  2800:3f0:4001:824::200e:

Starting reverse DNS Lookup on:  142.250.219.174:
142.250.219.174:  gru06s63-in-f14.1e100.net.
```

### Full Scan 

```bash
cassler --url https://tls-v1-2.badssl.com --port 1012 --mode scan

Checking Certificates: tls-v1-2.badssl.com on port 1012

DNS Lookup on: 8.8.8.8

Server Certificate:
Common Name: *.badssl.com
Issuer: CN=R3,O=Let's Encrypt,C=US
Subject: CN=*.badssl.com
Signature Algorithm: SHA256-RSA
Created: 2022-08-12 14:57:46 +0000 UTC
Expires: 2022-11-10 14:57:45 +0000 UTC
Expiration time: 83 days
Certificate Version: 3

DNS Names:
- *.badssl.com
- badssl.com

Issuing Certificate URL's:
- http://r3.i.lencr.org/

Server IP's:
* 104.154.89.105

Certificate Authority:

R3
Issuer: CN=ISRG Root X1,O=Internet Security Research Group,C=US
Subject: CN=R3,O=Let's Encrypt,C=US
Signature Algorithm: SHA256-RSA
Created: 2020-09-04 00:00:00 +0000 UTC
Expires: 2025-09-15 16:00:00 +0000 UTC
Expiration time: 1123 days
Certificate Version: 3


Issuing Certificate URL's:
- http://x1.i.lencr.org/


ISRG Root X1
Issuer: CN=DST Root CA X3,O=Digital Signature Trust Co.
Subject: CN=ISRG Root X1,O=Internet Security Research Group,C=US
Signature Algorithm: SHA256-RSA
Created: 2021-01-20 19:14:03 +0000 UTC
Expires: 2024-09-30 18:14:03 +0000 UTC
Expiration time: 773 days
Certificate Version: 3


Issuing Certificate URL's:
- http://apps.identrust.com/roots/dstrootcax3.p7c



Testing TLS Versions: tls-v1-2.badssl.com on port 1012

DNS Lookup on: 8.8.8.8

TLS Versions Enabled on 104.154.89.105:
- tls1.0: false
- tls1.1: false
- tls1.2: true
- tls1.3: false


Starting reverse DNS Lookup on:  104.154.89.105:
104.154.89.105:  105.89.154.104.bc.googleusercontent.com.
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
