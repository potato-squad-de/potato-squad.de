package ui

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"github.com/potato-squad-org/site/globals"
)

func Index() g.Node {
	return g.Group([]g.Node{
		Img(
			Src("/images/logo.png"),
			Alt("logo"),
		),

		Div(
			Class("flex flex-col items-center gap-4 font-riffic text-2xl uppercase text-white md:flex-row"),

			A(
				Href(globals.DiscordInviteURL),
				Target("_blank"),
				Rel("noopener noreferrer"),
				Class("flex min-w-fit items-center gap-2 rounded-lg bg-[#7289da] p-2 transition duration-100 ease-in-out hover:bg-[#97a4d4]"),

				Span(
					Class("icon-[mdi--discord]"),
				),
				g.Text("Join"),
			),
			Button(
				g.Attr("_", "on click call openRandomWebsite()"),
				Class("min-w-fit rounded-lg bg-[#4b0082] p-2 uppercase transition duration-100 ease-in-out hover:bg-[#7700ce]"),
				g.Text("Traust dich eh nicht"),
			),
		),
	})
}
