package matcher

import (
  "fmt"
  "regexp"
  "log"
  "path/filepath"
  "os"
)

type FileList struct {
  params *MatcherParams

  Photos []PhotoFile
  Raws []PhotoFile
}

func (fl FileList) MatchRawForPhoto(photo PhotoFile) *PhotoFile {
  for i, raw := range fl.Raws {
    _ = i
    if photo.equal(raw) {
      return &raw
    }
  }

  return nil
}

func jpegRegexp() *regexp.Regexp {
  r, e := regexp.Compile("(?i)^.+\\.(jpg|jpeg)$")
  if e != nil {
      log.Fatal(e)
  }
  return r
}

func rawRegexp() *regexp.Regexp {
  // https://stackoverflow.com/questions/15326421/how-do-i-do-a-case-insensitive-regular-expression-in-go
  r, e := regexp.Compile("(?i)^.+\\.(dng|pef|arw)$")
  if e != nil {
      log.Fatal(e)
  }
  return r
}

func scanPhotoFiles(rxp *regexp.Regexp, path string) []PhotoFile {
  var photoFiles []PhotoFile

  e := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    if err == nil && rxp.MatchString(info.Name()) {
      // fmt.Println(info.Name())

      photoFiles = append(photoFiles, NewPhotoFile(path, info))
    }
    return nil
  })
  if e != nil {
      log.Fatal(e)
  }

  return photoFiles
}

func scanPhotos(photosPath string) []PhotoFile {
  log.Print("scanPhotos")
  result := scanPhotoFiles(jpegRegexp(), photosPath)
  log.Print(fmt.Sprint("done with ", cap(result)))
  return result
}

func scanRaws(rawsPath string) []PhotoFile {
  log.Print("scanRaws")
  result := scanPhotoFiles(rawRegexp(), rawsPath)
  log.Print(fmt.Sprint("done with ", cap(result)))
  return result
}

func NewFileList(params *MatcherParams) *FileList {
  log.Print("FileList start")

  return &FileList {
    params: params,
    Photos: scanPhotos(params.PhotosPath),
    Raws: scanRaws(params.RawsPath) }
}
