package main

import (
	"github.com/hoanhan101/stargazer"
)

func main() {
	repos := []string{
		"hoanhan101/ultimate-go",
		"hoanhan101/algo",
		"hoanhan101/request",
		"hoanhan101/dotfiles",
		"hoanhan101/lang",
		"vapor-ware/kubetest",
		"vapor-ware/synse-server",
		"Bennington-Distributed-Systems-2017/DarkDarkGo",
		"bennington-hardware-hacking-2019/pos_system",
	}

	stargazer.Get(repos)
}
