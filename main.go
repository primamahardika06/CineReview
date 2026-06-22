package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CineReview struct{
	judulFilm string
	genre string
	tahunRilis int
	deskripsi string
	skorFilm float64
}

func tampilkanData(data []CineReview){
	if len(data) == 0{
		fmt.Println("Belum ada data")
		return
	}
	for i := 0; i < len(data); i++{
		fmt.Printf("%d. Judul: %s | Genre: %s | Tahun: %d | Deskripsi: %s | Skor: %.1f\n", i+1, data[i].judulFilm, data[i].genre, data[i].tahunRilis, data[i].deskripsi, data[i].skorFilm)
	}
}
func seqSearchJudul(data []CineReview, target string) bool {
	if len(data) == 0 {
		fmt.Println("Tidak data film untuk dicari.")
		return false
	}

	j:= 0
	found := false
	nobaru:= 0

	fmt.Println("Hasil Pencarian:")
	for j < len(data) {
		if data[j].judulFilm == target {
			found = true
			nobaru++
			fmt.Printf("%d. Judul: %s | Genre: %s | Tahun: %d | Deskripsi: %s | Skor: %.1f\n", nobaru, data[j].judulFilm, data[j].genre, data[j].tahunRilis, data[j].deskripsi, data[j].skorFilm)
		}
		j++
	}
	return found
}

func seqSearchGenre(data []CineReview, target string) bool {
	if len(data) == 0 {
		fmt.Println("Tidak data film untuk dicari.")
		return false
	}

	j:= 0
	found := false
	nobaru:= 0
	fmt.Println("Hasil Pencarian:")
	for j < len(data) {
		if data[j].genre == target{
			found = true
			nobaru++
			fmt.Printf("%d. Judul: %s | Genre: %s | Tahun: %d | Deskripsi: %s | Skor: %.1f\n", nobaru, data[j].judulFilm, data[j].genre, data[j].tahunRilis, data[j].deskripsi, data[j].skorFilm)
		}
		j++
	}
	return found


}

func binarySearchJudul(data []CineReview, k string) int{
	kiri := 0
	kanan := len(data) -1
	
	for kiri <= kanan{
		mid := (kiri + kanan) / 2
		Mid := data[mid].judulFilm
		if Mid == k{
			return mid
		}else if Mid < k{
			kiri = mid + 1
		}else{
			kanan = mid -1
		}
	}
	return -1
}
func binarySearchGenre(data []CineReview, k string) int{
	kiri := 0
	kanan := len(data) -1
	
	for kiri <= kanan{
		mid := (kiri + kanan) / 2
		Mid := data[mid].genre
		if Mid == k{
			return mid
		}else if Mid < k{
			kiri = mid + 1
		}else{
			kanan = mid -1
		}
	}
	return -1
}


func sortingByJudul(data []CineReview) []CineReview{
	diurutkan := make([]CineReview, len(data))
	for i := 0; i < len(data); i++{
		diurutkan[i] = data[i]
	}
	n := len(diurutkan)
	for i := 0; i < n - 1; i++{
		for j := 0; j < n-i-1; j++{
			if diurutkan[j].judulFilm > diurutkan[j+1].judulFilm {
				diurutkan[j], diurutkan[j+1] = diurutkan[j+1], diurutkan[j]
			}
		} 
	}
	return diurutkan
}

