package middlewares

//TSVInsertVote - inserts vote
const TSVInsertVote = `
	INSERT INTO votes ("voice", "thread", "nickname") 
	VALUES ($1, $2, $3)
	`

//TSVUpdateVote - updates vote
const TSVUpdateVote = `
	UPDATE votes SET 
	"voice" = $1
	WHERE "thread" = $2 AND "nickname" = $3
	`

//TSVSelectVoteByID - selects vote by id
const TSVSelectVoteByID = `
	SELECT threads.id, threads.votes, votes.voice, usr.nickname
	FROM (SELECT 1) AS tmp_table
	LEFT JOIN threads ON threads.id = $1
	LEFT JOIN "users" AS usr ON usr.nickname = $2
	LEFT JOIN votes ON threads.id = votes.thread AND usr.nickname = votes.nickname
	`

//TSVSelectVoteBySlug - selects vote by slug
const TSVSelectVoteBySlug = `
	SELECT threads.id, threads.votes, votes.voice, usr.nickname
	FROM (SELECT 1) AS tmp_table
	LEFT JOIN threads ON threads.slug = $1
	LEFT JOIN users AS usr ON usr.nickname = $2
	LEFT JOIN votes ON threads.id = votes.thread AND usr.nickname = votes.nickname
	`

//TSVUpdateVotes - updates votes
const TSVUpdateVotes = `
	UPDATE threads SET
	"votes" = $1
	WHERE "id" = $2
	RETURNING "slug", "title", "id", "votes", "author", "created", "forum", "message"
	`

//TCMInsertValues - used for ThreadCreateMiddleware as request text
const TCMInsertValues = `
	INSERT INTO posts("author", "created", "forum", "message", "parent", "thread", "path") 
	VALUES($1, $2, $3, $4, $5, $6, (SELECT path FROM posts WHERE id = $5) || (select currval(pg_get_serial_sequence('posts', 'id'))))
	RETURNING "id", "created"
	`

//TCMFindForumByThread - used for ThreadCreateMiddleware as request text
const TCMFindForumByThread = `
	SELECT "forum", "id"
	FROM threads
	WHERE "id" = $1
	`

//TCMFindForumBySlug - used for ThreadCreateMiddleware as request text
const TCMFindForumBySlug = `
	SELECT "forum", "id"
	FROM threads
	WHERE "slug" = $1
	`

//TFBySlug - thread find by slug
const TFBySlug = `
	SELECT "author", "created", "forum", "id", "message", "slug", "title", "votes"
	FROM threads
	WHERE "slug" = $1
	`

//TFByID - thread find by ID
const TFByID = `
	SELECT "author", "created", "forum", "id", "message", "slug", "title", "votes"
	FROM threads
	WHERE "id" = $1
	`

//TCMFindPostByParent - finds parent by value
const TCMFindPostByParent = `
	SELECT "id"
	FROM posts
	WHERE "id" = $1 AND "thread" = $2
	`

//TCMUpdateForumPostsCount - updates posts count (forum)
const TCMUpdateForumPostsCount = `
	UPDATE forums
	SET "posts" = "posts" + 1
	WHERE "slug" = $1
	`

//TDPUpdateMessageID - updates thread message
const TDPUpdateMessageID = `
	UPDATE threads
	SET "message" = $1
	WHERE "id" = $2
	`

//TDPUpdateMessageSlug - updates thread message
const TDPUpdateMessageSlug = `
	UPDATE threads
	SET "message" = $1
	WHERE "slug" = $2
	`

//TDPUpdateTitleID - updates thread message
const TDPUpdateTitleID = `
	UPDATE threads
	SET "title" = $1
	WHERE "id" = $2
	`

//TDPUpdateTitleSlug - updates thread message
const TDPUpdateTitleSlug = `
	UPDATE threads
	SET "title" = $1
	WHERE "slug" = $2
	`

//TCMUpdatePath - updates post path
const TCMUpdatePath = `
	UPDATE posts
	SET "path" = concat("path", $1, ".")
	WHERE "id" = $2
	`

//TPSinceDescLimitTree - since desc limit tree
const TPSinceDescLimitTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND ("path" < (SELECT "path" FROM posts WHERE "id" = $2::TEXT::INTEGER))
	ORDER BY "path" DESC
	LIMIT $3::TEXT::INTEGER
	`

//TPSinceDescLimitParentTree - since desc limit parent tree
const TPSinceDescLimitParentTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE path[1] IN (
		SELECT "id"
		FROM posts
		WHERE "thread" = $1 AND "parent" = 0 AND "id" < (SELECT path[1] FROM posts WHERE "id" = $2::TEXT::INTEGER)
		ORDER BY "id" DESC
		LIMIT $3::TEXT::INTEGER
	)
	ORDER BY "path"
	`

//TPSinceDescLimitFlat - since desc limit flat
const TPSinceDescLimitFlat = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND "id" < $2::TEXT::INTEGER
	ORDER BY "id" DESC
	LIMIT $3::TEXT::INTEGER
	`

//TPSinceAscLimitTree - since asc limit tree
const TPSinceAscLimitTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND ("path" > (SELECT "path" FROM posts WHERE "id" = $2::TEXT::INTEGER))
	ORDER BY "path"
	LIMIT $3::TEXT::INTEGER
	`

//TPSinceAscLimitParentTree - since asc limit parent tree
const TPSinceAscLimitParentTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE path[1] IN (
		SELECT "id"
		FROM posts
		WHERE "thread" = $1 AND "parent" = 0 AND "id" > (SELECT path[1] FROM posts WHERE "id" = $2::TEXT::INTEGER)
		ORDER BY "id" LIMIT $3::TEXT::INTEGER
	)
	ORDER BY "path"
	`

//TPSinceAscLimitFlat - since asc limit flat
const TPSinceAscLimitFlat = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND "id" > $2::TEXT::INTEGER
	ORDER BY "id"
	LIMIT $3::TEXT::INTEGER
	`

//TPDescLimitTree - desc limit tree
const TPDescLimitTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 
	ORDER BY "path" DESC
	LIMIT $2::TEXT::INTEGER
	`

//TPDescLimitParentTree - desc limit parent tree
const TPDescLimitParentTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND path[1] IN (
		SELECT path[1]
		FROM posts
		WHERE "thread" = $1
		GROUP BY path[1]
		ORDER BY path[1] DESC
		LIMIT $2::TEXT::INTEGER
	)
	ORDER BY path[1] DESC, "path"
	`

//TPDescLimitFlat - desc limit flat
const TPDescLimitFlat = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1
	ORDER BY "id" DESC
	LIMIT $2::TEXT::INTEGER
	`

//TPAscLimitTree - asc limit tree
const TPAscLimitTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 
	ORDER BY "path"
	LIMIT $2::TEXT::INTEGER
	`

//TPAscLimitParentTree - asc limit parent tree
const TPAscLimitParentTree = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 AND path[1] IN (
		SELECT path[1] 
		FROM posts 
		WHERE "thread" = $1 
		GROUP BY path[1]
		ORDER BY path[1]
		LIMIT $2::TEXT::INTEGER
	)
	ORDER BY "path"
	`

//TPAscLimitFlat - asc limit flat
const TPAscLimitFlat = `
	SELECT "id", "author", "parent", "message", "forum", "thread", "created"
	FROM posts
	WHERE "thread" = $1 
	ORDER BY "id"
	LIMIT $2::TEXT::INTEGER
	`
