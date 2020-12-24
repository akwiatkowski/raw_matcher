package matcher

import (
  "fmt"
  "log"
)

type Matcher struct {
  params *MatcherParams

  FileList *FileList
}

func New(params *MatcherParams) *Matcher {
  return &Matcher {
    FileList: NewFileList(params),
    params: params }
}

func (matcher Matcher) Match() {
  log.Print("Start matching")

  for i, photo := range matcher.FileList.Photos {
    _ = i
    result := matcher.FileList.MatchRawForPhoto(photo)
    if result == nil {
      log.Print(fmt.Sprint(" not found for ", photo.dateName()))
    } else {
      log.Print(fmt.Sprint(" found ", photo.dateName(), " -> ", result.dateName()))
    }
  }

  log.Print("End matching")
}
