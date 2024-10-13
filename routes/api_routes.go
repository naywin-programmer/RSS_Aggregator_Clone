package routes

import (
	"github.com/go-chi/chi/v5"
	ct "github.com/naywin-programmer/RSS_Aggregator/controllers/api"
	mi "github.com/naywin-programmer/RSS_Aggregator/middlewares"
)

func ApiRoutes(mainRouter *chi.Mux) {

	r := chi.NewRouter()
	defer mainRouter.Mount("/api/v1", r)

	// === NON AUTH ROUTES - START ===

	r.Post("/sign-up", ct.UserController.SignUp)
	r.Post("/sign-in", ct.UserController.SignIn)

	r.Get("/feeds", ct.FeedController.Index)
	r.Get("/feeds/{id}", ct.FeedController.Show)

	r.Get("/feed-follow-count/{feedId}", ct.FeedFollowController.GetFeedFollowerCount)

	// === NON AUTH ROUTES - END ===

	r.With(mi.BasicApiAuthMiddleware.Authenticate).Group(func(r chi.Router) {

		// === AUTH ROUTES - START ===

		r.Get("/users/feeds", ct.UserController.GetUserCreatedFeeds)
		r.Get("/users/feed-follow", ct.UserController.GetUserFollowedFeeds)

		r.Post("/feeds", ct.FeedController.Store)
		r.Delete("/feeds/{id}", ct.FeedController.Delete)

		r.Post("/feed-follow/{feedId}", ct.FeedFollowController.Store)
		r.Delete("/feed-follow/{feedId}", ct.FeedFollowController.Delete)

		r.Get("/posts", ct.PostController.Index)

		// === AUTH ROUTES - END ===
	})

}
