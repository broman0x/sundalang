# 06_io_error.sl - File I/O & Error Handling

# 1. Nyerat & Maca File
tanda file = "catatan.txt"
cetakkeun("Nyerat data ka file " + file + "...")

cobaan {
    nyerat(file, "Ieu catetan rahasiah SundaLang.")
    tanda eusi = maca(file)
    cetakkeun("Eusi file: " + eusi)
} sanya {
    cetakkeun("Gagal maca/nulis file!")
}

# 2. Cobaan (Try-Catch)
cetakkeun("\nNgetes Error Handling...")
cobaan {
    tanda a = 10
    tanda b = 0
    # Ieu bakal error (integer divide by zero)
    tanda c = a / b
    cetakkeun("Ieu moal tampil")
} sanya {
    cetakkeun("Katerima Error: Teu bisa ngabagi jeung nol!")
}
