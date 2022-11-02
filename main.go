package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/timrxd/mdlist-gen/gen"
)

func main() {

	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(-1)
	}

	gen.Login()

	// err = gen.FetchManga()
	// if err != nil {
	// 	fmt.Printf("\nerror: %s\n\n", err)
	// }

}
