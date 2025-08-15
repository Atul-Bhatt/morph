package main

import (
	"fmt"
	"image/png"
	"bytes"
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
	"github.com/otiai10/gosseract/v2"
	"github.com/gen2brain/go-fitz"
)

func main() {
	var text string

	client := gosseract.NewClient()
	defer client.Close()

	doc, err := fitz.New("sample_text.pdf")
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	// Extract pages as images and pass to tesseract
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		png.Encode(&buf, img)
		b := buf.Bytes()

		client.SetImageFromBytes(b)
		text, err = client.Text()
		if err != nil {
			fmt.Println(err)
		}
	}

	// pass the text to UI
	t := Text(Font("helvetica", 10), Padx("2m"), Pady("2m"))
	t.InsertML(text)
	ActivateTheme("azure light")
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
}
