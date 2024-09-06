import * as htmx from "htmx.org"
import * as _hyperscript from "hyperscript.org"
import { openRandomWebsite } from "./useless_web"

declare global {
	interface Window {
		htmx: typeof htmx
		_hyperscript: typeof _hyperscript
		openRandomWebsite: typeof openRandomWebsite
	}
}

window.htmx = htmx
window._hyperscript = _hyperscript
window.openRandomWebsite = openRandomWebsite

_hyperscript.browserInit()
