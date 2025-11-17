package main

import (
	"fmt"

	"github.com/Vinayakatk/go-learn/config"
	"github.com/vrischmann/envconfig"
)

func main() {
	//  := models.User{
	// 	Id: 1,
	// }
	c := config.Config{}
	if err := envconfig.Init(&c); err != nil {
		fmt.Printf("err=%s\n", err)
	}
	println(c.DBHost)
}