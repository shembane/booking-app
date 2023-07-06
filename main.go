package main // This line organises the application into a package

import "fmt" // All functions in Go come from a package, in this case the Print function comes from the fmt package (fmt is a core Go formatting package)

func main() { // This function tells the compiler where to start running the application's code. There can only be one of these 'main' functions.
	fmt.Println("Hello world") // You must add the package name before a funtion to tell the app where that function comes from, in this case Print or Println (ln writes the text in a new line)
}
