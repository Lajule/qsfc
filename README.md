# qsfc

Query Salesforce contacts from command line.

## Installation

```sh
go get github.com/Lajule/qsfc
```

> This binary uses [go-force][1] library.

## Usage

Type the following command `qsfc -h` to display this help message:

```
qsfc [flags] configfile
  -n	Counts the number of Salesforce contacts
  -w	Displays results with header
```

## Configuration

The configuration file looks like this:

```json
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

## Output

This program outputs following fields for all contacts:

* Id
* Name
* Email
* MobilePhone
* Title
* Department

> Fields are separated by `\t`.

## Example

```sh
qsfc config.json | cut -d$'\t' -f3 | xargs -I email mail -s "Happy new year" email <<EOF
> Dear customer,
> We wish you a happy new year.
>
> Best regards,
> The DEV team
> EOF
```

[1]: https://github.com/nimajalali/go-force "go-force"
