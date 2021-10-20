package views

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type NavLink struct {
	Path string
	Name string
}

// Page with a title, head, and a basic body layout.
func Page(title, path string, body ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			Script(Src("/static/js/dropdown.js")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/base.min.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/components.min.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/@tailwindcss/typography@0.4.0/dist/typography.min.css"), Type("text/css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/utilities.min.css"), Type("text/css")),
		},
		Body: []g.Node{
			c.Classes{"antialiased font-sans bg-gray-200 overflow-hidden": true},
			Navbar(path, []NavLink{
				{Path: "/team", Name: "Team"},
				{Path: "/projects", Name: "Projects"},
				{Path: "/calendar", Name: "Calendar"},
			}),
			Container(true,
				Prose(g.Group(body)),
			),
			PageFooter(),
		},
	})
}

// Navigation bar
func Navbar(currentPath string, links []NavLink) g.Node {
	return Nav(Class("bg-gray-800 sticky top-0 z-50"),
		Container(false,
			Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"),
				Div(Class("relative flex items-center justify-between h-16"),
					Div(Class("flex-1 flex items-center justify-center sm:items-stretch sm:justify-start"),
						Div(Class("flex-shrink-0 flex items-center"),
							Img(Class("block lg:hidden h-8 w-auto"), Src("https://tailwindui.com/img/logos/workflow-mark-indigo-500.svg"), Alt("Workflow")),
							Img(Class("hidden lg:block h-8 w-auto"), Src("https://tailwindui.com/img/logos/workflow-logo-indigo-500-mark-white-text.svg"), Alt("Workflow")),
						),
						Div(Class("hidden sm:block sm:ml-6"),
							Div(Class("flex space-x-4"),
								NavbarLink("/", "Home", currentPath == "/"),
								g.Group(g.Map(len(links), func(i int) g.Node {
									return NavbarLink(links[i].Path, links[i].Name, currentPath == links[i].Path)
								})),
							),
						),
					),
					Div(Class("absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0"),
						ButtonNotifications(),
						Div(Class("ml-3 relative"),
							Div(
								ButtonUserMenu(),
								DropdownMenu(),
							),
						),
					),
				),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return A(Href(path), g.Text(text),
		c.Classes{
			"text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium": true,
			"bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium":                           active,
		})
}

// Notification button
func ButtonNotifications() g.Node {
	return Button(c.Classes{
		"bg-gray-800 p-1 rounded-full text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white": true},
		Span(Class("sr-only"), g.Text("View notifications")),
		g.Raw(`<svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">`),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />`),
		g.Raw(`</svg>`),
	)
}

// User menu button
func ButtonUserMenu() g.Node {
	return Button(c.Classes{
		"bg-gray-800 flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white": true},
		ID("user-menu-button"),
		Aria("expanded", "false"),
		Aria("haspopup", "true"),
		Span(Class("sr-only"), g.Text("Open user menu")),
		Img(Class("h-8 w-8 rounded-full"),
			Src("https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"),
			Alt(""),
		),
	)
}

// Dropdown menu
func DropdownMenu() g.Node {
	return Div(DataAttr("dropdown-toggle", "dropdown"),
		Class("origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"),
		Role("menu"),
		Aria("orientation", "vertical"),
		Aria("labelledby", "user-menu-button"),
		TabIndex("-1"),
		A(Href("#"), g.Text("Your Profile"), c.Classes{"block px-4 py-2 text-sm text-gray-700": true}, Role("menuitem"), TabIndex("-1"), ID("user-menu-item-0")),
		A(Href("#"), g.Text("Settings"), c.Classes{"block px-4 py-2 text-sm text-gray-700": true}, Role("menuitem"), TabIndex("-1"), ID("user-menu-item-1")),
		A(Href("#"), g.Text("Sign out"), c.Classes{"block px-4 py-2 text-sm text-gray-700": true}, Role("menuitem"), TabIndex("-1"), ID("user-menu-item-2")),
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

// Footer
func PageFooter() g.Node {
	return Footer(Class("bg-gray-800"),
		Container(false,
			Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"),
				Div(Class("relative flex items-center justify-between h-16")),
			),
		),
	)
}
