package matcher

import (
  "fmt"
  "log"
)

type Matcher struct {
  params *MatcherParams

  FileList *FileList
  ScriptGenerator *ScriptGenerator
}

func New(params *MatcherParams) *Matcher {
  return &Matcher {
    FileList: NewFileList(params),
    ScriptGenerator: NewScriptGenerator(params),
    params: params }
}

func (matcher Matcher) Match() {
  log.Print("Start matching")

  for i, photo := range matcher.FileList.Photos {
    _ = i

    // check if RAW exists
    rawExist, _ := photo.RawFileExists()

    if rawExist {
      // add photo to `ExistRawPhotos`
      matcher.ScriptGenerator.addToRawExist(photo)

      log.Print(fmt.Sprint(photo.DateName(), " has raw"))
    } else {
      // search for best RAW
      matchedRaw := matcher.FileList.MatchRawForPhoto(photo)

      if matchedRaw == nil {
        // RAW not found
        matcher.ScriptGenerator.addToNotFound(photo)

        log.Print(fmt.Sprint(photo.DateName(), " NOT found"))
      } else {
        // RAW not found
        matcher.ScriptGenerator.addToMatched(photo, *matchedRaw)

        log.Print(fmt.Sprint(photo.DateName(), " found RAW: ", matchedRaw.DateName()))
      }
    }
  }

  log.Print("End matching")

  matcher.ScriptGenerator.Run()
}
