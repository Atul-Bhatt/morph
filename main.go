package main

import (
	"fmt"
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	fmt.Println("Program started!")
	client := gosseract.NewClient()
	defer client.Close()

	client.SetPDF("sample_text.pdf")
	text, err := client.Text()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
	fmt.Println("Program was immensely successful!")

	t := Text(Font("helvetica", 10), Padx("2m"), Pady("2m"))
	t.InsertML(text)
	ActivateTheme("azure light")
	Grid(t, Sticky("news"), Pady("2m"))
	Grid(TExit(), Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	App.Center().Wait()
}
