package shared

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	root "github.com/potato-squad-org/site"
	"golang.org/x/exp/slices"
)

type PageProps struct {
	Title string
	Body  g.Node
}

var titleExcludes = []string{
	root.Meta().Name(),
	"Home",
}

func title(title string) string {
	if slices.Contains(titleExcludes, title) {
		return root.Meta().Name()
	}

	return fmt.Sprintf("%s | %s", title, root.Meta().Name())
}

func Page(props PageProps) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: title(props.Title),
		Head: []g.Node{
			Script(Defer(), Src("/main.min.js")),
			Link(Rel("stylesheet"), Href("/main.min.css")),
			Link(Rel("shortcut icon"), Href("/favicon.png"), Type("image/x-icon")),
		},
		Body: []g.Node{
			Class("h-screen select-none bg-[#99ccff] flex flex-col items-center justify-center gap-4"),

			props.Body,
		},
	})
}
