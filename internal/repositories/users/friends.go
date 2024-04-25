package users

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateRequest(ctx context.Context, req *models.FriendRequest) error {
	_, err := r.db.Exec(ctx, createFriendRequestStmt, req.UserID, req.FriendID, req.Status)
	return err
}

func (r *repository) AcceptRequest(ctx context.Context, req *models.FriendRequest) error {
	_, err := r.db.Exec(ctx, acceptFriendRequestStmt, req.UserID, req.FriendID, req.Status)
	return err
}

func (r *repository) GetFriends(ctx context.Context, userID int) ([]*models.User, error) {
	rows, err := r.db.Query(ctx, getFriendsStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.User
	for rows.Next() {
		friend := &userModel{}
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.Password, &friend.Role, &friend.City); err != nil {
			return nil, err
		}

		friends = append(friends, friend.convert())
	}

	return friends, nil
}

func (r *repository) GetFriendRequest(ctx context.Context, userID, friendID int) (*models.FriendRequest, error) {
	row := r.db.QueryRow(ctx, getFriendRequestStmt, userID, friendID)

	out := &models.FriendRequest{}
	if err := row.Scan(&out.UserID, &out.FriendID, &out.Status); err != nil {
		if err == pgx.ErrNoRows {
			return &models.FriendRequest{
				UserID:   userID,
				FriendID: friendID,
				Status:   "",
			}, nil
		}
		return nil, err
	}

	return out, nil
}

func (r *repository) GetSentRequestFriends(ctx context.Context, userID int) ([]*models.User, error) {
	rows, err := r.db.Query(ctx, getSentRequestFriendsStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.User
	for rows.Next() {
		friend := &userModel{}
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.City); err != nil {
			return nil, err
		}

		friends = append(friends, friend.convert())
	}

	return friends, nil
}

func (r *repository) GetReceivedRequestFriends(ctx context.Context, userID int) ([]*models.User, error) {
	rows, err := r.db.Query(ctx, getReceivedRequestFriendsStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.User
	for rows.Next() {
		friend := &userModel{}
		if err := rows.Scan(&friend.ID, &friend.Username, &friend.Email, &friend.City); err != nil {
			return nil, err
		}

		friends = append(friends, friend.convert())
	}

	return friends, nil
}
