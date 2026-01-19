# 03_kontrol_alur.sl - Loop & Break

# 1. Pikeun (For Loop)
# Catetan: Tanda titik-koma (;) WAJIB dipake di jero header pikeun misahkeun bagean
cetakkeun("--- Pikeun (For) ---")
pikeun tanda i = 1; i <= 5; i = i + 1 {
    cetakkeun("Angka: " + i)
    lamun i == 3 {
        cetakkeun("(Eureun heula di 3)")
        eureun # Break
    }
}

# 2. Kedap (While Loop)
cetakkeun("\n--- Kedap (While) ---")
tanda j = 0
kedap j < 3 {
    cetakkeun("Looping... " + j)
    j = j + 1
}
cetakkeun("Beres!")
