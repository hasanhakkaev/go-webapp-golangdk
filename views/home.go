package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Home() g.Node {
	return Page(
		"Workflow - Home",
		"/",
		H1(g.Text(`Home Page`)),
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
