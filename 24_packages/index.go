package main

import (
	"github.com/fatih/color"
	"github.com/iarpitnagpure/GoLearning/auth"
)

// To create mod file use go mod init github.com/iarpitnagpure/GoLearning, This file includes all your dependencies
// To Install any third party dependency use go get <library_path>
// This will add dependency in go.mod and go.sum is like package-lock file includes internal depedency
func main() {
	auth.ValidateLoginInfo("TestUser", "123") // "username, password ---->" TestUser 123
	color.Cyan("Prints text in cyan.")        // "Prints text in cyan."
}
