# stargazer

Stargazing using Github API.

```go
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
```
```
name                                            stars   forks   watch
hoanhan101/ultimate-go                          6926    451     207
hoanhan101/algo                                 77      14      6
hoanhan101/request                              0       0       1
hoanhan101/dotfiles                             2       0       0
hoanhan101/lang                                 0       0       2
vapor-ware/kubetest                             81      25      14
vapor-ware/synse-server                         21      7       16
Bennington-Distributed-Systems-2017/DarkDarkGo  279     36      18
bennington-hardware-hacking-2019/pos_system     27      10      5
```
