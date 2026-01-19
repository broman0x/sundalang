# 04_pungsi.sl - Pungsi (Functions)

# 1. Pungsi Biasa
tanda sapa = pungsi(ngaran) {
    cetakkeun("Wilujeng sumping, " + ngaran + "!")
}

sapa("Asep")

# 2. Pungsi jeung Balikkeun (Return)
# "pungsi" oge bisa diganti ku "fungsi"
tanda kali = fungsi(a, b) {
    balikkeun a * b
}

tanda hasil = kali(5, 4)
cetakkeun("Hasil 5 x 4 = " + hasil)

# 3. Pungsi salaku Argumen (Callback)
tanda jalankeun = pungsi(fn) {
    cetakkeun("Ngajalankeun pungsi...")
    fn()
}

jalankeun(pungsi() {
    cetakkeun("Ieu pungsi di jero pungsi!")
})
