package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Page with a title, head, and a basic body layout.
func Page(title, path string, body ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/base.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/components.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/@tailwindcss/typography@0.4.0/dist/typography.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/utilities.min.css")),
		},
		Body: []g.Node{
			Navbar(path),
			Container(true,
				Prose(g.Group(body)),
			),
		},
	})
}

func Navbar(path string) g.Node {
	return Nav(Class("bg-gray-800"),
		Container(false,
			Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), // mian div
				Div(Class("relative flex items-center justify-between h-16"),
					Div(Class("flex-1 flex items-center justify-center sm:items-stretch sm:justify-start"),
						Div(Class("flex-shrink-0 flex items-center"),
							Img(Class("block lg:hidden h-8 w-auto"), Src("https://tailwindui.com/img/logos/workflow-mark-indigo-500.svg"), Alt("Workflow")),
							Img(Class("hidden lg:block h-8 w-auto"), Src("https://tailwindui.com/img/logos/workflow-logo-indigo-500-mark-white-text.svg"), Alt("Workflow")),
						),
						Div(Class("hidden sm:block sm:ml-6"),
							Div(Class("flex space-x-4"),
								NavbarLinkCurrent("/", "Home", path),
								NavbarLink("/team", "Team", path),
								NavbarLink("/projects", "Projects", path),
								NavbarLink("/calendar", "Calendar", path),
							),
						),
					),
				),
			),
		),
	)
}

func NavbarLinkCurrent(path, text, currentPath string) g.Node {
	//active := path == currentPath
	return A(Href(path), g.Text(text),
		c.Classes{
			"bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium": true,
		},
		Aria("current", "page"),
	)
}

func NavbarLink(path, text, currentPath string) g.Node {
	//active := path == currentPath
	return A(Href(path), g.Text(text),
		c.Classes{
			"text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium": true,
		},
	)
}

func Container(padY bool, children ...g.Node) g.Node {
	return Div(
		c.Classes{
			"max-w-7xl mx-auto px-4 sm:px-6 lg:px-8": true,
			"py-4 sm:py-6 lg:py-8":                   padY,
		},
		g.Group(children),
	)
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose lg:prose-lg xl:prose-xl prose-indigo"), g.Group(children))
}
