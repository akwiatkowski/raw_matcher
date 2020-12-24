package main

import (
	"flag"
	"fmt"
	"os"

  "raw_matcher/matcher"
)

func main() {
	var path, output_script string
	flag.StringVar(&path, "path", ".", "Input path")
  flag.StringVar(&output_script, "output", "do.sh", "Output script")
  flag.Parse()

	if len(path) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify input path")
	}

	fmt.Println(fmt.Sprint("Input path: ", path))
  fmt.Println(fmt.Sprint("Output script: ", output_script))

  params := &matcher.MatcherParams{
    PhotosPath: path,
    RawsPath: path,
    OutputScriptName: output_script }

  instance := matcher.New(params)
  instance.Match()
}
