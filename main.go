package main

import (
  "fmt"

  "raw_matcher/matcher"
)

func main() {
  fmt.Println("test")

  params := &matcher.MatcherParams{
    PhotosPath: "data/photos",
    RawsPath: "data/raws" }

  instance := matcher.New(params)
  fmt.Println(instance)
  fmt.Println(instance.FileList.Photos)
}
