package main

import "fmt"

func main() {
	var size int
	fmt.Println("Введите размер доски: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil || size <= 0 {
		fmt.Println("введено не число или 0.\nвведите целое цисло > 0 ...")
		return
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
