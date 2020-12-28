package raw_matcher

import (
  "log"
  "fmt"
  "bufio"
  "os"
)

type ScriptGenerator struct {
  params *MatcherParams
  file *os.File
}

func NewScriptGenerator(params *MatcherParams) *ScriptGenerator {
  log.Print("ScriptGenerator start")

  file, err := os.Create(params.OutputScriptName)
  logError(err)

  return &ScriptGenerator {
    params: params,
    file: file}
}

func logError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func (sg ScriptGenerator) Close() {
  sg.file.Close()
}

func (sg ScriptGenerator) GenerateForMatched(photo_raws [][2]PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  for i, photo_raw := range photo_raws {
    _ = i

    photo := photo_raw[0]
    raw := photo_raw[1]

    line := fmt.Sprint("# ", i+1, " matched \n")
    _, err = writer.WriteString(line)
    logError(err)

    line = fmt.Sprint("cp -n \"", raw.Path, "\" \"", photo.DirRawPath(), "\" \n")
    _, err = writer.WriteString(line)
    logError(err)
  }

  writer.Flush()
}

func (sg ScriptGenerator) GenerateForNotFound(photos []PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  for i, photo := range photos {
    _ = i

    line := fmt.Sprint("# ", photo.DateName(), " NOT found RAW file\n")
    _, err = writer.WriteString(line)
    logError(err)
  }

  writer.Flush()
}

func (sg ScriptGenerator) GenerateForRawExist(photos []PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  for i, photo := range photos {
    _ = i

    line := fmt.Sprint("# ", photo.DateName(), " already exists\n")
    _, err = writer.WriteString(line)
    logError(err)
  }

  writer.Flush()
}
