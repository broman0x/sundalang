$ErrorActionPreference = "Stop"

Write-Host "==================================" -ForegroundColor Cyan
Write-Host " SundaLang Uninstaller - Windows " -ForegroundColor Cyan
Write-Host "==================================" -ForegroundColor Cyan
Write-Host ""

$InstallDir = Join-Path $env:USERPROFILE ".sundalang"

if (-not (Test-Path $InstallDir)) {
    Write-Host "[ERROR] SundaLang teu kapanggih. Meureun geus diuninstall." -ForegroundColor Red
    exit 0
}

Write-Host "[INFO] Kapanggih instalasi di: $InstallDir" -ForegroundColor Yellow
Write-Host ""

$Confirm = Read-Host "Rek diuninstall SundaLang? (y/n)"
if ($Confirm -ne "y" -and $Confirm -ne "Y") {
    Write-Host "Uninstall dibatalkeun." -ForegroundColor Yellow
    exit 0
}

Write-Host "[INFO] Ngahapus file..." -ForegroundColor Yellow
try {
    Remove-Item -Path $InstallDir -Recurse -Force
    Write-Host "[OK] File geus dihapus" -ForegroundColor Green
} catch {
    Write-Host "[ERROR] Gagal ngahapus file: $_" -ForegroundColor Red
    exit 1
}

Write-Host "[INFO] Ngahapus tina PATH..." -ForegroundColor Yellow
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
$BinPath = Join-Path $InstallDir "bin"

if ($UserPath -like "*$BinPath*") {
    $NewPath = ($UserPath -split ';' | Where-Object { $_ -ne $BinPath }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Host "[OK] PATH geus diperbarui" -ForegroundColor Green
} else {
    Write-Host "[OK] PATH teu ngandung SundaLang" -ForegroundColor Green
}

Write-Host ""
Write-Host "[SUCCESS] SundaLang geus berhasil diuninstall!" -ForegroundColor Green
Write-Host "Hatur nuhun geus make SundaLang!" -ForegroundColor Cyan
Write-Host ""
