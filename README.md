<p align="center">
  <img src="https://sundalang.emandev.xyz/logo.png" alt="Logo SundaLang" width="180"/>
</p>

<h1 align="center">SundaLang</h1>
<p align="center">
  <strong>Bahasa Pemrograman Sunda Pandeglang</strong><br>
  <em>Sederhana, Modern</em>
</p>

<p align="center">
  <a href="https://sundalang.emandev.xyz">
    <img src="https://img.shields.io/badge/web-official-blue" alt="web" />
  </a>
  <a href="https://raw.githubusercontent.com/broman0x/sundalang/refs/heads/main/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT" />
  </a>
  <a href="https://github.com/broman0x/sundalang">
    <img src="https://img.shields.io/github/stars/broman0x/sundalang?style=social" alt="GitHub stars" />
  </a>
  <a href="https://github.com/broman0x/sundalang/releases">
    <img src="https://img.shields.io/github/v/release/broman0x/sundalang?label=release" alt="Latest Release" />
  </a>
</p>

<p align="center">
  <a href="#-fitur-utama">Fitur</a> •
  <a href="#-cara-install">Install</a> •
  <a href="#-kontribusi">Kontribusi</a>
</p>

> **"Diajar koding bari ngamumule bahasa indung."**

**SundaLang** adalah bahasa pemrograman esoteric (esolang) berbasis interpreter yang menggunakan kosa kata asli **Bahasa Sunda Pandeglang**, Banten. Dibuat untuk edukasi, melatih logika berpikir, sekaligus ngamumule budaya lokal lewat coding.

## ✨ Fitur Utama

- **Sederhana** → Sintaks mirip Python/Go, mudah dipelajari pemula  
- **Nyunda 100%** → Keyword pakai dialek Pandeglang (kedap, lamun, cetakkeun, dsb)  
- **Ringan & Cepat** → Dibangun dengan Go, tanpa dependency berat  
- **Lengkap** → Variabel, if-else, loop, fungsi, operasi matematika  
- **Self-Installer** → Binary dengan installer built-in
- **Open Source** → Bebas dikembangkan bareng-bareng

## 📚 Kamus Syntax
 ```
 tanda > var / let
 ```
 ```
 cetakkeun > print
 ```
 ```
 tanyakeun > input
 ```
 ```
 lamun > if
 ```
 ```
 lamunteu > else
 ```
 ```
 kedap > while
 ```
 ```
 fungsi > function
 ```
 ```
 balik > return
 ```
 ```
 bener > true
 ```
 ```   
 salah > false
 ```
 ```   
 % > modulo
 ```
 ```
 && > AND
 ```


## 📦 Cara Install

### ⚡ Quick Install (Recommended)

Download dan jalankan installer otomatis dengan one-liner:

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/broman0x/sundalang/main/install.ps1 | iex
```

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/broman0x/sundalang/main/install.sh | bash
```

Script akan otomatis:
- Download binary terbaru dari GitHub Releases
- Install ke `~/.sundalang/bin/`
- Tambahkan ke PATH

**Verifikasi instalasi:**
```bash
sundalang --version
```

**Uninstall:**

Windows:
```powershell
irm https://raw.githubusercontent.com/broman0x/sundalang/main/uninstall.ps1 | iex
```

Linux/macOS:
```bash
curl -fsSL https://raw.githubusercontent.com/broman0x/sundalang/main/uninstall.sh | bash
```

---

### 🎯 Manual Install via Binary

Download binary dari [GitHub Releases](https://github.com/broman0x/sundalang/releases) sesuai platform:

**Windows:** `sundalang.exe`  
**Linux:** `sundalang`  
**macOS Intel:** `sundalang-macos`  
**macOS ARM:** `sundalang-macos-arm64`

Jalankan installer built-in:

```bash
./sundalang install
```

Binary akan otomatis:
- Copy dirinya ke `~/.sundalang/bin/`
- Update PATH (Windows otomatis, Linux/macOS perlu manual)

**Linux/macOS - Tambahkan ke PATH:**
```bash
echo 'export PATH="$HOME/.sundalang/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

**Uninstall:**
```bash
sundalang uninstall
```

---

### 🔧 Interactive Launcher

Jalankan binary tanpa argument untuk membuka **interactive menu**:

```bash
sundalang
```

Menu menyediakan:
- Install SundaLang ke sistem
- Uninstall SundaLang
- Jalankan file .sl
- Buka REPL mode
- Keluar

---

## 🚀 Cara Pakai

**Jalankan file .sl:**
```bash
sundalang hello.sl
```

**REPL Mode (Interactive):**
```bash
sundalang
# Pilih [4] Buka REPL mode
```

**Lihat versi:**
```bash
sundalang --version
```

**Help:**
```bash
sundalang --help
```

---

## � Development

Untuk development atau build dari source:

**Prasyarat:** Go 1.22+

```bash
git clone https://github.com/broman0x/sundalang.git
cd sundalang
go build -o sundalang ./cmd/...
```

**Jalankan:**
```bash
./sundalang examples/hello.sl
```

---

## 🤝 Kontribusi

Kami sangat senang kalau baraya rek nyumbangkeun!

1. Fork repo ini
2. Buat branch baru (`git checkout -b fitur-anyar`)
3. Commit (`git commit -m 'Nambihan fitur mantap'`)
4. Push & buat Pull Request

Ayo bareng-bareng ngajadikeun SundaLang leuwih gagah!

---

## 📄 Lisensi

Dirilis di bawah MIT License. Lihat file [LICENSE](LICENSE) untuk detail.

---
