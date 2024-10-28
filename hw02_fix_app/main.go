package main

import (
	"fmt"

	"github.com/Denis9111/otus_hw_go_basic/hw02_fix_app/printer"
	"github.com/Denis9111/otus_hw_go_basic/hw02_fix_app/reader"
	"github.com/Denis9111/otus_hw_go_basic/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	printer.PrintStaff(staff)
}
