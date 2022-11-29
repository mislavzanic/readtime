# readtime
Go library for calculating the read time of a blog post.

This repo is a (not great, but will be) Go implementation of https://github.com/karthik512/estimated_read_time.

# Installation

``` go
go get github.com/mislavzanic/readtime
```


# Basic usage

``` go
package main

import (
    "fmt"
 
    "github.com/mislavzanic/readtime"
)

func main()  {
	myArticle := "some text ... ... ..."
    option, _ := readtime.NewOption()
        .IsTechnical(true)
        .TechnicalDifficulty(1)
    readTime := readTime.CalcReadTime(myArticle, *(option)).Seconds / 60
	fmt.Println(readTime) // read time in minutes
}
```