func sortingByGenre(data []CineReview) []CineReview{
	diurutkan := make([]CineReview, len(data))
	for i := 0; i < len(data); i++{
		diurutkan[i] = data[i]
	}
	n := len(diurutkan)
	for i := 0; i < n - 1; i++{
		for j := 0; j < n-i-1; j++{
			if diurutkan[j].genre > diurutkan[j+1].genre {
				diurutkan[j], diurutkan[j+1] = diurutkan[j+1], diurutkan[j]
			}
		} 
	}
	return diurutkan
}

 func selectionsortRating(T []CineReview, n int) {
 	var i,j, idx_max int
	var t CineReview
	i = 1
	for i <= n-1 {
		idx_max = i-1 
		j=i
		for j < n {
			if T[idx_max].skorFilm < T[j].skorFilm {
				idx_max = j
			}
			j++
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i++
	}
 }

 func selectionsortTahun(T []CineReview, n int) {
	var i,j, idx_max int
	var t CineReview
	i = 1
	for i <= n-1 {
		idx_max = i-1 
		j=i
		for j < n {
			if T[idx_max].tahunRilis < T[j].tahunRilis {
				idx_max = j
			}
			j++
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i++
	}
 }

func insertionSortTahun(data []CineReview, n int) {
    var i, j int
    i = 1
    for i <= n-1 {
        j = i
        temp := data[j]
        for j > 0 && temp.tahunRilis > data[j-1].tahunRilis {

            data[j] = data[j-1]
            j--
        }
        data[j] = temp
        i++
    }
}

func insertionSortRating(data []CineReview, n int) {
    var i, j int
    i = 1
    for i <= n-1 {
        j = i
        temp := data[j]
        for j > 0 && temp.skorFilm > data[j-1].skorFilm {

            data[j] = data[j-1]
            j--
        }
        data[j] = temp
        i++
    }
}

func jumlahPerGenre(data []CineReview, genreDicari string) int{
	total := 0

	for i:= 0; i < len(data); i++{
		if data[i].genre == genreDicari{
			total++
			fmt.Printf("%d. Judul: %s | Genre: %s | Tahun: %d | Deskripsi: %s | Skor: %.1f\n",total, data[i].judulFilm, data[i].genre, data[i].tahunRilis, data[i].deskripsi, data[i].skorFilm)
		}
	}
	return total

}

func ratarating(data []CineReview) {
	var totalRating float64	
	for i:= 0; i< len(data) ;i++ {
		totalRating+= data[i].skorFilm
	}
	ratarata:= totalRating / float64(len(data))
	fmt.Printf("Total Film : %d\n", len(data))
	fmt.Printf("Rata-rata Rating: %.1f\n",ratarata)
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var data []CineReview
	fmt.Println("Selamat Datang di CineReview🙌")
	for{
		fmt.Println("=======================")
		fmt.Println(" 1. Menambahkan")
		fmt.Println(" 2. Menghapus")
		fmt.Println(" 3. Mengubah")
		fmt.Println(" 4. Tampilkan Film")
		fmt.Println(" 5. Cari Film")
		fmt.Println(" 6. Urutkan Film")
		fmt.Println(" 7. Statistik")
		fmt.Println(" 0. Keluar")
		fmt.Println("=======================")

		var input int
		fmt.Print("Input Pilihan Anda: ")
		fmt.Scan(&input)
		fmt.Println("=======================")


		switch input {
		case 1:
			var n int
			fmt.Print("Masukkan Jumlah Film Yang Mau Ditambahkan: ")
			fmt.Scan(&n)
			fmt.Println("=======================")
			dataBaru := make([]CineReview, n)

			fmt.Scanln()

			for i := 0; i < n; i++{
				fmt.Printf("--- Film ke-%d ---\n", i+1)
				fmt.Print("Masukkan Judul Film: ")
				if scanner.Scan() {
					dataBaru[i].judulFilm = scanner.Text()
					dataBaru[i].judulFilm = strings.ToLower(scanner.Text())
				}

				fmt.Print("Masukkan Genre Film: ")
				fmt.Scan(&dataBaru[i].genre)
				
				fmt.Print("Masukkan Tahun Rilis: ")
				fmt.Scan(&dataBaru[i].tahunRilis)

				fmt.Scanln()
				
				fmt.Print("Masukkan Deskripsi Film: ")
				if scanner.Scan() {
					dataBaru[i].deskripsi = scanner.Text()
				}

				fmt.Print("Masukkan Skor Film (0-10): ")
				fmt.Scan(&dataBaru[i].skorFilm)

				fmt.Scanln()
				fmt.Println()
			}
			data = append(data, dataBaru...)
			fmt.Println("Data saat ini:")
			tampilkanData(data)
	
		case 2:
			tampilkanData(data)

			if len(data) == 0{
				fmt.Println("Tidak ada data untuk dihapus")
				continue
			}
			var hapus int
			fmt.Print("Data yang mau dihapus: ")
			fmt.Scan(&hapus)

			
			indexToRemove := hapus - 1
			if indexToRemove >= 0 && indexToRemove < len(data){
				data = append(data[:indexToRemove], data[indexToRemove+1:]...)
				fmt.Println("Data berhasil dihapus")
			}else{
				fmt.Println("Data tidak ditemukan")
			}
			
		case 3:
			tampilkanData(data)
			if len(data) == 0{
				fmt.Println("Tidak ada data untuk diubah")
				continue
			}

			var ubah int
			fmt.Print("ID Data yang mau diubah: ")
			fmt.Scan(&ubah)

			indexToChange := ubah - 1
			fmt.Scanln()

			if indexToChange >= 0 && indexToChange < len(data){
				fmt.Print("Ubah judul film: ")
				if scanner.Scan() {
					data[indexToChange].judulFilm = scanner.Text()
				}

				fmt.Print("Ubah genre film: ")
				fmt.Scan(&data[indexToChange].genre)

				fmt.Print("Ubah tahun rils film: ")
				fmt.Scan(&data[indexToChange].tahunRilis)

				fmt.Scanln()

				fmt.Print("Ubah deskripsi film: ")
				if scanner.Scan() {
					data[indexToChange].deskripsi = scanner.Text()
				}

				fmt.Print("Ubah skor film (0-10): ")
				fmt.Scan(&data[indexToChange].skorFilm)
				fmt.Println()

				fmt.Scanln()
				fmt.Println("Data berhasil diubah")
			}else{
				fmt.Println("Data tidak ditemukan")
			}
			
		case 4:
			tampilkanData(data)
			
		case 5:
			var kataKunci string
			// var kataKunci2 int
			var metode int
			var searchBy int

			fmt.Println("1. Cari berdasarkan Judul")
			fmt.Println("2. Cari berdasarkan Genre")
			fmt.Print("Cari berdasarkan: ")
			fmt.Scan(&searchBy)
			fmt.Scanln()
			if searchBy == 1{
				fmt.Print("Judul yang dicari: ")
				if scanner.Scan() {
        			kataKunci = strings.ToLower(scanner.Text())
    			}
			}else if searchBy == 2{
				fmt.Print("Genre yang dicari: ")
				// fmt.Scan(&kataKunci2)
				fmt.Scan(&kataKunci)
			}
			
			fmt.Println("1.BinarySearch")
			fmt.Println("2.Sequential Search")
			fmt.Print("Metode: ")
			fmt.Scan(&metode)
			fmt.Scanln()

			if metode == 1 && searchBy == 1 {
				terurut := sortingByJudul(data)
				idx := binarySearchJudul(terurut, kataKunci)
				if idx != -1{
					fmt.Println("Hasil Pencarian:")
					tampilkanData(terurut[idx:idx+1])
				}else{
					fmt.Println("Hasil Pencarian:")
					fmt.Println("Film dengan Judul tersebut tidak ditemukan")
				}	
			}else if metode == 1 && searchBy == 2{
				terurut := sortingByGenre(data)
				idx := binarySearchGenre(terurut, kataKunci)
				if idx != -1{
					fmt.Println("Hasil Pencarian:")
					awal := idx
					for awal > 0 && terurut[awal-1].genre == kataKunci{
						awal--
					}
					akhir := idx
					for akhir < len(terurut)-1 && terurut[akhir+1].genre == kataKunci{
						akhir++
					}
					tampilkanData(terurut[awal:akhir+1])
				}else{
					fmt.Println("Hasil Pencarian:")
					fmt.Println("Film dengan Genre tersebut tidak ditemukan")
				}	
			}else if metode == 2 && searchBy == 1{
				idx := seqSearchJudul(data, kataKunci)
				if !idx {
					fmt.Println("Film dengan judul tersebut tidak ditemukan")
				}
			}else if metode == 2 && searchBy == 2{
				idx := seqSearchGenre(data, kataKunci)
				if !idx {
					fmt.Println("Film dengan Genre tersebut tidak ditemukan")
				}
			} else {
				fmt.Println("Metode tersebut tidak ada ")
			}
		case 6:
			var pilihUrutan int
			var metode int
			fmt.Println("1. Urutkan berdasarkan tahun")
			fmt.Println("2. Urutkan berdasarkan rating")
			fmt.Print("Urutkan berdasarkan: ")
			fmt.Scan(&pilihUrutan)
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			fmt.Print("Metode: ")
			fmt.Scan(&metode)
			terurut := make([]CineReview, len(data))

			for i := 0; i < len(data); i++{
				terurut[i] = data[i]
			}
			if pilihUrutan == 1 && metode == 2{
				insertionSortTahun(terurut, len(terurut))
				tampilkanData(terurut)
			}else if pilihUrutan == 2 && metode == 2{
				insertionSortRating(terurut, len(terurut))
				tampilkanData(terurut)
			}else if pilihUrutan == 1 && metode == 1{
				selectionsortTahun(terurut, len(terurut))
				tampilkanData(terurut)
			}else if pilihUrutan == 2 && metode == 1{
				selectionsortRating(terurut, len(terurut))
				tampilkanData(terurut)
			}else{
				fmt.Println("Metode tersebut tidak ada ")
			}

		case 7:
			var pilih int
			fmt.Println("1. Jumlah film per genre")
			fmt.Println("2. Rata-rata rating seluruh film")
			fmt.Print("Tampilkan: ")
			fmt.Scan(&pilih)
			fmt.Println("===================================")
			if pilih == 1{
				var genreDicari string
				fmt.Print("Genre yang dicari: ")
				fmt.Scan(&genreDicari)
				jumlah := jumlahPerGenre(data, genreDicari)
				fmt.Println("===================================")
				fmt.Printf("Genre %s ditemukan sebanyak: %d\n", genreDicari, jumlah)
			}else if pilih == 2{
				fmt.Println("===================================")
				ratarating(data)
			}else{
				fmt.Println("Data tersebut tidak ditemukan")
			}
			

			
		case 0:
			fmt.Println("Terima kasih telah menggunakan CineReview! Sampai jumpa di review film selanjutnya 🎬👋")
			fmt.Println("===================================")
			return

		default:
			fmt.Println("Input tidak valid")
		}
	}

}




