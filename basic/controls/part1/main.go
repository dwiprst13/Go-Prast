package main

import ("fmt")

func kelulusan() {
	var input1 int
	var input2 int

	fmt.Println("Masukkan nilai ujian:")
	fmt.Scan(&input1)
	fmt.Println("Masukkan nilai tugas:")
	fmt.Scan(&input2)
	var grade string

	if input1 >= 85 && input2 >= 85 {
		grade = "A"
	} else if input1 >= 70 && input2 >= 70 {
		grade = "B"
	} else if input1 >= 60 && input2 >= 60 {
		grade = "C"
	} else if input1 >= 50 && input2 >= 50 {	
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println("Nilai ujian:", input1)
	fmt.Println("Nilai tugas:", input2)
	fmt.Println("Grade:", grade)
	fmt.Println("Kelulusan:", grade)
	if grade == "A" || grade == "B" {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

}

func main() {
	kelulusan()
}