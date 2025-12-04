@echo off

rem Menjalankan compiler Go
rem %* artinya meneruskan semua argumen (nama file atau flag -v) ke program Go
rem Jika kosong, program Go akan otomatis masuk mode REPL.

go run cmd/sundalang.go %*