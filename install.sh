#!/bin/bash

# The version to download
VERSION="v0.7.1"

# Determine the machine architecture
ARCH=$(uname -m)
echo "Detected architecture: $ARCH"

# Define the URL for downloading the file based on architecture
if [ "$ARCH" = "arm64" ]; then
    URL="https://github.com/commondatageek/mark/releases/download/${VERSION}/mark-${VERSION}-darwin-arm64.tar.gz"
elif [ "$ARCH" = "x86_64" ]; then
    URL="https://github.com/commondatageek/mark/releases/download/${VERSION}/mark-${VERSION}-darwin-amd64.tar.gz"
else
    echo "Unsupported architecture."
    exit 1
fi

# Download the file
echo "Downloading $URL..."
curl -L $URL -o mark.tar.gz

# Extract the file and remove the archive
tar -xzf mark.tar.gz
rm mark.tar.gz

# Make the binary executable
chmod +x mark

# Move the binary to /usr/local/bin
echo "Moving mark to /usr/local/bin. This may require your password."
sudo mv mark /usr/local/bin/

# macOS Gatekeeper handling (requires user to perform manually)
echo "To authorize the binary to run, you may need to:"
echo "  1. try to run 'mark' from the command line"
echo "  2. go to System Preferences -> Privacy & Security, and allow 'mark' to run"
