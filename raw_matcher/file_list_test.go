package raw_matcher

import (
  "testing"
  "os"
)

func TestMatchRawForPhoto(t *testing.T) {
  fileInfo, _ := os.Stat("README.md")

  rawFile := NewPhotoFile(
    "2020-11-11-test/IMG0001.arw",
    fileInfo )

  photoFile := NewPhotoFile(
    "2020-11-11-test/IMG0001_raw1.jpg",
    fileInfo )

  fl := FileList {
    Params: &MatcherParams {},
    Raws: []PhotoFile{ rawFile },
    Photos: []PhotoFile{ photoFile } }

  result := fl.MatchRawForPhoto(photoFile)

  if result == nil {
    t.Errorf("MatchRawForPhoto %s should not be nil", result)
  }
}

func TestIsNotBlacklisted(t *testing.T) {
    testTable := []struct {
      fileName string
      output bool
    } {
      { "2019-12-12/@eaDir/IMG1.jpg", false },
      { "2019-12-12/DSC05439raw1.jpg", true },
      { "DSC05439-hdr.jpg", true },
      { "SYNOPHOTO_THUMB/2020_11_03__10_16_DSC05439.jpg", false },
    }

    for _, row := range testTable {
		  result := IsNotBlacklisted(row.fileName)

      if result != row.output {
			t.Errorf("IsNotBlacklisted was incorrect, from %s, got %t, want: %t", row.fileName, result, row.output)
		}
	}
}
