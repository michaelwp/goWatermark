# goWatermark
This package allows you to add customizable watermarks to images using the Go programming language.
It provides functionalities for positioning, repeating text, and adjusting font properties.

### Installation

```sh
go get github.com/michaelwp/goWatermark
```

### Example

```go
package main

import (
	"fmt"
	"image/color"

	"github.com/michaelwp/goWatermark"
)

func main() {
	err := gowatermark.AddWatermark(
		&gowatermark.Watermark{
			Image:      "input1.jpeg",
			OutputFile: "output.jpeg",
			Text:       "79995782-PTGLOBALPRADANASEJAHTERA-227",
			Position: gowatermark.Position{
				PosAY: 10,
			},
			Font: gowatermark.Font{
				FontSize: 12,
			},
			Color: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 80,
			},
			Align: gowatermark.AlignCenter,
			Repeat: gowatermark.Repeat{
				RepY: 20,
				RepX: 10,
			},
			LineSpacing: 25,
			Rotate:      -30,
			ImgSize: gowatermark.ImgSize{
				Width: 250,
			},
		},
	)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Watermark added successfully!")
	}
}
```

This example demonstrates how to configure and apply a watermark to an image using the `go_watermark` package. Adjust the parameters as needed to fit your specific use case.

for detail explanation visit: [How to Add a Watermark onto an Image Using Go](https://www.goblog.dev/articles/33)
