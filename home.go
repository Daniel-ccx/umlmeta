package main

import (
	"fmt"
	"models"
)

func main() {
    var cfg = models.ParseCfg("config");
    fmt.Println(cfg)

    fmt.Println("adf")
}
