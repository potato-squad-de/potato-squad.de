import { integer, sqliteTable, text } from 'drizzle-orm/sqlite-core'

export const user = sqliteTable('user', {
	id: integer('id').primaryKey(),
	email: text('user_email').unique().notNull(),
	passwordHash: text('password_hash').notNull()
})

export const userProfile = sqliteTable('user_profile', {
	userId: integer('user_id')
		.notNull()
		.unique()
		.references(() => user.id),
	name: text('display_name').notNull().default('Anonymous'),
	bio: text('bio').notNull().default(''),
	avatar_ref: text('avatar_ref')
		.notNull()
		.default('/images/avatar_placeholder.png')
})

export const userAliases = sqliteTable('user_aliases', {
	userId: integer('user_id')
		.notNull()
		.references(() => user.id),
	alias: text('alias').notNull()
})

export const socials = sqliteTable('socials', {
	id: integer('id').primaryKey(),
	platform: text('platform').notNull(),
	baseURL: text('url').notNull()
})

export const userSocials = sqliteTable('user_socials', {
	userId: integer('user_id')
		.notNull()
		.references(() => user.id),
	platform: integer('platform')
		.notNull()
		.references(() => socials.id),
	url: text('url').notNull()
})

export const userReflinks = sqliteTable('user_reflinks', {
	userId: integer('user_id')
		.notNull()
		.references(() => user.id),
	reflinkId: integer('reflink_id')
		.notNull()
		.references(() => user.id)
})
