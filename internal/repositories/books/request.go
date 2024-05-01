package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
)

func (r *repository) CreateRequest(ctx context.Context, req *models.ShareRequest) error {
	_, err := r.db.Exec(ctx, createNewRequestStmt, req.SenderID, req.ReceiverID, req.SenderBookID, req.ReceiverBookID, req.SenderStatus, req.ReceiverStatus)
	return err
}

func (r *repository) UpdateRequest(ctx context.Context, req *models.ShareRequest) error {
	if req.SenderStatus == models.StatusSenderProved && req.ReceiverStatus == models.StatusReceiverProved {
		return r.finishRequest(ctx, req)
	}

	_, err := r.db.Exec(ctx, updateRequestStmt, req.ID, req.SenderStatus, req.ReceiverStatus, req.SenderBookID)
	return err
}

func (r *repository) finishRequest(ctx context.Context, req *models.ShareRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, updateRequestStmt, req.ID, req.SenderStatus, req.ReceiverStatus, req.SenderBookID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	id := req.ID
	req = &models.ShareRequest{}
	err = tx.QueryRow(ctx, getRequestStmt, id).Scan(&req.ID, &req.SenderID, &req.ReceiverID, &req.SenderBookID, &req.ReceiverBookID, &req.SenderStatus, &req.ReceiverStatus, &req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, updateStockBookUserStmt, req.SenderBookID, req.ReceiverID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, updateStockBookUserStmt, req.ReceiverBookID, req.SenderID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *repository) GetRequest(ctx context.Context, id int) (*models.ShareRequest, error) {
	req := &models.ShareRequest{}
	err := r.db.QueryRow(ctx, getRequestStmt, id).Scan(&req.ID, &req.SenderID, &req.ReceiverID, &req.SenderBookID, &req.ReceiverBookID, &req.SenderStatus, &req.ReceiverStatus, &req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.ErrRequestNotFound
		}

		return nil, err
	}

	return req, nil
}

func (r *repository) ListRequests(ctx context.Context, userID int) ([]*models.ShareRequest, error) {
	rows, err := r.db.Query(ctx, getRequestsStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.ShareRequest
	for rows.Next() {
		req := &models.ShareRequest{}
		err := rows.Scan(&req.ID, &req.SenderID, &req.ReceiverID, &req.SenderBookID, &req.ReceiverBookID, &req.SenderStatus, &req.ReceiverStatus, &req.CreatedAt, &req.UpdatedAt)
		if err != nil {
			return nil, err
		}

		out = append(out, req)
	}

	return out, nil
}

func (r *repository) GetSuccessRequestCount(ctx context.Context, userID int) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, getSuccessRequestCountStmt, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
