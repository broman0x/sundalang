// Contoh Pungsi Waktu & Jeda (Sleep)

// Ningali waktu ayeuna
cetakkeun("Waktu ayeuna: " + waktu())

cetakkeun("Ngagoan 2 detik (nganggo sare)...")
// sare(milidetik) -> 2000ms = 2 detik
sare(2000)
cetakkeun("Hudang sare! Waktu ayeuna: " + waktu())

cetakkeun("Reureuh heula 1 detik (nganggo reureuh)...")
// reureuh sami sareng sare (alias)
reureuh(1000)
cetakkeun("Lanjut deui! Waktu ayeuna: " + waktu())

// Countdown
cetakkeun("Itungan mundur mimitian:")
pikeun tanda i = 3; i > 0; i = i - 1 {
    cetakkeun(i + "...")
    sare(1000)
}
cetakkeun("DUAR!!! ")
