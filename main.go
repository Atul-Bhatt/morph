package main

import (
	"fmt"
	"image"
	"image/png"
	"bytes"
	"strings"
	"time"
	tk "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
	"github.com/otiai10/gosseract/v2"
	"github.com/gen2brain/go-fitz"
	"github.com/go-pdf/fpdf"
)

const APPNAME = "morph"

func main() {
	now := time.Now()
	app := NewApp()
	app.Run()

	doc, err := fitz.New("Metamorphosis.pdf")
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	processPages := func() {
		// Extract pages as images and pass to tesseract
		for n := 0; n < doc.NumPage(); n++ {
			img, err := doc.Image(n)
			if err != nil {
				fmt.Println("Error processing image")
			}
			_ = ImageToText(img, n)
			//uiText.InsertML(text + "<br><br>")
		}
	}
	processPages()

	//TclAfter(time.Second * 1, processPages)
	fmt.Println("Time taken: ", now.Sub(time.Now()))
}

func NewApp() *App {
	app := &App{name: "morph"}
	tk.StyleThemeUse("clam")
	tk.WmWithdraw(tk.App)
	tk.WmAttributes(tk.App, tk.Topmost(true))
	tk.App.WmTitle(APPNAME)
	tk.App.Configure(tk.Background(tk.LightYellow), tk.Pady(0), tk.Padx(0))
	tk.WmProtocol(tk.App, tk.WM_DELETE_WINDOW, app.onQuit)
	for _, key := range []string{"<Escape>", "<q>", "<Return>"} {
		tk.Bind(tk.App, key, tk.Command(app.onQuit))
	}
	tk.StyleConfigure("TButton", tk.Font(tk.HELVETICA, 36, tk.BOLD),
		tk.Background(tk.LightYellow), tk.Foreground(tk.Red))
	/*app.button = tk.TButton(tk.Txt(app.getMesage()), tk.Command(app.onQuit),
		tk.Justify(tk.CENTER))
	app.update()
	tk.Pack(app.button, tk.Fill(tk.FILL_BOTH), tk.Expand(true),
		tk.Ipadx(15), tk.Ipady(15))*/
	uiText := tk.Text(tk.Font("helvetica", 10), tk.Padx("2m"), tk.Pady("2m"))
	uiText.InsertML("Metamorphosis<br>")
	out := tk.Label(tk.Height(2), tk.Anchor("e"), tk.Txt("Morph PDF Editor"))
	tk.Grid(out, tk.Columnspan(1), tk.Sticky("e"))
	tk.Grid(uiText, tk.Padx("2m"), tk.Pady("2m"))
	tk.Grid(tk.TButton(tk.Txt("Save PDF"), tk.Command(func() { SavePDF(uiText.Get("1.0", "end-1c")) })))
	tk.Grid(tk.TExit(), tk.Padx("1m"), tk.Pady("2m"), tk.Ipadx("1m"), tk.Ipady("1m"))
	return app
}

func ImageToText(img *image.RGBA, n int) string {
	client := gosseract.NewClient()
	defer client.Close()

	fmt.Println("Processing page: ", n)

	var buf bytes.Buffer
	png.Encode(&buf, img)
	b := buf.Bytes()

	client.SetImageFromBytes(b)
	text, err := client.Text()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Processing complete: ", n)
	return text
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
