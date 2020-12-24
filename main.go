package main

import (
  "log"
  "fmt"

  "raw_matcher/matcher"
)

func main() {
  log.Print("Start")

  params := &matcher.MatcherParams{
    PhotosPath: "data/photos",
    RawsPath: "data/raws" }

  instance := matcher.New(params)
  log.Print(fmt.Sprint("instance ", instance))
  // fmt.Println(instance.FileList.Photos)
}
