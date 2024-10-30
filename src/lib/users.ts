export interface User {
	displayName?: string
	aliases?: string[]
	bio?: string
	avatarRef?: string
	potatoAvatarRef?: string
	links?: {
		url: string
		name: string
	}[]
}

export const users: User[] = []
