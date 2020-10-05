# Cassler - SSL Validator Tool

Validate SSL around web :spider: 

## Running Tests

```bash
go test -v -race
```

## Installation

### Using Go tools

```bash
go get github.com/msfidelis/cassler
```

## Usage

### Check Certificates

```bash
cassler --url https://google.com

Checking Certificates: google.com on port 443

Server Certificate:
Common Name: *.google.com
Signature Algorithm: SHA256-RSA
Created: 2020-09-03 06:36:33 +0000 UTC
Expires: 2020-11-26 06:36:33 +0000 UTC
Expiration time: 52 days

Server IP's:
* 2800:3f0:4001:813::200e
* 172.217.172.142

Certificate Authority:
* GTS CA 1O1
```

### Check TLS Versions Enabled on Servers

```bash
cassler --url https://google.com --mode tls

Testing TLS Versions: google.com on port 443

TLS Versions Enabled on 2800:3f0:4001:813::200e:
* tls1.0: true
* tls1.1: true
* tls1.2: true
* tls1.3: true

TLS Versions Enabled on 172.217.162.142:
* tls1.0: true
* tls1.1: true
* tls1.2: true
* tls1.3: true
```

### Full Scan 

```bash
cassler --url https://tls-v1-2.badssl.com --port 1012 --mode scan

Checking Certificates: tls-v1-2.badssl.com on port 1012

Server Certificate:
Common Name: *.badssl.com
Signature Algorithm: SHA256-RSA
Created: 2020-03-23 00:00:00 +0000 UTC
Expires: 2022-05-17 12:00:00 +0000 UTC
Expiration time: 589 days

Server IP's:
* 104.154.89.105

Certificate Authority:
* DigiCert SHA2 Secure Server CA

Testing TLS Versions: tls-v1-2.badssl.com on port 1012

TLS Versions Enabled on 104.154.89.105:
* tls1.0: false
* tls1.1: false
* tls1.2: true
* tls1.3: false

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)