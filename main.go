package main

import (
	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

func main() {
	ActivateTheme("azure light")
	Pack(Button(Txt("Hello"), Command(func() { Destroy(App) })))
	App.Center().Wait()
}
