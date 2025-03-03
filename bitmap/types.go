package bitmap

import "image/color"

type Bitmap struct {
	Width, Height int
	Pixels        [][]color.NRGBA // https://stackoverflow.com/a/54309181
}
