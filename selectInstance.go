package main

import (
	"fmt"
	"strconv"

	"sshaws/helpers"
)

func selectInstanceIndex(instance_list []helpers.Instance) int {
	var input string
	var index int
	if len(instance_list) == 1 {
		fmt.Printf("\n\n")
		index = 0
	} else {
		fmt.Println("\nWhich one do you want to ssh in?")
		fmt.Scanln(&input)
		index, err := strconv.Atoi(input)
		if err != nil || index > len(instance_list)-1 || index < 0 {
			fmt.Println("Please enter a valid number.")
			index = selectInstanceIndex(instance_list)
		}
	}
	return index
}
