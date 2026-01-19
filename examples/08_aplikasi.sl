# 08_aplikasi.sl - Aplikasi Lengkep

# Aplikasi Pencatat Tugas (To-Do List Sederhana)

tanda tugas = ["Diajar Golang", "Nyobaan SundaLang"]
tanda jalan = bener

tanda tampilkeun_menu = pungsi() {
    cetakkeun("\n=== MENU TUGAS ===")
    cetakkeun("1. Tampilkeun Tugas")
    cetakkeun("2. Tambah Tugas")
    cetakkeun("3. Keluar")
    cetakkeun("==================")
}

cetakkeun("Wilujeng Sumping di Aplikasi Tugas!")

kedap jalan {
    tampilkeun_menu()

    # Demo Loop otomatis
    pikeun tanda i = 0; i < 3; i = i + 1 {
        milih i {
            kasus 0:
                cetakkeun("\n[Aksi] Nampilkeun Tugas:")
                # Loop pikeun nampilkeun array
                pikeun tanda j = 0; j < panjang(tugas); j = j + 1 {
                    cetakkeun("- " + tugas[j])
                }
            kasus 1:
                cetakkeun("\n[Aksi] Nambah Tugas Baru...")
                tanda random_id = acak(100)
                tugas = asupkeun(tugas, "Tugas Anyar " + random_id)
                cetakkeun("Tugas ditambihkeun!")
            kasus 2:
                cetakkeun("\n[Aksi] Keluar...")
                jalan = salah
        }
        reureuh(500) # Jeda satengah detik
    }
}

cetakkeun("Hatur nuhun tos nganggo aplikasi ieu!")
