# Commands

* [`atools version`](#version) - Print the current version of atools.
* [`atools completion`](#completion) - Generate the autocompletion script for the specified shell.
* [`atools linecounter`](#linecounter) - Count the number of lines in files with a specific extension.
* [`atools portcheck`](#portcheck) - Checks if a port is occupied or free on the local machine.
* [`atools ipinfo`](#ipinfo) - Display the IP info of the current machine.
* [`atools jsonfmt`](#jsonfmt) - Pretty print JSON.
* [`atools construct`](#construct) - Construct your atools blueprints.

## Version
Prints the current installed version of atools.

Usage:
```
atools version
```

Example Output:
```
atools version 0.1.0
```

## Completion
Generate the autocompletion script for the specified shell.

Usage:
```
atools completion <shell>
```
Supported Shells:

- `bash`
- `zsh`
- `fish`
- `powershell`

## Linecounter
Count the number of lines in files with a specific extension.

Usage:
```
atools linecounter <extension> [flags]
```

> The extension is required and can be provided without the dot (e.g., `txt` for `.txt` files).

Flags:

- `-h`, `--help`              Help for linecounter.
- `-d`, `--directory` string  The directory to scan. Default is the current directory.
- `-r`, `--recursive`         Enable recursive search in subdirectories.
- `-l`, `--list`              List all scanned files with their line count (high to low).
- `-s`, `--spread`            Show the amount of files scanned.

## Portcheck
Check if a port is occupied or free on the local machine.

Usage:
```
atools portcheck <port>
```

## Ipinfo
Display the IP info of the current machine.

Usage:
```
atools ipinfo
```

Example Output:
```
Hostname:   my-computer
Private IP: 192.168.1.100
Public IP:  203.0.113.5
```

## Jsonfmt
Pretty print JSON.

Usage:
```
atools jsonfmt <json-file> [flags]
```
Or use it in a pipe:
```
cat file.json | atools jsonfmt
```

Flags:

- `-h`, `--help`     Help for jsonfmt.
- `-c`, `--compact`  Compact JSON output.

## Construct
Construct your atools blueprints.

Usage:
```
atools construct <blueprint-name> [flags]
```
> See how to create a blueprint [here](./blueprints.md).
>
> To call your blueprints, you can also use the name of the file without the extension (e.g., `my-blueprint` for `my-blueprint.yaml`).

Flags:

- `-h`, `--help`    Help for construct.
- `-s`, `--silent`  Silent mode (no output).
