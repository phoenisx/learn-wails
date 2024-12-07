package main

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func addEventListeners(app *application.App) {
	app.OnEvent("calculator:test", func(e *application.CustomEvent) {
		app.Logger.Info(">>>>>>>>>", fmt.Sprintf("%v, %v", e, e.Data))
	})
}
