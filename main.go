package main

import (
	"fmt"
	"image"
	"image/png"
	"bytes"
	"strings"
	"time"
	"sync"
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
	"github.com/otiai10/gosseract/v2"
	"github.com/gen2brain/go-fitz"
	"github.com/go-pdf/fpdf"
)

func main() {
	now := time.Now()
	uiText := Text(Font("helvetica", 10), Padx("2m"), Pady("2m"))
	uiText.InsertML("Metamorphosis<br>")
	uiTextChan := make(chan string, 5)

	doc, err := fitz.New("Metamorphosis.pdf")
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	ActivateTheme("azure light")
	out := Label(Height(2), Anchor("e"), Txt("Morph PDF Editor"))
	Grid(out, Columnspan(1), Sticky("e"))

	var wg sync.WaitGroup
	// Extract pages as images and pass to tesseract
	for n := 0; n < doc.NumPage(); n++ {
		wg.Add(1)
		img, err := doc.Image(n)
		if err != nil {
			fmt.Println("Error processing image")
		}
		go ImageToText(img, n, &wg, uiTextChan)
	}

	var poll func()
	poll = func() {
		select {
		case t := <-uiTextChan:
			uiText.InsertML(t + "<br><br>")
		default:
			TclAfter(100, poll)
		}
	}
	TclAfter(100, poll)

	Grid(uiText, Padx("2m"), Pady("2m"))
	Grid(TButton(Txt("Save PDF"), Command(func() { SavePDF(uiText.Get("1.0", "end-1c")) })))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
	wg.Wait()
	fmt.Println("Time taken: ", now.Sub(time.Now()))
}

func ImageToText(img *image.RGBA, n int, wg *sync.WaitGroup, uiTextChan chan string) {
	client := gosseract.NewClient()
	defer client.Close()
	defer wg.Done()

	fmt.Println("Processing page: ", n)

	var buf bytes.Buffer
	png.Encode(&buf, img)
	b := buf.Bytes()

	client.SetImageFromBytes(b)
	text, err := client.Text()
	if err != nil {
		fmt.Println(err)
	}
	uiTextChan <- text
	fmt.Println("Processing complete: ", n)
}

func SavePDF(text []string) {
	pdf := fpdf.New(fpdf.OrientationPortrait, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("helvetica", "", 12)
	pdf.MultiCell(0, 5, strings.Join(text, " "), "", "", false)
	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Println("Error creating pdf: ", err)
	}
}
