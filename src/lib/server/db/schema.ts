import { integer, sqliteTable, text } from 'drizzle-orm/sqlite-core'

export const user = sqliteTable('user', {
	id: integer('id').primaryKey(),
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

export const userLinks = sqliteTable('user_links', {
	userId: integer('user_id')
		.notNull()
		.references(() => user.id),

})
