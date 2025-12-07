# 1. Definisi Fungsi
tanda tambah = fungsi(a, b) {
    balik a + b
}

# 2. Ngagunakeun Array 
cetakkeun("--- Tes Array ---")
tanda nilai = [80, 90, 100]
cetakkeun("Nilai kahiji:")
cetakkeun(nilai[0]) 
cetakkeun("Nilai katilu:")
cetakkeun(nilai[2])

# 3. Ngagunakeun Kamus 
cetakkeun("--- Tes Kamus ---")
tanda identitas = {
    "ngaran": "Kabayan",
    "asal": "Pandeglang",
    "umur": 30
}

cetakkeun("Ngaran:")
cetakkeun(identitas["ngaran"])
cetakkeun("Asal:")
cetakkeun(identitas["asal"])

# 4. Kombinasi
cetakkeun("--- Tes Itungan ---")
tanda total = tambah(nilai[0], identitas["umur"])
cetakkeun("80 + 30 =")
cetakkeun(total)