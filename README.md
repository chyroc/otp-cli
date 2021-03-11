# otp-cli

otp-cli is a tool for generate otp code in terminal.

## Install

```shell
go get github.com/chyroc/otp-cli
```

check

```shell
otp-cli version # expect output: otp-cli version v0.2.0
```

## Usage

```text
NAME:
   otp-cli - generate otp client

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value       otp secret text
   --secret-file value, -f value  otp secret file
   --copy, -c                     copy to clipboard (default: false)
   --quiet, -q                    not output to console (default: false)
   --help, -h                     show help (default: false)
```

- generate from secret text string

```shell
otp-cli -s '<secret>'
```

- generate from secret file

```shell
otp-cli -f '<secret file>'
```

- generate and copy to clipboard

```shell
otp-cli -s '<secret>' -c
```

- generate ,copy to clipboard and not output to console

```shell
otp-cli -s '<secret>' -c -q
```