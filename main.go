package main

import (
  "fmt"

  "raw_matcher/matcher"
)

func main() {
  fmt.Println("test")

  instance := matcher.New()
  fmt.Println(instance)
  fmt.Println(instance.FileList.Photos)
}
