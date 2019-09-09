package main

import (
	"fmt"
)

// verilen sayının tüm asal çarpanlarını al n

func PrimeFactors(n int) (pfs []int) {
	// N bölen 2s sayısını al
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// Bu noktada n tuhaf olmalı. böylece bir elementi atlayabiliriz
	// (note i = i + 2)
	var i int
	for i = 3; i*i <= n; i = i + 2 {
		// n'yi böldüğümde, n'i ekle ve i böl
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// Bu koşul, n bir asal sayı olduğunda durumu ele almaktır.
	// 2'den büyük
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func main() {
	fmt.Println(PrimeFactors(600851475143))
}
