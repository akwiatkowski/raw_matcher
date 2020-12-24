package matcher

import (
  "log"
  "fmt"
  "bufio"
  "os"
)

type ScriptGenerator struct {
  params *MatcherParams

  MatchedPhotos [][2]PhotoFile
  NotFoundPhotos []PhotoFile
  RawExistPhotos []PhotoFile
}

func NewScriptGenerator(params *MatcherParams) *ScriptGenerator {
  log.Print("ScriptGenerator start")

  return &ScriptGenerator {
    params: params }
}

func (sg ScriptGenerator) addToRawExist(photo PhotoFile) {
  sg.RawExistPhotos = append(sg.RawExistPhotos, photo)

  log.Print(fmt.Sprint("addToRawExist ", photo.DateName(), " len=", len(sg.RawExistPhotos)))
}

func (sg ScriptGenerator) addToNotFound(photo PhotoFile) {
  sg.NotFoundPhotos = append(sg.NotFoundPhotos, photo)

  log.Print(fmt.Sprint("addToNotFound ", photo.DateName(), " len=", len(sg.NotFoundPhotos)))
}

func (sg ScriptGenerator) addToMatched(photo PhotoFile, raw PhotoFile) {
  var row [2]PhotoFile
  row[0] = photo
  row[1] = raw
  sg.MatchedPhotos = append(sg.MatchedPhotos, row)

  log.Print(fmt.Sprint("addToMatched ", photo.DateName(), " len=", len(sg.MatchedPhotos)))
}

func logError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func (sg ScriptGenerator) Run() {
  log.Print("Start script")

  file, err := os.Create("script.sh")
  logError(err)

  defer file.Close()

  writer := bufio.NewWriter(file)

  log.Print(fmt.Sprint("Script Matched len=", len(sg.MatchedPhotos )))
  for i, photo_raw := range sg.MatchedPhotos {
    _ = i

    photo := photo_raw[0]
    raw := photo_raw[1]

    _, err = writer.WriteString(fmt.Sprint("cp \"", raw.Path, "\" \"", photo.DirRawPath(), "\" \n"))
    logError(err)
  }

  log.Print(fmt.Sprint("Script exists len=", len(sg.RawExistPhotos )))
  for i, photo := range sg.RawExistPhotos {
    _ = i

    _, err := writer.WriteString(fmt.Sprint("# ", photo.DateName(), " RAW file exists\n"))
    logError(err)
  }

  log.Print(fmt.Sprint("Script not found len=", len(sg.NotFoundPhotos )))
  for i, photo := range sg.NotFoundPhotos {
    _ = i

    _, err := writer.WriteString(fmt.Sprint("# ", photo.DateName(), " NOT found RAW file\n"))
    logError(err)
  }

  writer.Flush()

  log.Print("End script")
}
