# === PANDUAN TIPE DATA SUNDALANG v1.0.1 ===

# 1. INTEGER (Bilangan Bulat)
# Biasa dipake jang itungan
tanda umur = 25
tanda taun = 2025
tanda hasil = 10 + 5 * 2
cetakkeun("--- Integer ---")
cetakkeun(umur)
cetakkeun(hasil)

# 2. STRING (Téks)
# Diapit ku kutip dua ("")
# Bisa digabung make tanda tambah (+)
tanda ngaran = "Kabayan"
tanda salam = "Wilujeng Enjing, " + ngaran
cetakkeun("--- String ---")
cetakkeun(salam)

# 3. BOOLEAN (Logika)
# Ngan aya dua: bener atawa salah
tanda lulus = bener
tanda gagal = salah
tanda cek = 10 > 5  # Ieu hasilna bener
cetakkeun("--- Boolean ---")
cetakkeun(lulus)
cetakkeun(cek)

# 4. ARRAY (Daptar) [Anyar!]
# Kumpulan data nu ngaruntuy. Indeks dimimitian ti 0.
# Format: [data1, data2, data3]
tanda hobi = ["Mancing", "Sare", "Ngoding"]
tanda angka = [10, 20, 30]

cetakkeun("--- Array ---")
cetakkeun("Hobi ka-1:")
cetakkeun(hobi[0])  # Akses elemen kahiji
cetakkeun("Hobi ka-3:")
cetakkeun(hobi[2])  # Akses elemen katilu
cetakkeun("Jumlah angka:")
cetakkeun(angka[0] + angka[1]) # 10 + 20

# 5. HASH MAP 
# Kumpulan data nu boga konci (key).
# Format: {"konci": nilai}
tanda identitas = {
    "ngaran": "Iteung",
    "alamat": "Bandung",
    "umur": 22,
    "menikah": salah
}

cetakkeun("--- Hash Map ---")
cetakkeun("Ngaran:")
cetakkeun(identitas["ngaran"]) # Akses make konci string
cetakkeun("Umur:")
cetakkeun(identitas["umur"])
cetakkeun("Status:")
lamun identitas["menikah"] == bener {
    cetakkeun("Geus kawin")
} lamunteu {
    cetakkeun("Jomblo keneh")
}

# 6. KOMBINASI (Complex)
# Hash di jero Array, atawa sabalikna
tanda data_kompleks = [
    {"id": 1, "ngaran": "Asep"},
    {"id": 2, "ngaran": "Ujang"}
]

cetakkeun("--- Kombinasi ---")
tanda user_kahiji = data_kompleks[0]
cetakkeun(user_kahiji["ngaran"]) # Output: Asep