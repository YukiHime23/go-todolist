package main

import "os"

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()
}

func Server()  {
	
}