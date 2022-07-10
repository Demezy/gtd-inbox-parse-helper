package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	displayHello()
	os.Mkdir(prefix, 0755)
}

type Actions = map[string]func(string)

func getFileForAppending(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func getWriteLineCallback(filename string) func(string) {
	f, err := getFileForAppending(filename)
	if err != nil {
		panic(err)
	}
	return func(str string) {
		f.WriteString(str)
		f.WriteString("\n")
	}
}

const prefix string = "output/"

func main() {
	displayHelp()

	actions := Actions{}
	setupHotKeys(actions)

	processFile("./example.txt",
		func(str string) { processLine(actions, str) })
}

func setupHotKeys(acts Actions) {
	acts["w"] = getWriteLineCallback(prefix + "w.txt")
	acts["a"] = getWriteLineCallback(prefix + "a.txt")
	acts["s"] = getWriteLineCallback(prefix + "s.txt")
	acts["d"] = getWriteLineCallback(prefix + "d.txt")
	acts["q"] = getWriteLineCallback(prefix + "q.txt")
	acts["e"] = getWriteLineCallback(prefix + "e.txt")
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

func displayHello() {
	fmt.Println(`  ___           _                        ____
 |_ _|  _ __   | |__     ___   __  __   |  _ \    __ _   _ __   ___    ___   _ __
  | |  | '_ \  | '_ \   / _ \  \ \/ /   | |_) |  / _\ | | '__| / __|  / _ \ | '__|
  | |  | | | | | |_) | | (_) |  >  <    |  __/  | (_| | | |    \__ \ |  __/ | |
 |___| |_| |_| |_.__/   \___/  /_/\_\   |_|      \__,_| |_|    |___/  \___| |_|`)
}

func readLine() string {
	var str string
	fmt.Scan(&str)
	return str
}
