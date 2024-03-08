package main

import "fmt"

var PUBLIC_NAME string
var SITE_URL string
var SEO_STATUS uint8

func main() {
	fmt.Println("STARTED MAIN")
	fmt.Println(PUBLIC_NAME, SITE_URL, SEO_STATUS)
}
func init() {
	fmt.Println("Here we add values to all the initialized variables")
	PUBLIC_NAME = "EUA"
	SITE_URL = "WWW.IBELIEVE.COM"
	SEO_STATUS = 201
}
