package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"../github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	fmt.Printf("%d issues:\n", result.TotalCount)

	fmt.Println("\n1ヶ月未満")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days < 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("\n1年未満")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days < 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("\n1年以上")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days >= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
