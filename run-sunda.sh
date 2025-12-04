#!/bin/sh

# Menjalankan compiler Go
# "$@" artinya meneruskan semua argumen ke program Go
# Jika kosong, masuk mode REPL.

go run cmd/sundalang.go "$@"