package main

import "savior_oss/main/router"

func main() {
	r := router.SetUpRoute()
	_ = r.Run(":8080")
}
