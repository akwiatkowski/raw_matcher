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
  instance.Match()
  instance.GenerateCopyScript()

  fmt.Println("done")
  // fmt.Println(instance.FileList.Photos)
}
