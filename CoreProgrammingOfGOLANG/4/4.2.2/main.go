package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var i *int = nil
	fmt.Printf("%p", i)

}
func example1() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var i io.Reader = f
	switch v := i.(type) {
	case io.ReadWriter:
		v.Write([]byte("io.ReadWriter\n"))
	case *os.File:
		v.Write([]byte("*os.File\n"))
		v.Sync()

	default:
		return
	}

}
func example2() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var i io.Reader = f
	switch v := i.(type) {
	case *os.File:
		v.Write([]byte("*os.File\n"))
		v.Sync()
	case io.ReadWriter:
		v.Write([]byte("io.ReadWriter\n"))
	default:
		return
	}

}
func example3() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var i io.Reader = f
	switch v := i.(type) {
	case *os.File, io.ReadWriter:
		if v == i {
			fmt.Println(true)
		}
	default:
		return
	}
}
