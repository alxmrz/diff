package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: diff file1 file2")
	}

	firstFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	secondFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	scannerFirst := bufio.NewScanner(firstFile)
	scannerSecond := bufio.NewScanner(secondFile)

	runFirstScan := true
	runSecondScan := true

	i := 0
	for {
		i++
		if !runFirstScan && !runSecondScan {
			break
		}

		if runFirstScan && !scannerFirst.Scan() {
			runFirstScan = false
		}

		if runSecondScan && !scannerSecond.Scan() {
			runSecondScan = false
		}

		firstFileLine := scannerFirst.Text()
		secondFileLine := scannerSecond.Text()

		if runFirstScan && runSecondScan {
			if firstFileLine == secondFileLine {
				printColoredLine(colorGreen, i, firstFileLine, firstFile.Name())
			} else {
				printColoredLine(colorRed, i, firstFileLine, firstFile.Name())
				printColoredLine(colorRed, i, secondFileLine, secondFile.Name())
			}
		} else {
			if runFirstScan {
				printColoredLine(colorRed, i, firstFileLine, firstFile.Name())
			}
			if runSecondScan {
				printColoredLine(colorRed, i, secondFileLine, secondFile.Name())
			}
		}
	}
}

func printColoredLine(color string, index int, line string, file string) {
	fmt.Print(color)
	fmt.Print(file, ":", index, "\n")
	fmt.Println(line)
	fmt.Print(colorReset)
}
