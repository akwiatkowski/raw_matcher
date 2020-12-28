package main

import (
	"flag"
	"fmt"
	"os"

  "github.com/akwiatkowski/raw_matcher"
)

func main() {
	var photosPath, rawsPath, outputPath string
	flag.StringVar(&photosPath, "photos_path", "./", "Photos path")
  flag.StringVar(&rawsPath, "raws_path", "./", "Raws path")
  flag.StringVar(&outputPath, "output", "do.sh", "Output script")
  flag.Parse()

	if len(photosPath) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify photos path")
	}

  if len(rawsPath) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify raws path")
	}

  fmt.Println(fmt.Sprint("Photos path: ", photosPath))
	fmt.Println(fmt.Sprint("Raws path: ", rawsPath))
  fmt.Println(fmt.Sprint("Output script: ", outputPath))

  params := &raw_matcher.MatcherParams{
    PhotosPath: photosPath,
    RawsPath: rawsPath,
    OutputScriptName: outputPath }

  instance := raw_matcher.New(params)
  instance.Match()
}
