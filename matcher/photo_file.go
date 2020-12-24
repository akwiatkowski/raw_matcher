package matcher

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "time"
  "strings"
)

type PhotoFile struct {
  // https://stackoverflow.com/questions/24216510/empty-or-not-required-struct-fields-in-golang

  DateString string
  Date time.Time
  RawFilenameString string

  Path string
  FileInfo os.FileInfo
}

func NewPhotoFile(path string, fileInfo os.FileInfo) PhotoFile {
  dateString := processDateFromPath(path)
  filenameString := processRawFilenameFromPath(path)
  date := processDate(dateString)

  return PhotoFile {
    Path: path,
    FileInfo: fileInfo,
    DateString: dateString,
    Date: date,
    RawFilenameString: filenameString }
}

// use date stored in path
// it's faster that using exif data
func processDateFromPath(path string) string {
  re := dateRegexp()
	matched := re.FindAllString(path, -1)
  if len(matched) > 0 {
    lastElement := matched[len(matched) - 1]
    log.Print(fmt.Sprint(" ", path, " -> date ", lastElement))
    return(lastElement)
  } else {
    return ""
  }

	// for _, element := range submatchall {
  //   fmt.Println(element)
	// }

  return ""
}

func processRawFilenameFromPath(path string) string {
  re := rawFilenameRegexp()
	matched := re.FindAllStringSubmatch(path, -1)

  if len(matched) > 0 {
    lastElement := matched[len(matched) - 1][1]
    log.Print(fmt.Sprint(" ", path, " -> filename ", lastElement))
    return(lastElement)
  } else {
    return ""
  }

  return ""
}

func processDate(dateString string) time.Time {
  layout := "2006_01_02"
  normalizedDateString := strings.ReplaceAll(dateString, "-", "_")

  t, err := time.Parse(layout, normalizedDateString)

  if err != nil {
    log.Fatal(err)
  }

  return t
}

func dateRegexp() *regexp.Regexp {
  return regexp.MustCompile(`\d{4}[-_]\d{2}[-_]\d{2}`)
}

func rawFilenameRegexp() *regexp.Regexp {
  return regexp.MustCompile(`_?([^_./]+)\.\w{3,4}`)
}
