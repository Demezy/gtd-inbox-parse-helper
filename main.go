package main

import (
	"fmt"
	// "io"
	"bufio"
	"log"
	"os"
)

func init() {
	fmt.Println("Init")
}

func main() {
	displayHello()
	processFile("./example.txt", func(s string) { fmt.Println(s) })
}

func processFile(filename string, callback func(string)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		callback(line)
	}

	return nil
}

func readEntireFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)

}

func displayHello() {
	fmt.Println(`  ___           _                        ____
 |_ _|  _ __   | |__     ___   __  __   |  _ \    __ _   _ __   ___    ___   _ __
  | |  | '_ \  | '_ \   / _ \  \ \/ /   | |_) |  / _\ | | '__| / __|  / _ \ | '__|
  | |  | | | | | |_) | | (_) |  >  <    |  __/  | (_| | | |    \__ \ |  __/ | |
 |___| |_| |_| |_.__/   \___/  /_/\_\   |_|      \__,_| |_|    |___/  \___| |_|`)
}

func getHello(num int) string {
	hellos := []string{
		"Hi, aboba",
		"Hello, biba",
		"Dobryi den, milsdar",
	}
	return (hellos[num%(len(hellos))])
}

func readLine() string {
	var str string
	fmt.Println("Enter string")
	fmt.Scan(&str)
	return str
}
