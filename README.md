# qsfc

Query Salesforce contacts from command line.

# Installation

```
go get github.com/Lajule/qsfc
```

> This binary uses [go-force][1] library.

# Build

```
go build
```

# Usage

Type the following command `qsfc -h` to display this help message:

```
qsfc [flags] configfile
  -n	Counts the number of Salesforce contacts
```

# Configuration

The configuration file looks like this:

```
{
  "version": "YOUR-API-VERSION",
  "clientid": "YOUR-CLIENT-ID",
  "clientsecret": "YOUR-CLIENT-SECRET",
  "username": "YOUR-USERNAME",
  "password": "YOUR-PASSWORD",
  "securitytoken": "YOUR-SECURITY-TOKEN",
  "environment": "YOUR-ENVIRONMENT"
}
```

# Example

```
qsfc config.json | cut -d$'\t' -f3 | xargs -I email mail -s "Happy new year" email <<EOF
> Dear customer,
> We wish you a happy new year.
>
> Best regards,
> The DEV team
> EOF
```

[1]: https://github.com/nimajalali/go-force "go-force"
