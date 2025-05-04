# Home

![GitHub Release](https://img.shields.io/github/v/release/avl-systems/atools?color=%237c3aed)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/avl-systems/atools/verify-installers.yml?label=installer%20tests)

**atools** is a collection of handy command-line interface(CLI) tools built using Go and [cobra-cli](github.com/spf13/cobra-cli).


## Quick Installation
### Linux
```
curl -fsSL https://atools.arnevanlooveren.be/installers/linux.sh | sh
```
### MacOS
A dedicated MacOS installer is not available yet.

Please use the [manual installation](#manual-installation) method or [build from source](#source-build).

### Windows
A dedicated Windows installer is not available yet.

Please use the [manual installation](#manual-installation) method or [build from source](#source-build).

## Manual Installation
### Linux
1. Download the latest release from the [releases page](https://github.com/avl-systems/atools/releases).
2. Unzip the downloaded file:
    ```
    tar -xzf atools_*_linux_*.tar.gz
    ```
3. Copy the `atools` binary to `/usr/bin`
    ```
    sudo cp atools /usr/bin
    ```
4. Make the binary executable:
    ```
    sudo chmod +x /usr/bin/atools
    ```
5. Verify the installation:
    ```
    atools version
    ```
(Or simply run `atools` to see the available commands.)

### MacOS
1. Download the latest release from the [releases page](https://github.com/avl-systems/atools/releases).
2. Unzip the downloaded file:
    ```
    tar -xzf atools_*_darwin_*.tar.gz
    ```
3. Copy the `atools` binary to `/usr/local/bin`
    ```
    sudo cp atools /usr/local/bin
    ```
4. Make the binary executable:
    ```
    sudo chmod +x /usr/local/bin/atools
    ```
5. Verify the installation:
    ```
    atools version
    ```
(Or simply run `atools` to see the available commands.)

### Windows
1. Download the latest release from the [releases page](https://github.com/avl-systems/atools/releases).
2. Unzip the downloaded file.
3. Copy the `atools.exe` binary to `C:\Program Files\atools` or any other directory in your PATH.
4. Open Command Prompt or Powershell and verify the installation with `atools version`.

## Source Build
1. Clone the repository:
    ```
    git clone https://github.com/avl-systems/atools.git
    ```
2. Navigate, build, and install:
    ```
    cd atools
    go mod tidy
    go build
    go install
    ```
3. Verify the installation:
    ```
    atools version
    ```
