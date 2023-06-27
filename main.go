package main

import (
	"fmt"

	"github.com/OSX-RSPlug-a/linuxtips-golang-week-d01/git"
)

func main() {
	fmt.Println("################# #VAAAAAAAAI #################")
	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	b := git.BuscadorGit{
		DonoDoRepositorio: donoDoRepositorio,
		Repo:              repo,
	}
	b.BuscaGitTag(tag)

	fmt.Println("################# Halting the program... #################")
}
