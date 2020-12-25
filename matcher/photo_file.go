package matcher

import (
  "fmt"
  "log"
  "os"
  "regexp"
  "time"
  "strings"
  "math"
  "path"
  "path/filepath"
)

const MaxHourOffset = 96.0
const RawSubdir = "raw"

type PhotoFile struct {
  // https://stackoverflow.com/questions/24216510/empty-or-not-required-struct-fields-in-golang

  DateString string
  Date time.Time
  // not filterd
  RealFilename string
  // filtered
  Filename string

  Path string
  FileInfo os.FileInfo

  // maybe if I convert into pointer/reference that could work`
  // but I won't do that now
  // AssignedRaw *PhotoFile
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
    Filename: filenameString }
}

func (pf PhotoFile) DirPath() string {
  return path.Dir(pf.Path)
}

func (pf PhotoFile) DirRawPath() string {
  return path.Join(pf.DirPath(), RawSubdir) + "/"
}

// w/o extension
func (pf PhotoFile) FilenameRawPath() string {
  return path.Join(pf.DirRawPath(), pf.Filename)
}

func (pf PhotoFile) RawFileExists() (bool, error) {
  matches, err := filepath.Glob(pf.FilenameRawPath() + ".*")
  if err != nil {
    return false, err
  }
  return len(matches) > 0, nil
}

func (pf PhotoFile) DateName() string {
  return fmt.Sprint(pf.DateString, ":", pf.Filename)
}

func (pf PhotoFile) Equal(other PhotoFile) bool {
  // check substring because darktable added `_01` intead of overwrite
  // and I quite often added `raw1` or `b` after
  if strings.Contains(pf.Filename, other.Filename) == false {
    return false
  }

  timeDiff := pf.Date.Sub(other.Date)
  result := math.Abs(timeDiff.Hours()) < MaxHourOffset

  log.Print(
    fmt.Sprint(
      "  ",
      pf.DateName(),
      " - ",
      other.DateName(),
      " timeDiff ",
      timeDiff,
      " result ",
      result ))

  return result
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

func FilterPhotoFilename(basePath string) string {
  // https://regex101.com/r/w4gLE3/1

  // not sure if/how I can get strin from regexp
  // `^\d{4}\D+\d{2}\D+\d{2}\D+\d{2}\D+\d{2}\D([^.]+)\..+`

  // log.Print(fmt.Sprint("filter0 ", basePath))

  reRemoveExtension := regexp.MustCompile(`\.[^.]+`)
  s := reRemoveExtension.ReplaceAllString(basePath, `$1`)

  // log.Print(fmt.Sprint("filter1 ", s))

  reRemoveTimedate := regexp.MustCompile(`^\d{4}\D+\d{2}\D+\d{2}\D+\d{2}\D+\d{2}\D`)
  s = reRemoveTimedate.ReplaceAllString(s, `$1`)

  // log.Print(fmt.Sprint("filter2 ", s))

  reRemoveSuffix := regexp.MustCompile(`^(\D*\d+)`)
  s = reRemoveSuffix.FindString(s)

  // log.Print(fmt.Sprint("filter3 ", s))

  return s
}

func processRawFilenameFromPath(fullPath string) string {
  re := rawFilenameRegexp()
  basePath := path.Base(fullPath)

  // filter not needed suffixes
  basePath = FilterPhotoFilename(basePath)

	matched := re.FindAllStringSubmatch(basePath, -1)

  if len(matched) > 0 {
    lastElement := matched[len(matched) - 1][1]
    log.Print(fmt.Sprint(" ", fullPath, " -> filename ", lastElement))
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
