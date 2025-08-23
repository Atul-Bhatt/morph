package main

import tk "modernc.org/tk9.0"

type App struct {
	name string
}

func (me *App) Run() {
	tk.App.SetResizable(false, false)
	tk.App.Center()
	tk.WmDeiconify(tk.App)
	tk.App.Wait()
}

func (me *App) onQuit() { tk.Destroy(tk.App) }
