#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
GRAY='\033[0;90m'
WHITE='\033[1;37m'
NC='\033[0m'

get_term_width() {
    if command -v tput &> /dev/null; then
        tput cols 2>/dev/null || echo 80
    else
        echo 80
    fi
}

center_text() {
    local text="$1"
    local color="${2:-$NC}"
    local width=$(get_term_width)
    local text_len=${#text}
    local padding=$(( (width - text_len) / 2 ))
    [ $padding -lt 0 ] && padding=0
    printf "%${padding}s" ""
    echo -e "${color}${text}${NC}"
}

draw_box() {
    local width=$(get_term_width)
    local box_width=44
    local padding=$(( (width - box_width) / 2 ))
    [ $padding -lt 0 ] && padding=0
    local pad=$(printf "%${padding}s" "")
    
    echo -e "${RED}${pad}╭──────────────────────────────────────────╮${NC}"
    echo -e "${RED}${pad}│                                          │${NC}"
    echo -e "${RED}${pad}│${WHITE}            ⬢  SUNDALANG  ⬢              ${RED}│${NC}"
    echo -e "${RED}${pad}│                                          │${NC}"
    echo -e "${RED}${pad}│${NC}        Uninstaller for Unix/Mac         ${RED}│${NC}"
    echo -e "${RED}${pad}│                                          │${NC}"
    echo -e "${RED}${pad}╰──────────────────────────────────────────╯${NC}"
}

draw_success_box() {
    local width=$(get_term_width)
    local box_width=36
    local padding=$(( (width - box_width) / 2 ))
    [ $padding -lt 0 ] && padding=0
    local pad=$(printf "%${padding}s" "")
    
    echo -e "${GREEN}${pad}╭──────────────────────────────────╮${NC}"
    echo -e "${GREEN}${pad}│                                  │${NC}"
    echo -e "${GREEN}${pad}│${WHITE}      ✓ UNINSTALL SUKSES!        ${GREEN}│${NC}"
    echo -e "${GREEN}${pad}│                                  │${NC}"
    echo -e "${GREEN}${pad}╰──────────────────────────────────╯${NC}"
}

draw_divider() {
    local width=$(get_term_width)
    local div_len=50
    [ $div_len -gt $((width - 10)) ] && div_len=$((width - 10))
    local padding=$(( (width - div_len) / 2 ))
    [ $padding -lt 0 ] && padding=0
    printf "%${padding}s" ""
    printf "${GRAY}"
    for ((i=0; i<div_len; i++)); do printf "─"; done
    printf "${NC}\n"
}

status_msg() {
    local status="$1"
    local text="$2"
    local color="$3"
    local msg="[${status}] ${text}"
    center_text "$msg" "$color"
}

clear
echo ""
draw_box
echo ""

INSTALL_DIR="$HOME/.sundalang"

if [ ! -d "$INSTALL_DIR" ]; then
    status_msg "INFO" "SundaLang teu kapanggih. Meureun geus diuninstall." "$YELLOW"
    echo ""
    exit 0
fi

status_msg "INFO" "Kapanggih instalasi di: $INSTALL_DIR" "$YELLOW"
echo ""
draw_divider
echo ""

center_text "Rek diuninstall SundaLang? (y/n)" "$CYAN"
echo ""
read -p "                                        " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo ""
    center_text "Uninstall dibatalkeun." "$YELLOW"
    echo ""
    exit 0
fi

echo ""
draw_divider
echo ""
status_msg "INFO" "Ngahapus file..." "$YELLOW"

rm -rf "$INSTALL_DIR"
status_msg "OK" "File geus dihapus" "$GREEN"

echo ""
draw_divider
echo ""
status_msg "INFO" "Ngahapus tina PATH..." "$YELLOW"
echo ""

for RC_FILE in "$HOME/.bashrc" "$HOME/.zshrc" "$HOME/.profile"; do
    if [ -f "$RC_FILE" ]; then
        if grep -q ".sundalang/bin" "$RC_FILE" 2>/dev/null; then
            grep -v ".sundalang/bin" "$RC_FILE" > "${RC_FILE}.tmp"
            mv "${RC_FILE}.tmp" "$RC_FILE"
            status_msg "OK" "Dihapus tina $(basename $RC_FILE)" "$GREEN"
        fi
    fi
done

echo ""
echo ""
draw_success_box
echo ""
center_text "Hatur nuhun geus make SundaLang!" "$CYAN"
echo ""
center_text "Tutup tur buka deui terminal pikeun nerapkeun parobahan" "$GRAY"
echo ""
