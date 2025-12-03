@echo off

rem Cek apakah user memberikan nama file
if "%1"=="" (
    echo [ERROR] Masukkan nama file!
    echo Cara pakai: run-sunda examples\hello.sl
    exit /b
)

rem Menjalankan compiler Go langsung
go run cmd/sundalang.go %1
echo.
