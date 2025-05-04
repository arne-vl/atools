# `atools` - AVL CLI Tools
![GitHub Release](https://img.shields.io/github/v/release/avl-systems/atools?color=%237c3aed)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/avl-systems/atools/verify-installers.yml?label=installer%20tests)

## Building manually
1. Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
2. Clone the repository:
   ```bash
    git clone git@github.com:avl-systems/atools.git
    cd atools
   ```
3. Build the project:
    ```bash
    go mod tidy
    go build
    go install
    ```
