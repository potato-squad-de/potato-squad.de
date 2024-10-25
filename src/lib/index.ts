// place files you want to import through the `$lib` alias in this folder.

export type Social = {
	platform: string
	username: string
	url: string
}

export type Ref = {
	name: string
	url: string
}

export type CardProps = {
	name: string
	aliases?: string[]
	imgUrl?: string
	bio?: string
	socials?: Social[]
	otherRefs?: Ref[]
}
