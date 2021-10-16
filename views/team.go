package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Team() g.Node {
	return Page(
		"Workflow - Team",
		"/team",
		H1(g.Text(`Team Page.`)),
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
