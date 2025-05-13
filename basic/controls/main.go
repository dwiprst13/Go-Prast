package main

import ("fmt")

func ifelse() {
	var nilai int = 60
	var grade string

	if nilai >= 85 {
		grade = "A"
	} else if nilai >= 70 {
		grade = "B"
	} else if nilai >= 60 {
		grade = "C"
	} else if nilai >= 50 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println("Nilai:", nilai)
	fmt.Println("Grade:", grade)
}

func switchcase() {
	var nilai int = 60
	var grade string

	switch {
	case nilai >= 85:
		grade = "A"
	case nilai >= 70:
		grade = "B"
	case nilai >= 60:
		grade = "C"
	case nilai >= 50:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Println("Nilai:", nilai)
	fmt.Println("Grade:", grade)
}

func forfunction() {
	for i := 2; i <= 4; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	ifelse()
	switchcase()
	forfunction()
}