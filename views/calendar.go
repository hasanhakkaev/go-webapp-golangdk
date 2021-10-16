package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Calendar() g.Node {
	return Page(
		"Workflow - Calendar",
		"/calendar",
		H1(g.Text(`Calendar Page.`)),
		P(g.Text(`------------------------------------------------------------.`)),
		P(g.Raw(`Then we created the <em>go-web-app</em> app, and now we don't! 😬`)),
		P(g.Text(`------------------------------------------------------------.`)),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
	)
}
