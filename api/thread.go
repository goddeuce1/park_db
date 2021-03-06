package api

import (
	"park_base/park_db/models"

	mw "park_base/park_db/middlewares"

	"github.com/valyala/fasthttp"
)

//log.Println(string(ctx.Request.Header.RequestURI()))

//ThreadCreate - creates posts for thread
func ThreadCreate(ctx *fasthttp.RequestCtx) {
	posts := models.Posts{}
	posts.UnmarshalJSON(ctx.PostBody())

	thread := ctx.UserValue("slug_or_id").(string)

	response, error := mw.ThreadCreateMiddleware(posts, thread)

	if error == nil {
		mw.SetHeaders(ctx, fasthttp.StatusCreated)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if error == models.ErrParentNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusConflict)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if error == models.ErrUserNotFound || error == models.ErrThreadNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusNotFound)
		result, _ := error.MarshalJSON()
		ctx.Write(result)
	}

	return
}

//ThreadVote - sets +-1 rating to thread
func ThreadVote(ctx *fasthttp.RequestCtx) {
	vote := models.Vote{}
	vote.UnmarshalJSON(ctx.PostBody())

	threadSlug := ctx.UserValue("slug_or_id").(string)

	response, error := mw.ThreadSlugVoteMiddleware(&vote, threadSlug)

	if error == nil {
		mw.SetHeaders(ctx, fasthttp.StatusOK)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if error == models.ErrThreadNotFound || error == models.ErrUserNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusNotFound)
		result, _ := error.MarshalJSON()
		ctx.Write(result)
	}

	return
}

//ThreadDetailsGet - get info about thread by slug/id
func ThreadDetailsGet(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)

	response, err := mw.ThreadDetailsGetMiddleware(slug)

	if err == nil {
		mw.SetHeaders(ctx, fasthttp.StatusOK)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if err == models.ErrThreadNotFound || err == models.ErrUserNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusNotFound)
		result, _ := err.MarshalJSON()
		ctx.Write(result)
	}

	return
}

//ThreadDetailsPost - updates thread info
func ThreadDetailsPost(ctx *fasthttp.RequestCtx) {
	threadUpdate := models.ThreadUpdate{}
	threadUpdate.UnmarshalJSON(ctx.PostBody())

	threadSlug := ctx.UserValue("slug_or_id").(string)

	response, error := mw.ThreadDetailsPostMiddleware(&threadUpdate, threadSlug)

	if error == nil {
		mw.SetHeaders(ctx, fasthttp.StatusOK)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if error == models.ErrThreadNotFound || error == models.ErrUserNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusNotFound)
		result, _ := error.MarshalJSON()
		ctx.Write(result)
	}

	return
}

//ThreadPosts - returns thread posts
func ThreadPosts(ctx *fasthttp.RequestCtx) {
	slug := ctx.UserValue("slug_or_id").(string)
	limit := string(ctx.FormValue("limit"))
	since := string(ctx.FormValue("since"))
	sort := string(ctx.FormValue("sort"))
	desc := string(ctx.FormValue("desc"))

	response, error := mw.ThreadPostsMiddleware(slug, limit, since, sort, desc)

	if error == nil {
		mw.SetHeaders(ctx, fasthttp.StatusOK)
		result, _ := response.MarshalJSON()
		ctx.Write(result)

	} else if error == models.ErrThreadNotFound || error == models.ErrUserNotFound {
		mw.SetHeaders(ctx, fasthttp.StatusNotFound)
		result, _ := error.MarshalJSON()
		ctx.Write(result)
	}

	return
}
