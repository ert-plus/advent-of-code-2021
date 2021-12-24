package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (mr MyReader) {

}

func main() {
	reader.Validate(MyReader{})
}