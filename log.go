package main

func (app *App) LogInfo(info ...interface{}) {
	app.infoLog.Println(info...)
}

func (app *App) LogError(err error) {
	app.errorLog.Panicln(err.Error()[:100])
}
