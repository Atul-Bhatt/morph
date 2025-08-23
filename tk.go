package main

import 
(
	tk "modernc.org/tk9.0"
	//"time"
)

type App struct {
	name string
	pdfText *tk.TextWidget
}

func (me *App) Run() {
	tk.App.SetResizable(false, false)
	tk.App.Center()
	tk.WmDeiconify(tk.App)
	tk.App.Wait()
}

/*func (me *App) update(processPages func()) {
	processPages()
	tk.TclAfter(time.Second*10, me.update(processPages))
}*/

func (me *App) onQuit() { tk.Destroy(tk.App) }
