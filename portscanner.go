package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func worker(address string, ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		target := net.JoinHostPort(address, fmt.Sprintf("%d", p))
		conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)

		if err != nil {
			wg.Done()
			continue
		}

		conn.Close()
		fmt.Printf("[+] Port %d terbuka\n", p)
		wg.Done()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan address yang ingin discan: ")
	scanner.Scan()
	address := scanner.Text()

	// Gunakan buffer yang sama dengan jumlah worker (1000)
	portChannel := make(chan int, 1000)
	var wg sync.WaitGroup

	fmt.Printf("Memulai scanning pada %s...\n", address)
	start := time.Now()

	// 1. Jalankan 1000 Worker
	for i := 0; i < 1000; i++ {
		go worker(address, portChannel, &wg)
	}

	// 2. Kirim port (Maksimal 65535)
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		portChannel <- i
	}

	// 3. Tunggu proses selesai
	wg.Wait()

	// 4. Tutup channel
	close(portChannel)

	duration := time.Since(start)
	fmt.Printf("\nSelesai dalam waktu: %v\n", duration)
}
