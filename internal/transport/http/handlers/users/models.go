package users

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
)

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r createUserRequest) convert() *models.User {
	return &models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type createUserResponse struct {
	response.Response
	Result createUserResp `json:"result"`
}

type createUserResp struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newCreateUserResponse(user *models.User) createUserResp {
	return createUserResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

type setCityRequest struct {
	City string `json:"city"`
}

type updateUsernameRequest struct {
	Username string `json:"username"`
}

type updatePasswordRequest struct {
	Password string `json:"password"`
}

type getUserByIDResponse struct {
	response.Response
	Result getUserByIDResp `json:"result"`
}

type getUserByIDResp struct {
	ID           int     `json:"id"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	City         *string `json:"city,omitempty"`
	ProfilePic   string  `json:"profile_pic"`
	FriendStatus string  `json:"friend_status"`
	BooksCount   int     `json:"books_count"`
	ShareCount   int     `json:"share_count"`
}

func newGetUserByIDResponse(user *models.UserWithCounts, req *models.FriendRequest) getUserByIDResp {
	friendStatus := "not_friends"

	if req != nil {
		if req.Status == models.FriendRequestAccepted {
			friendStatus = "friends"
		}

		if req.Status == models.FriendRequestSent {
			if req.UserID == user.ID {
				friendStatus = "request_received"
			} else {
				friendStatus = "request_sent"
			}
		}
	}

	return getUserByIDResp{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Password:     user.Password,
		City:         user.City,
		ProfilePic:   user.ProfilePic,
		FriendStatus: friendStatus,
		BooksCount:   user.BooksCount,
		ShareCount:   user.ShareCount,
	}
}

type listFriendsResponse struct {
	response.Response
	Result listFriendsResp `json:"result"`
}

type listFriendsResp struct {
	Friends []*listFriendsRespItem `json:"friends"`
}

type listFriendsRespItem struct {
	ID         int     `json:"id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	City       *string `json:"city,omitempty"`
	ProfilePic string  `json:"profile_pic"`
}

func newListFriendsResponse(friends []*models.User) listFriendsResp {
	resp := make([]*listFriendsRespItem, 0, len(friends))
	for _, friend := range friends {
		resp = append(resp, &listFriendsRespItem{
			ID:         friend.ID,
			Username:   friend.Username,
			Email:      friend.Email,
			City:       friend.City,
			ProfilePic: friend.ProfilePic,
		})
	}

	return listFriendsResp{
		Friends: resp,
	}
}

type addLikedBookRequest struct {
	BookID int `json:"book_id"`
}

type listLikedBooksResponse struct {
	response.Response
	Result []*models.Book `json:"result"`
}

type removeLikedBookRequest struct {
	BookID int `json:"book_id"`
}

type uploadProfilePicResponse struct {
	response.Response
	Result string `json:"result"`
}
