package stargazer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hoanhan101/request"
)

type Repo struct {
	Name   string `json:"full_name"`
	Stars  int    `json:"stargazers_count"`
	Forks  int    `json:"forks"`
	Watch  int    `json:"subscribers_count"`
	Issues int    `json:"open_issues_count"`
}

func Get(repos []string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", "name", "stars", "forks", "watch", "issues")

	for _, repo := range repos {
		r := getRepo(repo)
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", r.Name, r.Stars, r.Forks, r.Watch, r.Issues)
	}

	w.Flush()
}

func getRepo(name string) *Repo {
	r := new(Repo)

	err := request.GetJSON(
		&request.Options{
			URL: fmt.Sprintf("https://api.github.com/repos/%s", name),
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	return r
}
