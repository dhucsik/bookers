package users

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	SetCity(ctx context.Context, userID int, city string) error
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetUsersByIDs(ctx context.Context, ids []int) ([]*models.User, error)
	CreateRequest(ctx context.Context, req *models.FriendRequest) error
	AcceptRequest(ctx context.Context, req *models.FriendRequest) error
	GetFriends(ctx context.Context, userID int) ([]*models.User, error)
	GetFriendRequest(ctx context.Context, userID, friendID int) (*models.FriendRequest, error)
	GetSentRequestFriends(ctx context.Context, userID int) ([]*models.User, error)
	GetReceivedRequestFriends(ctx context.Context, userID int) ([]*models.User, error)
	UpdateUsername(ctx context.Context, userID int, username string) error
	UpdatePassword(ctx context.Context, userID int, password string) error
	SetPhone(ctx context.Context, userID int, phone string) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	model := convertUser(user)

	row := r.db.QueryRow(ctx, createUserStmt, model.Username, model.Email, model.Password)
	if err := row.Scan(&model.ID); err != nil {
		return nil, err
	}

	return model.convert(), nil
}

func (r *repository) SetCity(ctx context.Context, userID int, city string) error {
	_, err := r.db.Exec(ctx, setCityStmt, userID, city)
	return err
}

func (r *repository) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	row := r.db.QueryRow(ctx, getUserByIDStmt, userID)

	model := &userModel{}
	if err := row.Scan(&model.ID, &model.Username, &model.Email, &model.Password, &model.Role, &model.City); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return model.convert(), nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	row := r.db.QueryRow(ctx, getUserByUsernameStmt, username)

	model := &userModel{}
	if err := row.Scan(&model.ID, &model.Username, &model.Email, &model.Password, &model.Role, &model.City); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return model.convert(), nil
}

func (r *repository) DeleteUser(ctx context.Context, userID int) error {
	_, err := r.db.Exec(ctx, deleteUserStmt, userID)
	return err
}

func (r *repository) GetUsersByIDs(ctx context.Context, ids []int) ([]*models.User, error) {
	rows, err := r.db.Query(ctx, getUsersByIDsStmt, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.User
	for rows.Next() {
		model := &userModel{}
		if err := rows.Scan(&model.ID, &model.Username, &model.Email, &model.Role, &model.City); err != nil {
			return nil, err
		}

		out = append(out, model.convert())
	}

	return out, nil
}

func (r *repository) UpdateUsername(ctx context.Context, userID int, username string) error {
	_, err := r.db.Exec(ctx, updateUsernameStmt, userID, username)
	return err
}

func (r *repository) UpdatePassword(ctx context.Context, userID int, password string) error {
	_, err := r.db.Exec(ctx, updatePasswordStmt, userID, password)
	return err
}

func (r *repository) SetPhone(ctx context.Context, userID int, phone string) error {
	_, err := r.db.Exec(ctx, setPhoneNumberStmt, userID, phone)
	return err
}
