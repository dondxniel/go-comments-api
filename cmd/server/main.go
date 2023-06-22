package main

import "fmt"

// * The Run() function is going to be
// * responsible for instantiation and
// * startup of our Go server
func Run() error {
	fmt.Println("Starting up go server")
	return nil
}

func main(){
	fmt.Println("Running")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}