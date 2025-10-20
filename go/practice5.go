package main

import (
	"fmt"
)

var appConfig = Config{Env: "test"}

type Config struct {
	Env string
}

// Cの宣言と同じで、返り値にポインタ型であるという宣言をしている
func getConfig() *Config {
	return &appConfig
}

func main() {
	c := getConfig()
	c.Env = "production"

	fmt.Println(c.Env)         // production
	fmt.Println(appConfig.Env) // testではなくproducionへ
}
