package matcher

import (
  "testing"
  "raw_matcher/matcher"
)

func TestFilterPhotoFilenameA(t *testing.T) {
    testTable := []struct {
      fileName string
      output string
    } {
      { "DSC05439.jpg", "DSC05439" },
      { "DSC05439raw1.jpg", "DSC05439" },
      { "DSC05439-hdr.jpg", "DSC05439" },
      { "DSC05439b.jpg", "DSC05439" },
      { "2020_11_03__10_16_DSC05439.jpg", "DSC05439" },
    }

    for _, row := range testTable {
		  result := matcher.FilterPhotoFilename(row.fileName)

      if result != row.output {
			t.Errorf("FilterPhotoFilename was incorrect, from %s, got %s, want: %s", row.fileName, result, row.output)
		}
	}
}
