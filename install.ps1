$ErrorActionPreference = "Stop"

Write-Host "==================================" -ForegroundColor Cyan
Write-Host "  SundaLang Installer - Windows  " -ForegroundColor Cyan
Write-Host "  Bahasa Pemrograman Sunda Pandeglang" -ForegroundColor Cyan
Write-Host "==================================" -ForegroundColor Cyan
Write-Host ""

$InstallDir = Join-Path $env:USERPROFILE ".sundalang\bin"
$BinaryPath = Join-Path $InstallDir "sundalang.exe"

if (-not (Test-Path $InstallDir)) {
    Write-Host "[INFO] Nyieun folder instalasi: $InstallDir" -ForegroundColor Yellow
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

Write-Host "[INFO] Nyari versi terbaru..." -ForegroundColor Yellow
try {
    $ReleaseInfo = Invoke-RestMethod -Uri "https://api.github.com/repos/broman0x/sundalang/releases/latest"
    $Version = $ReleaseInfo.tag_name
    
    $Asset = $ReleaseInfo.assets | Where-Object { $_.name -eq "sundalang.exe" }
    
    if (-not $Asset) {
        Write-Host "[ERROR] Gagal manggihan binary Windows" -ForegroundColor Red
        exit 1
    }
    
    $DownloadUrl = $Asset.browser_download_url
    Write-Host "[OK] Kapanggih: SundaLang $Version" -ForegroundColor Green
    
} catch {
    Write-Host "[ERROR] Gagal nyambung ka GitHub. Pastikeun koneksi internet aktif." -ForegroundColor Red
    Write-Host "Error: $_" -ForegroundColor Red
    exit 1
}

Write-Host "[INFO] Ngeundeur SundaLang $Version..." -ForegroundColor Yellow
try {
    $TempFile = Join-Path $env:TEMP "sundalang.exe"
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $TempFile
    
    Move-Item -Path $TempFile -Destination $BinaryPath -Force
    Write-Host "[OK] Hasil ngundeur" -ForegroundColor Green
    
} catch {
    Write-Host "[ERROR] Gagal ngundeur binary" -ForegroundColor Red
    Write-Host "Error: $_" -ForegroundColor Red
    exit 1
}

Write-Host "[INFO] Ngatur PATH..." -ForegroundColor Yellow
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($UserPath -notlike "*$InstallDir*") {
    $NewPath = "$UserPath;$InstallDir"
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Host "[OK] PATH geus diupdate" -ForegroundColor Green
    Write-Host ""
    Write-Host "[WARNING] PENTING: Tutup tur buka deui terminal pikeun nerapkeun parobahan PATH" -ForegroundColor Yellow
} else {
    Write-Host "[OK] PATH geus aya SundaLang" -ForegroundColor Green
}

Write-Host ""
Write-Host "[SUCCESS] Instalasi rengse!" -ForegroundColor Green
Write-Host ""
Write-Host "Kumaha carana make:" -ForegroundColor Cyan
Write-Host "  1. Tutup terminal ieu tur buka deui (pikeun PATH update)" -ForegroundColor White
Write-Host "  2. Jalankeun: sundalang --version" -ForegroundColor White
Write-Host "  3. Jalankeun file .sl: sundalang namafile.sl" -ForegroundColor White
Write-Host ""
Write-Host "Lokasi instalasi: $BinaryPath" -ForegroundColor Gray
Write-Host ""
Write-Host "Pikeun uninstall: irm https://raw.githubusercontent.com/broman0x/sundalang/main/uninstall.ps1 | iex" -ForegroundColor Gray
