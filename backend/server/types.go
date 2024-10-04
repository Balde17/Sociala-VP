package server

import (
	"backend/app/internals/handlers/auth"
	"backend/app/internals/handlers/chat"
	"backend/app/internals/handlers/comment"
	"backend/app/internals/handlers/events"
	"backend/app/internals/handlers/followers"
	"backend/app/internals/handlers/group"
	"backend/app/internals/handlers/invite"
	"backend/app/internals/handlers/like"
	"backend/app/internals/handlers/notifications"
	"backend/app/internals/handlers/post"
	"backend/app/internals/handlers/profile"
	"backend/app/internals/handlers/request"
	"backend/app/internals/handlers/ws"

	"net/http"
)

type EndPoint struct {
	Path    string
	Handler RouteHandler
	Method  []string
}

type RouteHandler func(http.ResponseWriter, *http.Request)

type Middleware func(RouteHandler) RouteHandler

type route struct {
	method  string
	pattern string
	handler RouteHandler
}

type Router struct {
	routes      []route
	middlewares []Middleware
}

var (
	Endpoints = []EndPoint{
		{Path: "/login", Handler: auth.Login, Method: []string{"POST"}},
		{Path: "/register", Handler: auth.Register, Method: []string{"POST"}},
	}

	Protected = []EndPoint{
		{Path: "/create-post", Handler: post.CreatePost, Method: []string{"POST"}},
		{Path: "/get-post", Handler: post.ListPost, Method: []string{"GET"}},

		{Path: "/create-comment", Handler: comment.CreateComment, Method: []string{"POST"}},
		{Path: "/get-comment", Handler: comment.ListComment, Method: []string{"GET"}},

		{Path: "/get-like", Handler: like.ListLike, Method: []string{"GET"}},

		{Path: "/invite", Handler: invite.InviteUser, Method: []string{"POST"}},
		{Path: "/update-invite", Handler: invite.UpdateInvite, Method: []string{"POST"}},
		{Path: "/delete-invite", Handler: invite.DeleteInvite, Method: []string{"POST"}},
		{Path: "/get-invite", Handler: group.UserIsInvited, Method: []string{"GET"}},
		{Path: "/get-listInvite", Handler: group.ListInviteToGroup, Method: []string{"GET"}},
		{Path: "/get-memberGroup", Handler: group.MemberGroup, Method: []string{"GET"}},
		{Path: "/get-ismemberGroup", Handler: group.IsMemberGroup, Method: []string{"GET"}},
		{Path: "/delete-follow", Handler: invite.DeleteInviteFollow, Method: []string{"POST"}},

		{Path: "/create-event", Handler: events.CreateEvent, Method: []string{"POST"}},
		{Path: "/get-events", Handler: events.ListEvent, Method: []string{"GET"}},
		{Path: "/create-eventOption", Handler: events.CreateEventOption, Method: []string{"POST"}},
		{Path: "/get-event", Handler: events.ListEventGroup, Method: []string{"GET"}},

		{Path: "/create-group", Handler: group.CreateGroup, Method: []string{"POST"}},
		{Path: "/get-group", Handler: group.ListGroup, Method: []string{"GET"}},
		{Path: "/get-groupById", Handler: group.ListIdGroup, Method: []string{"GET"}},
		{Path: "/get-othergroup", Handler: group.ListOtherGroup, Method: []string{"GET"}},
		{Path: "/get-groupUser", Handler: group.UserGroup, Method: []string{"GET"}},

		{Path: "/get-profilUser", Handler: profile.UserProfile, Method: []string{"GET"}},
		{Path: "/get-postUser", Handler: profile.ListPostUser, Method: []string{"GET"}},
		{Path: "/get-follower", Handler: profile.ListFollower, Method: []string{"GET"}},
		{Path: "/get-followee", Handler: profile.ListFollowee, Method: []string{"GET"}},
		{Path: "/get-otherUser", Handler: profile.ListOtherUsers, Method: []string{"GET"}},

		{Path: "/get-userId", Handler: profile.UserProfileId, Method: []string{"GET"}},
		// {Path: "/get-groupUser", Handler: profile.GroupUserList, Method: []string{"GET"}},
		{Path: "/update-profil", Handler: profile.UpdateProfil, Method: []string{"POST"}},
		{Path: "/get-friends", Handler: profile.ListFriends, Method: []string{"GET"}},
		{Path: "/get-followerId", Handler: followers.FollowerById, Method: []string{"GET"}},
		{Path: "/post-follower", Handler: followers.FollowUser, Method: []string{"POST"}},

		{Path: "/get-request", Handler: request.Request, Method: []string{"GET"}},
		{Path: "/update-request", Handler: request.UpdateProfil, Method: []string{"POST"}},
		{Path: "/delete-request", Handler: request.DeleteRequest, Method: []string{"POST"}},

		{Path: "/list-notificationRequest", Handler: notifications.ListNotification, Method: []string{"GET"}},
		{Path: "/list-notificationEvent", Handler: notifications.ListNotificationEvent, Method: []string{"GET"}},
		{Path: "/list-notificationGroup", Handler: notifications.ListNotificationGroup, Method: []string{"GET"}},

		{Path: "/delete-notification", Handler: notifications.DeleteNotification, Method: []string{"POST"}},
		{Path: "/delete-notif", Handler: notifications.DeleteNotif, Method: []string{"POST"}},
		{Path: "/update-notif", Handler: notifications.UpdateNotification, Method: []string{"POST"}},

		{Path: "/get-discussions", Handler: chat.ListMessage, Method: []string{"GET"}},
		{Path: "/get-discussions-group", Handler: chat.ListMessageGroup, Method: []string{"GET"}},

		{Path: "/ws", Handler: ws.WS, Method: []string{"GET"}},
	}
)
