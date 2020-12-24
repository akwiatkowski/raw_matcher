package matcher

import (
  "log"
  "fmt"
  "bufio"
  "os"
)

type ScriptGenerator struct {
    MatchedPhotos [][]PhotoFile
}

func (matcher Matcher) GenerateCopyScript() {
  log.Print("Start script")

  file, err := os.Create("script.sh")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  writer := bufio.NewWriter(file)

  for i, photo := range matcher.FileList.Photos {
    _ = i
    _ = photo

    // check if there is no raw file already
    rawExists, _ := photo.RawFileExists()

    if rawExists {
      _, err := writer.WriteString(fmt.Sprint("# ", photo.DateName(), " RAW file exists\n"))
      if err != nil {
        log.Fatal(err)
      }

    } else {
      if photo.AssignedRaw != nil {
        _, err := writer.WriteString(fmt.Sprint("# ", photo.DateName(), " copy RAW file\n"))
        if err != nil {
          log.Fatal(err)
        }

        _, err = writer.WriteString(fmt.Sprint("cp \"", photo.AssignedRaw.Path, "\" \"", photo.DirRawPath(), "\" \n"))
        if err != nil {
          log.Fatal(err)
        }

      } else {
        _, err := writer.WriteString(fmt.Sprint("# ", photo.DateName(), " NOT found RAW file\n"))
        if err != nil {
          log.Fatal(err)
        }
      }
    }
  }

  writer.Flush()

  log.Print("End script")
}
