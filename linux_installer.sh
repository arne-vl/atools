#!/bin/sh

set -e # Exit on error

echo "Welcome to the atools linux installer!"
echo
LATEST_VERSION=$(curl -s https://api.github.com/repos/arne-vl/atools/tags | grep '"name":' | sed -E 's/.*"v?([^"]+)".*/\1/' | sort -Vr | head -n 1)

if command -v atools > /dev/null; then
        echo "atools already installed"
        echo "Checking version..."
        CURRENT_VERSION=$(atools version | sed -n 's/.*: \([0-9.]*\)/\1/p')

        if [ $CURRENT_VERSION != $LATEST_VERSION ]; then
                echo "Current version:" $CURRENT_VERSION
                echo "Updating to:" $LATEST_VERSION
                echo
        else
                echo "Most recent version is installed, aborting..."
                exit 0
        fi
fi

OS=$(uname | tr "[:upper:]" "[:lower:]")
echo "Found OS:" $OS
if [ "$OS" != "linux" ]; then
        echo "Operating system is not Linux. Aborting installer..."
        exit 1
fi

ARCH=$(uname -m | tr "[:upper:]" "[:lower:]")
echo "Found ARCH:" $ARCH

echo

echo "Finding latest matching binary..."

if [ "$ARCH" = "x86_64" ]; then
        BINARY_NAME="atools_${LATEST_VERSION}_linux_amd64.tar.gz"
elif [ "$ARCH" = "aarch64" ]; then
        BINARY_NAME="atools_${LATEST_VERSION}_linux_arm64.tar.gz"
elif [ "$ARCH" = "i386" ]; then
        BINARY_NAME="atools_${LATEST_VERSION}_linux_386.tar.gz"
else
        echo "Could not find binary for architecture '${ARCH}'. Aborting..."
        exit 1
fi
echo "Found matching binary:" $BINARY_NAME

echo

echo "Downloading binary..."
DOWNLOAD_URL="https://github.com/arne-vl/atools/releases/download/v${LATEST_VERSION}/${BINARY_NAME}"
FILE_PATH="/tmp/${BINARY_NAME}"
curl -Lso $FILE_PATH $DOWNLOAD_URL
echo "Finished downloading binary"

echo

echo "Unpacking..."
tar -xzf $FILE_PATH -C /tmp
echo "Installing..."
sudo mv /tmp/atools /bin/atools

echo

echo "Creating config directory..."
if [ ! -f ~/.config/atools/example.yml ]; then
    mkdir -p ~/.config/atools/blueprints
    cat <<EOF > ~/.config/atools/blueprints/example.yml
# Example config file for atools construct.
# This file is used to configure an atools blueprint for the construct command.
# You can find more information about the blueprint file format here:
# https://atools.arnevanlooveren.be/#docs

blueprint:
  directories:
    - example
  files:
    - path: example/example.txt
      content: |
        This is an example file.
        You can use this file to test atools.
EOF
fi

echo

echo "Install successful!"

echo

atools
