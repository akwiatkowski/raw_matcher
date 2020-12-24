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
      log.Print(fmt.Sprint(" not found for ", photo.DateName()))
    } else {
      // assign
      // photo.AssignedRaw = result
      log.Print(fmt.Sprint(" found ", photo.DateName(), " -> ", result.DateName()))

      //log.Print(fmt.Sprint("cp ", result.Path, " ", photo.DirRawPath() ))
    }
  }

  log.Print("End matching")
}
