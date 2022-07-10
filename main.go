package main

import (
	"fmt"
	"time"
	// "io"
	"bufio"
	"log"
	"os"
)

func init() {
	displayHello()
}

func expensiveOperation(c chan int) {
	c <- 4
	time.Sleep(5 * time.Second)
	c <- 2
}

func test() {
	c := make(chan int)
	go expensiveOperation(c)
	// fmt.Println(<-c)
	fmt.Println(<-c)
}

// TODO use decorator to keep only one file open
type Actions = map[string]func(string)

func myPrint(str string) {
	fmt.Println(str)
}

func getFileForAppending(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func saveLineToFile(line string) {
	f, _ := getFileForAppending("./example2.txt")
	defer f.Close()
	f.WriteString(line)
}

func main() {
	displayHelp()

	// actions := Actions{
	// 	"w": myPrint,
	// 	"a": myPrint,
	// 	"s": myPrint,
	// 	"d": myPrint,
	// 	"q": myPrint,
	// 	"e": myPrint,
	// }
	saveLineToFile("test line")

	// processFile("./example.txt",
	// 	func(str string) { processLine(actions, str) })
}

func processLine(act Actions, line string) {
	fmt.Println()
	fmt.Println(line)
	choice := readLine()
	if command, ok := act[choice]; !ok {
		displayHelp()
		processLine(act, line)
	} else {
		command(line)
	}

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

func displayHelp() {
	fmt.Println(
		` use w,a,s,d,q,e to do something
		`)
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
	fmt.Scan(&str)
	return str
}
