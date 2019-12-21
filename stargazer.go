package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hoanhan101/request"
)

type Repo struct {
	Name  string `json:"full_name"`
	Stars int    `json:"stargazers_count"`
	Forks int    `json:"forks"`
	Watch int    `json:"subscribers_count"`
}

func GetRepo(name string) *Repo {
	r := new(Repo)

	err := request.GetJSON(fmt.Sprintf("https://api.github.com/repos/%s", name), r)
	if err != nil {
		fmt.Println(err)
	}

	return r
}

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

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

	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "name", "stars", "forks", "watch")

	for _, repo := range repos {
		r := GetRepo(repo)
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", r.Name, r.Stars, r.Forks, r.Watch)
	}

	w.Flush()
}
