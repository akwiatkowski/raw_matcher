package matcher

import (
  "fmt"
  "regexp"
  "log"
  "path/filepath"
  "os"
)

type PhotoFile struct {
  Path string
  FileInfo os.FileInfo
}

type FileList struct {
  Photos []PhotoFile
  Raws []PhotoFile
}


func jpegRegexp() *regexp.Regexp {
  r, e := regexp.Compile("^.+\\.(jpg|jpeg)$")
  if e != nil {
      log.Fatal(e)
  }
  return r
}

func rawRegexp() *regexp.Regexp {
  r, e := regexp.Compile("^.+\\.(dng|pef|arw)$")
  if e != nil {
      log.Fatal(e)
  }
  return r
}

func scanPhotoFiles(rxp *regexp.Regexp, path string) []PhotoFile {
  var photoFiles []PhotoFile

  e := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    if err == nil && rxp.MatchString(info.Name()) {
      fmt.Println(info.Name())
      // fmt.Println(info)

      var photo PhotoFile
      photo.Path = path
      photo.FileInfo = info

      photoFiles = append(photoFiles, photo)
    }
    // fmt.Println(path)
    return nil
  })
  if e != nil {
      log.Fatal(e)
  }

  fmt.Println(photoFiles)

  return photoFiles
}

func scanPhotos() {

}


func NewFileList() *FileList {
  fmt.Println("File list START")

  var photos = scanPhotoFiles(jpegRegexp(), "data")

  return &FileList{Photos: photos}
}
