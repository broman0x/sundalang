tanda angka = tanyakeun("Masukan Angka Ganjil/Genap: ")
cetakkeun("Angka nu dicek: ")
cetakkeun(angka)

tanda sisa = angka
kedap sisa > 1 {
tanda sisa = sisa - 2
}

lamun sisa == 0 {
    cetakkeun("ini teh genap")
} lamunteu {
    cetakkeun("ini teh ganjil")
}