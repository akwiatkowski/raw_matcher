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

func (sg ScriptGenerator) GenerateStats(matchedSize int, notFound int, alreadyExist int) {
  writer := bufio.NewWriter(sg.file)
  var err error

  line := fmt.Sprint("# STATS\n# matched: ", matchedSize, "\n# not found: ", notFound, "\n# already exists: ", alreadyExist ,"\n \n\n")
  _, err = writer.WriteString(line)
  logError(err)

  writer.Flush()
}

func (sg ScriptGenerator) GenerateForMatched(photo_raws [][2]PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  var totalSize int64
  totalSize = 0

  totalLength := len(photo_raws)

  for i, photo_raw := range photo_raws {
    _ = i

    photo := photo_raw[0]
    raw := photo_raw[1]

    line := fmt.Sprint("# ", i + 1, "/", totalLength, " matched \n")
    _, err = writer.WriteString(line)
    logError(err)

    line = fmt.Sprint("mkdir -p \"", photo.DirRawPath(), "\" \n")
    _, err = writer.WriteString(line)
    logError(err)

    line = fmt.Sprint("cp -n \"", raw.Path, "\" \"", photo.DirRawPath(), "\" \n")
    _, err = writer.WriteString(line)
    logError(err)

    totalSize += raw.Size

    if (i + 1) % 200 == 0 {
      line = fmt.Sprint("echo \"copied ", i + 1, " files, ", totalSize / 1048576, "MB \" \n")
      _, err = writer.WriteString(line)
      logError(err)
    }
  }

  writer.Flush()
}

func (sg ScriptGenerator) GenerateForNotFound(photos []PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  totalLength := len(photos)

  for i, photo := range photos {
    _ = i

    line := fmt.Sprint("# ", i, "/", totalLength, " - ", photo.Path, " NOT found RAW file\n")
    _, err = writer.WriteString(line)
    logError(err)
  }

  writer.Flush()
}

func (sg ScriptGenerator) GenerateForRawExist(photos []PhotoFile) {
  writer := bufio.NewWriter(sg.file)
  var err error

  totalLength := len(photos)

  for i, photo := range photos {
    _ = i

    line := fmt.Sprint("# ", i, "/", totalLength, " - ", photo.Path, " RAW already exists\n")
    _, err = writer.WriteString(line)
    logError(err)
  }

  writer.Flush()
}
