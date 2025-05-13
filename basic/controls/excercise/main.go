package main

import (
	"fmt"
	"os"
	"bufio"
)

type Task struct {
	title string
	desc string
	done bool
}
type TodoList struct {
	tasks []Task
}

func scanInput() string {
	var input string
	bufioReader := bufio.NewReader(os.Stdin)
	input, _ = bufioReader.ReadString('\n')
	return input
}

func addTask(todolist *TodoList) {
	var title, desc string
	fmt.Println("masukkan Judul Tugas:")
	title = scanInput()
	fmt.Println("masukkan Deskripsi Tugas:")
	desc = scanInput()
	todolist.tasks = append(todolist.tasks, Task{title: title, desc: desc, done: false})
	fmt.Println("Tugas berhasil ditambahkan")
}

func listTask(todolist *TodoList ) {
	if len(todolist.tasks) == 0 {
		fmt.Println("Tidak ada tugas yang tersedia")
		return
	} else {
		fmt.Println("Daftar Tugas:")
		for i, task := range todolist.tasks {
			fmt.Println(i+1, "Judul:", task.title)
			fmt.Println("Deskripsi:", task.desc)
			if task.done {
				fmt.Println("Status: Selesai")
			} else {
				fmt.Println("Status: Belum selesai")
			}
			fmt.Println()
		}
	}
}

func markTaskAsDone(todolist *TodoList) {
	var index int
	fmt.Println("Masukkan nomor tugas yang ingin ditandai sebagai selesai:")
	fmt.Scanln(&index)
	if index < 1 || index > len(todolist.tasks) {
		fmt.Println("Nomor tugas tidak valid")
		return
	} else {
		todolist.tasks[index-1].done = true
		fmt.Println("Tugas berhasil ditandai sebagai selesai")
	}
}

func deleteTask(todolist *TodoList) {
	var index int
	fmt.Println("Pilih nomor tugas yang ingin dihapus:")
	fmt.Scanln(&index)
	if index < 1 || index > len (todolist.tasks) {
		fmt.Println("Nomor tugas tidak valid")
		return
	} else {
		todolist.tasks = append(todolist.tasks[:index-1], todolist.tasks[index:]...) // ternyata Go manipulatif kek dia bjir
		fmt.Println("Tugas berhasil dihapus")
	}
}

func main() {
	var todolist TodoList
	var choice int
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Daftar Tugas")
		fmt.Println("3. Tandai Tugas sebagai Selesai")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)
		fmt.Println()
		switch choice {
		case 1:
			addTask(&todolist)
		case 2:
			listTask(&todolist)
		case 3:
			markTaskAsDone(&todolist)
		case 4:
			deleteTask(&todolist)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
		fmt.Println()
	}
}
