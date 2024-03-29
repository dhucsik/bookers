package users

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	SetCity(ctx context.Context, userID int, city string) error
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, userID int) error
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
		return nil, err
	}

	return model.convert(), nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	row := r.db.QueryRow(ctx, getUserByUsernameStmt, username)

	model := &userModel{}
	if err := row.Scan(&model.ID, &model.Username, &model.Email, &model.Password, &model.Role, &model.City); err != nil {
		return nil, err
	}

	return model.convert(), nil
}

func (r *repository) DeleteUser(ctx context.Context, userID int) error {
	_, err := r.db.Exec(ctx, deleteUserStmt, userID)
	return err
}
