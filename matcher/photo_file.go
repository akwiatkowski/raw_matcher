package matcher

import (
  _ "fmt"
  "os"
  "regexp"
)

type PhotoFile struct {
  // https://stackoverflow.com/questions/24216510/empty-or-not-required-struct-fields-in-golang
  DateString string

  Path string
  FileInfo os.FileInfo
}

func NewPhotoFile(path string, fileInfo os.FileInfo) PhotoFile {
  return PhotoFile {
    Path: path,
    FileInfo: fileInfo,
    DateString: processDateFromPath(path) }
}

// date should be somewhere in path
func processDateFromPath(path string) string {
  re := dateRegexp()
	matched := re.FindAllString(path, -1)
  if len(matched) > 0 {
    lastElement := matched[len(matched) - 1]
    return(lastElement)
  } else {
    return ""
  }

	// for _, element := range submatchall {
  //   fmt.Println(element)
	// }

  return ""
}

func dateRegexp() *regexp.Regexp {
  return regexp.MustCompile(`\d{4}[-_]\d{2}[-_]\d{2}`)
}
