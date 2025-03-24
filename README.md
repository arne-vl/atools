# `atools` - Arne's CLI Tools
## 📦 Installation
### 🐧 Linux
```sh
curl https://raw.githubusercontent.com/arne-vl/atools/refs/heads/main/linux_installer.sh | sh
```
### Manual build
1. Clone the repository:
```
git clone https://github.com/arne-vl/atools.git
```
2. Build & install the project:
```
cd atools
go mod tidy
go build
go install
```
3. Check command installation:
```
atools
```

## 🚀 Features

- `linecounter` - Count the number of lines for a specified file extension.
- `portcheck` - Check if a given port is occupied or free.
- `ipinfo` - Get your hostname, private ip and public ip.

(More features coming soon!)
