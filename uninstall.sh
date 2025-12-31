#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}==================================${NC}"
echo -e "${CYAN} SundaLang Uninstaller - Unix/Mac${NC}"
echo -e "${CYAN}==================================${NC}"
echo ""

INSTALL_DIR="$HOME/.sundalang"

if [ ! -d "$INSTALL_DIR" ]; then
    echo -e "${RED}[ERROR] SundaLang teu kapanggih. Meureun geus diuninstall.${NC}"
    exit 0
fi

echo -e "${YELLOW}[INFO] Kapanggih instalasi di: $INSTALL_DIR${NC}"
echo ""

read -p "Rek diuninstall SundaLang? (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}Uninstall dibatalkeun.${NC}"
    exit 0
fi

echo -e "${YELLOW}[INFO] Ngahapus file...${NC}"
rm -rf "$INSTALL_DIR"
echo -e "${GREEN}[OK] File geus dihapus${NC}"

echo -e "${YELLOW}[INFO] Ngahapus tina PATH...${NC}"

for RC_FILE in "$HOME/.bashrc" "$HOME/.zshrc" "$HOME/.profile"; do
    if [ -f "$RC_FILE" ]; then
        if grep -q ".sundalang/bin" "$RC_FILE" 2>/dev/null; then
            grep -v ".sundalang/bin" "$RC_FILE" > "${RC_FILE}.tmp"
            mv "${RC_FILE}.tmp" "$RC_FILE"
            echo -e "${GREEN}[OK] Dihapus tina $RC_FILE${NC}"
        fi
    fi
done

echo ""
echo -e "${GREEN}[SUCCESS] SundaLang geus berhasil diuninstall!${NC}"
echo -e "${CYAN}Hatur nuhun geus make SundaLang!${NC}"
echo ""
echo -e "${YELLOW}[INFO] Catetan: Tutup tur buka deui terminal pikeun nerapkeun parobahan${NC}"
