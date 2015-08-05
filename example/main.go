package main

import (
	"fmt"
	"github.com/conc/goconfig"
)

func main() {

	testConfig := goconfig.LoadConfig("../testconfig/test.ini")

	fmt.Println(testConfig.GetString("user"))
	fmt.Println(testConfig.GetString("passwd"))
	fmt.Println(testConfig.GetInt("id"))
	fmt.Println(testConfig.GetBool("key5"))

}
