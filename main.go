package main

import (
	"bufio"
	"errors"
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
	actions := Actions{}
	setupHotKeys(actions)
	filename, err := getFilenameFromCli()
	if err != nil {
		displayHelp()
		return
	}

	processFile(filename,
		func(str string) { processLine(actions, str) })
}

func getFilenameFromCli() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Filename is not provided")
	}
	return os.Args[1], nil
}

var keybindings = map[string]string{
	"w": "todo.txt",
	"s": "someday.txt",
	"a": "calendar.txt",
	"d": "project.txt",
	"q": "deleted.txt",
	"e": "notes.txt",
	"f": "waiting.txt",
}

func setupHotKeys(acts Actions) {
	for key, file := range keybindings {
		acts[key] = getWriteLineCallback(prefix + file)
	}
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
	helpMessage := ` usage <program name :)> filename
Use keys quickly sort lines for files.`
	for key, file := range keybindings {
		helpMessage += fmt.Sprintf("\n\t%s - %s", key, file)
	}
	fmt.Println(helpMessage)
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
