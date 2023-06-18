package main

import "fmt"

// * This function is responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application");
	return nil
}

func main(){
	fmt.Println("Hello from main.go")
	if err := Run(); err != nil {
		// * Handle error
		fmt.Println(err);
	}
}