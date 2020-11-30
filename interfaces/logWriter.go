package main

import "fmt"

type logWriter struct{}

func (logWriter) Write(buff []byte) (int, error) {
	fmt.Println(string(buff))
	return len(buff), nil
}
