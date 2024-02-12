package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// ?
// create a method on the autehnticationInfo struct called getBAsicAuth that returns 
//the formatted string.
func (authInfo authenticationInfo) getBasicAuth() string {
	return ("Authorization: Basic " + authInfo.username + authInfo.password)
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
