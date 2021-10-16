package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Projects() g.Node {
	return Page(
		"Workflow - Projects",
		"/projects",
		H1(g.Text(`Projects Page.`)),
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
