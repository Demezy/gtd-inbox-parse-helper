package main

import "fmt"

func init(){
	fmt.Println("Init")
}

func main(){
	displayHello()
}

func displayHello(){
	fmt.Println(`  ___           _                        ____
 |_ _|  _ __   | |__     ___   __  __   |  _ \    __ _   _ __   ___    ___   _ __
  | |  | '_ \  | '_ \   / _ \  \ \/ /   | |_) |  / _\ | | '__| / __|  / _ \ | '__|
  | |  | | | | | |_) | | (_) |  >  <    |  __/  | (_| | | |    \__ \ |  __/ | |
 |___| |_| |_| |_.__/   \___/  /_/\_\   |_|      \__,_| |_|    |___/  \___| |_|`)
}


func getHello(num int) string{
	hellos := []string{
		"Hi, aboba",
		"Hello, biba" ,
		"Dobryi den, milsdar",
	}
	return(hellos[num % (len(hellos))])
}

func readLine() string{
	var str string
	fmt.Println("Enter string")
	fmt.Scan(&str)
	return str
}
