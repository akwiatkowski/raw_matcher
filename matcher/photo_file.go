package matcher

import (
  "os"
)

type PhotoFile struct {
  Path string
  FileInfo os.FileInfo
}

func NewPhotoFile(path string, fileInfo os.FileInfo) PhotoFile {
  return PhotoFile {
    Path: path,
    FileInfo: fileInfo }
}
