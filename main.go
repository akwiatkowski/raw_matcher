package main

import (
  "log"
  "fmt"

  "raw_matcher/matcher"
)

func main() {
  log.Print("Start")

  params := &matcher.MatcherParams{
    PhotosPath: "/home/olek/projects/koty_polskie/input",
    RawsPath: "/home/olek/projects/koty_polskie/input",
    OutputScriptName: "script.sh" }

  instance := matcher.New(params)
  instance.Match()

  fmt.Println("done")
  // fmt.Println(instance.FileList.Photos)
}
