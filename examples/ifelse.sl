cetakkeun("--- Program Peunteun (Nilai) Sakola ---")

tanda nilai = tanyakeun("Masuken Nilai Maneh: ")
cetakkeun("Nilai maneh: " + nilai )

# Logika Percabangan Bertingkat
# lamun          -> if
# lamunteu lamun -> else if
# lamunteu       -> else

lamun nilai >= 90 {
    cetakkeun("Grade: A (Hade Pisan)")
    cetakkeun("Pertahankeun prestasina!")

} lamunteu lamun nilai >= 80 {
    cetakkeun("Grade: B (Hade)")
    cetakkeun("Saeutik deui meunang A.")

} lamunteu lamun nilai >= 70 {
    cetakkeun("Grade: C (Cukup)")
    cetakkeun("Kudu leuwih giat diajar.")

} lamunteu lamun nilai >= 60 {
    cetakkeun("Grade: D (Kurang)")
    cetakkeun("Awas ulah loba teuing ulin.")

} lamunteu {
    cetakkeun("Grade: E (Gagal)")
    cetakkeun("Wayahna kudu ngulang taun hareup.")
}

