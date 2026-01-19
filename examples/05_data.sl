# 05_data.sl - Struktur Data (Array & Wadah)

# 1. Array (Daptar)
tanda buah = ["Cau", "Anggur", "Jeruk"]
cetakkeun("Buah kahiji: " + mimiti(buah))
cetakkeun("Buah panungtung: " + tungtung(buah))
cetakkeun("Jumlah buah: " + kana_tulisan(panjang(buah)))

# Nambahkeun data
tanda buah_anyar = asupkeun(buah, "Manggu")
cetakkeun("Buah anyar: " + tungtung(buah_anyar))

# 2. Wadah (Struct/Map)
# Wadah bisa ditulis siga JSON objek
tanda mahasiswa = wadah {
    "NIM": 123456,
    "Jurusan": "Informatika",
    "Aktif": bener
}

cetakkeun("Jurusan: " + mahasiswa["Jurusan"])
cetakkeun("Status Aktif: " + mahasiswa["Aktif"])
