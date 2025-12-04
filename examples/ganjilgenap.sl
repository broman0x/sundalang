# Cek Ganjil Genap make Operator Modulo (%)

cetakkeun("=== Cek Ganjil Genap ===")
tanda angka = tanyakeun("Masukan Angka: ")

cetakkeun("Angka nu dicek: ")
cetakkeun(angka)

# Make operator % (sisa bagi) nu leuwih gampang
# Di versi heubeul kudu make looping, ayeuna mah langsung weh
tanda sisa = angka % 2

lamun sisa == 0 {
    cetakkeun("Hasilna: Ieu teh angka GENAP.")
} lamunteu {
    cetakkeun("Hasilna: Ieu teh angka GANJIL.")
}