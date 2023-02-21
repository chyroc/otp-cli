# otp-cli

Generate OTP Code Tool.

## Install

### By Go Get

```shell
go install github.com/chyroc/otp-cli@latest
```

### By Brew

```shell
brew tap chyroc/tap
brew install chyroc/tap/otp-cli
```

### By Docker

```shell
docker pull ghcr.io/chyroc/otp-cli
```

## Usage

```text
NAME:
   otp-cli - generate otp client

USAGE:
   otp-cli [global options] command [command options] [arguments...]

COMMANDS:
   version    show otp-cli version
   set-scope  set scope secret
   del-scope  delete scope secret
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value       otp secret text [$OTP_SECRET]
   --secret-file value, -f value  otp secret file [$OTP_SECRET_FILE]
   --scope value                  otp scope [$OTP_SCOPE]
   --copy, -c                     copy to clipboard (default: false) [$OTP_COPY]
   --quiet, -q                    not output to console (default: false) [$OTP_QUIET]
   --help, -h                     show help
```

- ***generate from secret text string***

```shell
otp-cli -s '<secret>'
```

- ***generate from secret file***

```shell
otp-cli -f '<secret file>'
```

- ***generate from scope***

```shell
# first: config scope
otp-cli set-scope --name <scope> --secret <secret>

# second: generate
otp-cli --scope <scope>
```

- ***generate and copy to clipboard***

```shell
otp-cli -s '<secret>' -c
```

- ***generate, copy to clipboard and not output to console***

```shell
otp-cli -s '<secret>' -c -q
```