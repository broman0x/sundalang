# 02_logika.sl - Logika & Percabangan

tanda nilai = 75

# 1. Lamun (If-Else)
lamun nilai >= 80 {
    cetakkeun("Hade Pisan! (A)")
} lamunteu nilai >= 70 {
    cetakkeun("Lumayan (B)")
} sanya {
    cetakkeun("Kudu diajar deui (C)")
}

# 2. Milih (Switch-Case)
tanda poe = "Salasa"

milih poe {
    kasus "Senin":
        cetakkeun("Upacara")
    kasus "Salasa":
        cetakkeun("Olahraga")
    kasus "Jumaah":
        cetakkeun("Jumatan")
    baku:
        cetakkeun("Belajar biasa")
}
