# Cassler - SSL Validator Tool

Validate SSL around web :spider: 

## Running Tests

```bash
go test -v
```

## Installation

### Using Go tools

```bash
go get github.com/msfidelis/cassler
```

## Usage

```bash
cassler --url https://sha512.badssl.com

Resolving: sha512.badssl.com on port 443

Server Certificate:
Common Name: sha512.badssl.com
Signature Algorithm: SHA512-RSA
Created: 2020-03-23 00:00:00 +0000 UTC
Expires: 2022-04-01 12:00:00 +0000 UTC
Expiration time: 543 days

Server IP's:
* 104.154.89.105

Certificate Authority:
* DigiCert SHA2 Secure Server CA
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)