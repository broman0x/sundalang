#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GRAY='\033[0;90m'
NC='\033[0m'

echo -e "${CYAN}==================================${NC}"
echo -e "${CYAN}  SundaLang Installer - Unix/Mac ${NC}"
echo -e "${CYAN}  Bahasa Pemrograman Sunda Pandeglang${NC}"
echo -e "${CYAN}==================================${NC}"
echo ""

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
    linux*)
        OS="linux"
        ;;
    darwin*)
        OS="darwin"
        ;;
    *)
        echo -e "${RED}[ERROR] OS teu disupport: $OS${NC}"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}[ERROR] Architecture teu disupport: $ARCH${NC}"
        exit 1
        ;;
esac

if [ "$OS" = "darwin" ]; then
    if [ "$ARCH" = "arm64" ]; then
        BINARY_NAME="sundalang-macos-arm64"
    else
        BINARY_NAME="sundalang-macos"
    fi
else
    BINARY_NAME="sundalang"
fi

echo -e "${YELLOW}[INFO] Platform: $OS-$ARCH${NC}"
echo ""

INSTALL_DIR="$HOME/.sundalang/bin"
BINARY_PATH="$INSTALL_DIR/sundalang"

if [ ! -d "$INSTALL_DIR" ]; then
    echo -e "${YELLOW}[INFO] Nyieun folder instalasi: $INSTALL_DIR${NC}"
    mkdir -p "$INSTALL_DIR"
fi

echo -e "${YELLOW}[INFO] Nyari versi terbaru...${NC}"
RELEASE_INFO=$(curl -fsSL https://api.github.com/repos/broman0x/sundalang/releases/latest)
VERSION=$(echo "$RELEASE_INFO" | grep '"tag_name"' | cut -d '"' -f 4)

if [ -z "$VERSION" ]; then
    echo -e "${RED}[ERROR] Gagal manggihan versi terbaru${NC}"
    exit 1
fi

echo -e "${GREEN}[OK] Kapanggih: SundaLang $VERSION${NC}"

DOWNLOAD_URL=$(echo "$RELEASE_INFO" | grep "browser_download_url.*$BINARY_NAME\"" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo -e "${RED}[ERROR] Gagal manggihan binary pikeun $OS-$ARCH${NC}"
    exit 1
fi

echo -e "${YELLOW}[INFO] Ngeundeur SundaLang $VERSION...${NC}"
TEMP_FILE="/tmp/sundalang-$$.tmp"

if ! curl -fsSL -o "$TEMP_FILE" "$DOWNLOAD_URL"; then
    echo -e "${RED}[ERROR] Gagal ngundeur binary${NC}"
    rm -f "$TEMP_FILE"
    exit 1
fi

mv "$TEMP_FILE" "$BINARY_PATH"
chmod +x "$BINARY_PATH"
echo -e "${GREEN}[OK] Hasil ngundeur${NC}"

echo -e "${YELLOW}[INFO] Ngatur PATH...${NC}"

SHELL_RC=""
if [ -n "$ZSH_VERSION" ] || [ -f "$HOME/.zshrc" ]; then
    SHELL_RC="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ] || [ -f "$HOME/.bashrc" ]; then
    SHELL_RC="$HOME/.bashrc"
else
    SHELL_RC="$HOME/.profile"
fi

if ! grep -q ".sundalang/bin" "$SHELL_RC" 2>/dev/null; then
    echo "" >> "$SHELL_RC"
    echo 'export PATH="$HOME/.sundalang/bin:$PATH"' >> "$SHELL_RC"
    echo -e "${GREEN}[OK] PATH geus diupdate dina $SHELL_RC${NC}"
    echo ""
    echo -e "${YELLOW}[WARNING] PENTING: Jalankeun command ieu pikeun nerapkeun parobahan:${NC}"
    echo -e "${CYAN}    source $SHELL_RC${NC}"
    echo -e "${YELLOW}    Atawa tutup tur buka deui terminal${NC}"
else
    echo -e "${GREEN}[OK] PATH geus aya SundaLang${NC}"
fi

echo ""
echo -e "${GREEN}[SUCCESS] Instalasi rengse!${NC}"
echo ""
echo -e "${CYAN}Kumaha carana make:${NC}"
echo -e "  1. Jalankeun: ${NC}source $SHELL_RC"
echo -e "  2. Test: ${NC}sundalang --version"
echo -e "  3. Jalankeun file .sl: ${NC}sundalang namafile.sl"
echo ""
echo -e "${GRAY}Lokasi instalasi: $BINARY_PATH${NC}"
echo ""
echo -e "${GRAY}Pikeun uninstall: curl -fsSL https://raw.githubusercontent.com/broman0x/sundalang/main/uninstall.sh | bash${NC}"
