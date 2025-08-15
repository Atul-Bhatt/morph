package main

import (
	"fmt"
	"os"
	"path/filepath"
	"image/jpeg"
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

		f, err := os.Create(filepath.Join(fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}
		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}
		f.Close()		

		client.SetImage(fmt.Sprintf("test%03d.jpg", n))
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
