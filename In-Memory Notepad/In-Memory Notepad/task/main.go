package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readByNewReader() {
	var reader = bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Enter a command and data: > ")
		args, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.ReplaceAll(string(args), "\n", "") == "exit" {
			fmt.Println("[Info] Bye!")
			break
		}

		fmt.Print(string(args))
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var notes []string
	var memoryLength int
	fmt.Print("Enter the maximum number of notes: > ")
	fmt.Scan(&memoryLength)

	for {
		fmt.Print("Enter a command and data: > ")
		args, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		var text = strings.ReplaceAll(string(args), "\n", "")
		splitString := strings.Split(text, " ")
		var command string
		if len(splitString) > 1 {
			command = splitString[0]
		} else {
			command = text
		}
		var isExit bool
		switch command {
		case "create":
			var data []string
			if len(splitString) > 1 {
				data = splitString[1:]
			}
			switch {
			case len(data) > 0:
				if len(notes) < memoryLength {
					notes = append(notes, strings.Join(data, " "))
					fmt.Println("[OK] The note was successfully created")
					break
				} else {
					fmt.Println("[Error] Notepad is full")
				}
			default:
				fmt.Println("[Error] Missing note argument")
			}
		case "list":
			if len(notes) == 0 {
				fmt.Println("[Info] Notepad is empty")
				break
			}
			for i, note := range notes {
				fmt.Printf("[Info] %d: %s\n", i+1, note)
			}
		case "update":
			if len(splitString) < 2 {
				fmt.Printf("[Error] Missing position argument\n")
				break
			} else if len(splitString) < 3 {
				fmt.Printf("[Error] Missing note argument\n")
				break
			}
			var id, err = strconv.Atoi(splitString[1])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splitString[1])
				break
			}

			if id > memoryLength+1 {
				fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", id, memoryLength)
				break
			}

			var idExist = false
			for i, _ := range notes {
				if i == id-1 {
					idExist = true
				}
			}
			if !idExist {
				fmt.Printf("[Error] There is nothing to update\n")
				break
			}

			notes[id-1] = strings.Join(splitString[2:], " ")
			fmt.Printf("[OK] The note at position %d was successfully updated\n", id)
		case "delete":
			if len(splitString) < 2 {
				fmt.Printf("[Error] Missing position argument\n")
				break
			}

			var id, err = strconv.Atoi(splitString[1])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splitString[1])
				break
			}
			if id > memoryLength+1 {
				fmt.Printf("[Error] Position %d is out of the boundary [1,%d]\n", id, memoryLength)
				break
			}
			var idExist = false
			for i, _ := range notes {
				if i == id-1 {
					idExist = true
				}
			}
			if !idExist {
				fmt.Printf("[Error] There is nothing to delete\n")
				break
			}

			notes[id-1] = notes[len(notes)-1]
			notes[len(notes)-1] = ""
			notes = notes[:len(notes)-1]
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", id)
		case "clear":
			notes = nil
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			isExit = true
			fmt.Println("[Info] Bye!")
		default:
			fmt.Println("[Error] Unknown command")
		}

		if isExit {
			break
		}
	}
}
