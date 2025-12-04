# Interaksi Input Output (IO) & Kondisional

cetakkeun("=== Formulir Pendataan Warga ===")

# 1. Input String
tanda ngaran = tanyakeun("Saha aran maneh? ")
tanda salam = "Sampurasun, " + ngaran
cetakkeun(salam)

# 2. Input Integer
tanda umur = tanyakeun("Sabaraha umur maneh? ")

# Interpreter SundaLang otomatis ngadeteksi mun inputna angka
lamun umur >= 17 {
    cetakkeun("Status: Geus boga KTP.")
    cetakkeun("Sok atuh geura nyieun SIM.")
} lamunteu {
    cetakkeun("Status: Budak keneh.")
    cetakkeun("Ulah waka udud siah!")
}