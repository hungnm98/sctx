#!/usr/bin/env bash
set -e

VERSION=${VERSION:-"0.3.0"}
REPO="hungnm98/sctx"
INSTALL_DIR="$HOME/.sctx/bin"

echo "🚀 Installing sctx v$VERSION..."

# Detect OS
OS=$(uname -s)
case "$OS" in
  Linux*)   PLATFORM="linux" ;;
  Darwin*)  PLATFORM="darwin" ;;
  *) echo "❌ Unsupported OS: $OS"; exit 1 ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)   ARCH="amd64" ;;
  arm64)    ARCH="arm64" ;;
  aarch64)  ARCH="arm64" ;;
  *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Download URL
FILE="sctx-${VERSION}-${PLATFORM}-${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${FILE}"

echo "📦 Downloading $URL ..."
curl -fsSL -o /tmp/${FILE} ${URL}

echo "📂 Extracting to $INSTALL_DIR ..."
mkdir -p ${INSTALL_DIR}
tar -xzf /tmp/${FILE} -C ${INSTALL_DIR}
rm -f /tmp/${FILE}

# Ensure executable name is just 'sctx'
if [ -f "${INSTALL_DIR}/sctx-${VERSION}-${PLATFORM}-${ARCH}" ]; then
  mv "${INSTALL_DIR}/sctx-${VERSION}-${PLATFORM}-${ARCH}" "${INSTALL_DIR}/sctx"
fi
chmod +x "${INSTALL_DIR}/sctx"

# Add PATH to shell rc file
SHELL_NAME=$(basename "$SHELL")
case "$SHELL_NAME" in
  zsh)  RC_FILE="$HOME/.zshrc" ;;
  bash) RC_FILE="$HOME/.bashrc" ;;
  *)    RC_FILE="$HOME/.profile" ;;
esac

if ! grep -q 'sctx/bin' "$RC_FILE"; then
  echo '# sctx binary' >> "$RC_FILE"
  echo 'export PATH="$HOME/.sctx/bin:$PATH"' >> "$RC_FILE"
  echo "✅ Added sctx to PATH in $RC_FILE"
fi

echo "🎉 Installation complete!"
echo "➡️  Restart your shell or run: source $RC_FILE"
echo "➡️  Then try: sctx -h"