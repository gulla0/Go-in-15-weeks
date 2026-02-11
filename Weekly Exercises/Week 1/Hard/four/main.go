package main

import (
	"Hard/four/utils"
	"fmt"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.twitter.com", "https://www.instagram.com"}
	results := utils.CheckURLs(urls)
	fmt.Println(results)
}
