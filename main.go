package main

import "fmt"

func main() {
	fmt.Println("demo golang wire")
	fmt.Printf("[%#v]\n", InitializeSetProvideFirstService())
	fmt.Printf("[%#v]\n", InitializeSetProvideBindService())
	fmt.Printf("[%#v]\n", InitializeSetProvideStructSummary())
	fmt.Printf("[%#v]\n", InitializeSetValueStructFirst())
	fmt.Printf("[%#v]\n", InitializeSetValueStructSecond())
	fmt.Printf("[%#v]\n", InitializeSetValueStructSummary())
	fmt.Printf("[%#v]\n", InitializeUser("Name", 1, "Message"))
	if database, err := InitializeDatabaseWithError(""); err == nil {
		fmt.Printf("[%#v]\n", database)
	} else {
		fmt.Printf("[%#v]\n", err)
	}
	if database, cleanup, err := InitializeDatabaseWithErrorAndCleanup(""); err == nil {
		fmt.Printf("[%#v]\n", database)
		cleanup()
	} else {
		fmt.Printf("[%#v]\n", err)
	}
}
