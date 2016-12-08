# qsfc

Query Salesforce contacts from command line.

# Installation

```
go get github.com/Lajule/qsfc
```

# Build

```
go build
```

# Usage

Type the following command `qsfc -h` to display this help message:

```
qsfc [flags] file
  -n	Counts the number of Salesforce contacts
```

# Example

```
qsfc config.json | cut -d$'\t' -f3 | xargs -I email mail -s "Happy new year" <<EOF
> Dear customer,
> We wish you a happy new year.
>
> Best regards,
> The DEV team
> EOF
```
