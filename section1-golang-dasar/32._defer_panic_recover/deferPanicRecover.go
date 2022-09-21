package main

import "fmt"

/**
defer function = adalah function yg bisa dijadwalkan untuk dieksekusi setelah sebuah func di eksekusi
defer function akan selalu di eksekusi walaupun terjadi error di function yg di eksekusi
*/

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

/**
Panic function adalah function yang bisa digunakan untuk menghentikan program
Panic func biasanya dipanggil ketika terjadi error saat program berjalan
Saat panic func dipanggil, program terhenti, namun defer function tetap dieksekusi
*/

func endApp() {
	message := recover() // recover berfungsi untuk menangkap data panic, dengan recover proses panic akan terhenti, sehingga program tetap berjalan
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi Selesai")

}

func runApp(error bool) {
	defer endApp()
	if error {
		fmt.Println("Aplikasi berjalan")
	}else{
		panic("Aplikasi ERROR")
	}
}

func main() {
	// runApplication(0)
	runApp(false)
	fmt.Println("EKo")
}