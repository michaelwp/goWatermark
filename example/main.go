package main

import (
	"fmt"
	"image/color"

	"github.com/michaelwp/goWatermark"
)

func main() {
	err := gowatermark.AddWatermark(
		&gowatermark.Watermark{
			Image:      "input.jpg",
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
