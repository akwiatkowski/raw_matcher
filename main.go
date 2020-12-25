package main

import (
  "log"
  "fmt"

  "raw_matcher/matcher"
)

func main() {
  log.Print("Start")

  photosPath := "data"
  rawsPath := "data"

  // photosPath = "/home/olek/projects/koty_polskie/input"
  // rawsPath = "/media/silver"

  params := &matcher.MatcherParams{
    PhotosPath: photosPath,
    RawsPath: rawsPath,
    OutputScriptName: "script.sh" }

  instance := matcher.New(params)
  instance.Match()

  fmt.Println("done")
  // fmt.Println(instance.FileList.Photos)
}
