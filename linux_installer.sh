#!/bin/sh

echo "Welcome to the atools linux installer!"
echo
OS=$(uname | tr "[:upper:]" "[:lower:]")
echo 'Found OS:' $OS

if [ "$OS" != "linux" ]; then
        echo "Operating system is not Linux. Aborting installer..."
        exit 1
fi

ARCH=$(uname -m | tr "[:upper:]" "[:lower:]")
echo "Found ARCH:" $ARCH
echo
echo "Finding latest matching binary..."
VERSION=$(curl -s https://api.github.com/repos/arne-vl/atools/tags | grep "name" | sed -E 's/.*v([^"]+)".*/\1/' | head -n 1)

if [ "$ARCH" = "x86_64" ]; then
        BINARY_NAME="atools_${VERSION}_linux_amd64.tar.gz"
elif [ "$ARCH" = "aarch64" ]; then
        BINARY_NAME="atools_${VERSION}_linux_arm64.tar.gz"
elif [ "$ARCH" = "i386" ]; then
        BINARY_NAME="atools_${VERSION}_linux_386.tar.gz"
else
        echo "Could not find binary for architecture '${ARCH}'. Aborting..."
        exit 1
fi
echo "Found matching binary:" $BINARY_NAME
echo
echo "Downloading binary..."
DOWNLOAD_URL="https://github.com/arne-vl/atools/releases/download/v${VERSION}/${BINARY_NAME}"
FILE_PATH="/tmp/${BINARY_NAME}"
curl -Lso $FILE_PATH $DOWNLOAD_URL
echo "Finished downloading binary"
echo
echo "Unpacking..."
tar -xzf $FILE_PATH -C /tmp
echo "Installing..."
sudo mv /tmp/atools /bin/atools
echo "Install successful!"
echo
atools
