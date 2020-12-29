package raw_matcher

import (
  "fmt"
  "log"
)

type Matcher struct {
  params *MatcherParams

  FileList *FileList
  ScriptGenerator *ScriptGenerator

  MatchedPhotos [][2]PhotoFile
  NotFoundPhotos []PhotoFile
  RawExistPhotos []PhotoFile
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
      matcher.RawExistPhotos = append(matcher.RawExistPhotos, photo)

      log.Print(fmt.Sprint(photo.DateName(), " has raw"))
    } else {
      // search for best RAW
      matchedRaw := matcher.FileList.MatchRawForPhoto(photo)

      if matchedRaw == nil {
        // RAW not found
        matcher.NotFoundPhotos = append(matcher.NotFoundPhotos, photo)

        log.Print(fmt.Sprint(photo.DateName(), " NOT found"))
      } else {
        // raw Found
        var row [2]PhotoFile
        row[0] = photo
        row[1] = *matchedRaw
        matcher.MatchedPhotos = append(matcher.MatchedPhotos, row)

        log.Print(fmt.Sprint(photo.DateName(), " found RAW: ", matchedRaw.DateName()))
      }
    }
  }

  log.Print("End matching")

  matcher.ScriptGenerator.GenerateStats(
    len(matcher.MatchedPhotos),
    len(matcher.NotFoundPhotos),
    len(matcher.RawExistPhotos) )

  matcher.ScriptGenerator.GenerateForMatched(matcher.MatchedPhotos)
  matcher.ScriptGenerator.GenerateForNotFound(matcher.NotFoundPhotos)
  matcher.ScriptGenerator.GenerateForRawExist(matcher.RawExistPhotos)

  matcher.ScriptGenerator.Close()
}
