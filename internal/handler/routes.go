// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "blog_server/internal/handler/admin"
	article "blog_server/internal/handler/article"
	blog "blog_server/internal/handler/blog"
	comment "blog_server/internal/handler/comment"
	dashboard "blog_server/internal/handler/dashboard"
	exhibition "blog_server/internal/handler/exhibition"
	like "blog_server/internal/handler/like"
	login "blog_server/internal/handler/login"
	notice "blog_server/internal/handler/notice"
	search "blog_server/internal/handler/search"
	star "blog_server/internal/handler/star"
	tag "blog_server/internal/handler/tag"
	tray "blog_server/internal/handler/tray"
	user "blog_server/internal/handler/user"
	"blog_server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/blog/create",
				Handler: blog.CreateBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/blog/update",
				Handler: blog.UpdateBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/delete",
				Handler: blog.DeleteBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/info",
				Handler: blog.BlogInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/thumbsup",
				Handler: blog.ThumbsUpBlogHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/create",
				Handler: comment.CreateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/delete",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/thumbsup",
				Handler: comment.ThumbsupCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/info",
				Handler: comment.InfoCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/dashboard",
				Handler: dashboard.DashboardHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/like",
				Handler: like.LikeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/likes/info",
				Handler: like.LikesInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/likes/info/ids",
				Handler: like.LikesInfoIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/follow",
				Handler: like.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/follow/info",
				Handler: like.FollowInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: login.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: exhibition.UploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/cos/upload",
				Handler: exhibition.CosUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/create",
				Handler: exhibition.CreateExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/exhibition/info",
				Handler: exhibition.ExhibitionInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/update",
				Handler: exhibition.UpdateExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/approval",
				Handler: exhibition.ApprovalExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/del",
				Handler: exhibition.DelExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/upload/del",
				Handler: exhibition.DelUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/update/download",
				Handler: exhibition.UpdateDownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/exhibition/similar",
				Handler: exhibition.SimilarExhibitionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/info",
				Handler: user.UserinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/update",
				Handler: user.UpdateUserinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/update/password",
				Handler: user.UpdateUserPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/star",
				Handler: star.StarHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/tags/create",
				Handler: tag.CreateTagHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tags/info",
				Handler: tag.TagsInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: search.SearchHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/tray",
				Handler: tray.TrayExhibitionHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/admin/exhibitions",
				Handler: admin.ExhibitionAdminInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/admin/stat",
				Handler: admin.StatAdminHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/article/create",
				Handler: article.CreateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/update",
				Handler: article.UpdateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/article/delete",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/article/info",
				Handler: article.InfoArticleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/notice/create",
				Handler: notice.CreateNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/notice/update",
				Handler: notice.UpdateNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/delete",
				Handler: notice.DeleteNoticeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/notice/info",
				Handler: notice.InfoNoticeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
