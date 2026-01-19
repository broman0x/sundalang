$ErrorActionPreference = "Stop"

function Get-TermWidth {
    try { return $Host.UI.RawUI.WindowSize.Width } catch { return 80 }
}

function Write-Center {
    param([string]$Text, [string]$Color = "White")
    $width = Get-TermWidth
    $padding = [Math]::Max(0, [Math]::Floor(($width - $Text.Length) / 2))
    Write-Host (" " * $padding + $Text) -ForegroundColor $Color
}

function Write-Box {
    param([string[]]$Lines, [string]$BorderColor = "Cyan", [string]$TextColor = "White")
    $width = Get-TermWidth
    $maxLen = ($Lines | Measure-Object -Maximum -Property Length).Maximum
    $boxWidth = [Math]::Min($maxLen + 6, $width - 4)
    $innerWidth = $boxWidth - 4
    $padding = [Math]::Max(0, [Math]::Floor(($width - $boxWidth) / 2))
    $pad = " " * $padding
    
    Write-Host ($pad + [char]0x256D + ([string][char]0x2500 * ($boxWidth - 2)) + [char]0x256E) -ForegroundColor $BorderColor
    foreach ($line in $Lines) {
        $textPad = [Math]::Max(0, [Math]::Floor(($innerWidth - $line.Length) / 2))
        $leftPad = " " * $textPad
        $rightPad = " " * ($innerWidth - $textPad - $line.Length)
        Write-Host ($pad + [char]0x2502 + " " + $leftPad + $line + $rightPad + " " + [char]0x2502) -ForegroundColor $BorderColor -NoNewline
        Write-Host ""
    }
    Write-Host ($pad + [char]0x2570 + ([string][char]0x2500 * ($boxWidth - 2)) + [char]0x256F) -ForegroundColor $BorderColor
}

function Write-Status {
    param([string]$Status, [string]$Text, [string]$Color)
    $width = Get-TermWidth
    $msg = "[$Status] $Text"
    $padding = [Math]::Max(0, [Math]::Floor(($width - $msg.Length) / 2))
    Write-Host (" " * $padding + $msg) -ForegroundColor $Color
}

function Write-Divider {
    $width = Get-TermWidth
    $divLen = [Math]::Min(50, $width - 10)
    $padding = [Math]::Max(0, [Math]::Floor(($width - $divLen) / 2))
    Write-Host (" " * $padding + ([string][char]0x2500 * $divLen)) -ForegroundColor DarkGray
}

[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

Clear-Host
Write-Host ""
Write-Box @(
    "",
    "   SUNDALANG   ",
    "",
    "Uninstaller for Windows",
    ""
) -BorderColor Red -TextColor White
Write-Host ""

$InstallDir = Join-Path $env:USERPROFILE ".sundalang"

if (-not (Test-Path $InstallDir)) {
    Write-Status "INFO" "SundaLang teu kapanggih. Meureun geus diuninstall." Yellow
    Write-Host ""
    exit 0
}

Write-Status "INFO" "Kapanggih instalasi di: $InstallDir" Yellow
Write-Host ""
Write-Divider
Write-Host ""

Write-Center "Rek diuninstall SundaLang? (y/n)" Cyan
$Confirm = Read-Host " "
if ($Confirm -ne "y" -and $Confirm -ne "Y") {
    Write-Host ""
    Write-Center "Uninstall dibatalkeun." Yellow
    Write-Host ""
    exit 0
}

Write-Host ""
Write-Divider
Write-Host ""
Write-Status "INFO" "Ngahapus file..." Yellow

try {
    Remove-Item -Path $InstallDir -Recurse -Force
    Write-Status "OK" "File geus dihapus" Green
} catch {
    Write-Host ""
    Write-Status "ERROR" "Gagal ngahapus file" Red
    exit 1
}

Write-Host ""
Write-Divider
Write-Host ""
Write-Status "INFO" "Ngahapus tina PATH..." Yellow

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
$BinPath = Join-Path $InstallDir "bin"

if ($UserPath -like "*$BinPath*") {
    $NewPath = ($UserPath -split ';' | Where-Object { $_ -ne $BinPath }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $NewPath, "User")
    Write-Status "OK" "PATH geus diperbarui" Green
} else {
    Write-Status "OK" "PATH teu ngandung SundaLang" Green
}

Write-Host ""
Write-Host ""
Write-Box @(
    "",
    "UNINSTALL SUKSES!",
    ""
) -BorderColor Green -TextColor Green
Write-Host ""
Write-Center "Hatur nuhun geus make SundaLang!" Cyan
Write-Host ""
