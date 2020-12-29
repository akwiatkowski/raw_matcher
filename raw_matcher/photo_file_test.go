package raw_matcher

import (
  "testing"
)

/*
  TODO:
    Equal
    processDateFromPath
    processDate
*/

func TestFilterPhotoFilename(t *testing.T) {
    testTable := []struct {
      fileName string
      output string
    } {
      { "DSC05439.jpg", "DSC05439" },
      { "DSC05439raw1.jpg", "DSC05439" },
      { "DSC05439-hdr.jpg", "DSC05439" },
      { "DSC05439b.jpg", "DSC05439" },
      { "2020_11_03__10_16_DSC05439.jpg", "DSC05439" },
      { "2020_11_03__10_16_DSC05439_01.jpg", "DSC05439" },
      { "2020_11_03__10_16_DSC05439_01a.jpg", "DSC05439" },
      { "2020_12_06__14_42_DSC09064-denoise-denoise.jpg", "DSC09064" },
    }

    for _, row := range testTable {
		  result := FilterPhotoFilename(row.fileName)

      if result != row.output {
			t.Errorf("FilterPhotoFilename was incorrect, from %s, got %s, want: %s", row.fileName, result, row.output)
		}
	}
}
