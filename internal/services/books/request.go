package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

func (s *service) CreateRequest(ctx context.Context, userID, stockBookID int) error {
	stockBook, err := s.bookRepo.GetStockBook(ctx, stockBookID)
	if err != nil {
		return err
	}

	if stockBook.UserID == userID {
		return errors.ErrForbiddenForUser
	}

	return s.bookRepo.CreateRequest(ctx, &models.ShareRequest{
		SenderID:       userID,
		ReceiverBookID: stockBookID,
		ReceiverID:     stockBook.UserID,
		SenderStatus:   models.StatusCreated,
		ReceiverStatus: models.StatusCreated,
	})
}

func (s *service) CancelRequest(ctx context.Context, userID, id int) error {
	req, err := s.bookRepo.GetRequest(ctx, id)
	if err != nil {
		return err
	}

	if req.SenderID != userID || req.ReceiverID != userID {
		return errors.ErrForbiddenForUser
	}

	req = &models.ShareRequest{
		ID:             id,
		SenderBookID:   req.SenderBookID,
		SenderStatus:   models.StatusCanceled,
		ReceiverStatus: models.StatusCanceled,
	}

	return s.bookRepo.UpdateRequest(ctx, req)
}

func (s *service) ReceiverRequested(ctx context.Context, stockBookID, userID, id int) error {
	req, err := s.bookRepo.GetRequest(ctx, id)
	if err != nil {
		return err
	}

	stockBook, err := s.bookRepo.GetStockBook(ctx, stockBookID)
	if err != nil {
		return err
	}

	if userID != req.ReceiverID || stockBook.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	if req.ReceiverStatus != models.StatusCreated || req.SenderStatus != models.StatusCreated {
		return errors.ErrWrongStatus
	}

	req = &models.ShareRequest{
		ID:             id,
		SenderBookID:   stockBookID,
		SenderStatus:   models.StatusCreated,
		ReceiverStatus: models.StatusReceiverRequested,
	}

	return s.bookRepo.UpdateRequest(ctx, req)
}

func (s *service) SenderAccepted(ctx context.Context, userID, id int) error {
	req, err := s.bookRepo.GetRequest(ctx, id)
	if err != nil {
		return err
	}

	if userID != req.SenderID {
		return errors.ErrForbiddenForUser
	}

	if req.SenderStatus != models.StatusCreated || req.ReceiverStatus != models.StatusReceiverRequested {
		return errors.ErrWrongStatus
	}

	req = &models.ShareRequest{
		ID:             id,
		SenderBookID:   req.SenderBookID,
		SenderStatus:   models.StatusSenderAccepted,
		ReceiverStatus: models.StatusReceiverRequested,
	}

	return s.bookRepo.UpdateRequest(ctx, req)
}

func (s *service) ApproveRequest(ctx context.Context, userID, id int) error {
	req, err := s.bookRepo.GetRequest(ctx, id)
	if err != nil {
		return err
	}

	senderStatus := req.SenderStatus
	receiverStatus := req.ReceiverStatus

	if senderStatus == models.StatusSenderProved && receiverStatus == models.StatusReceiverProved {
		return errors.ErrWrongStatus
	}

	if senderStatus == models.StatusCreated || senderStatus == models.StatusCanceled {
		return errors.ErrWrongStatus
	}

	if receiverStatus == models.StatusCreated || receiverStatus == models.StatusCanceled {
		return errors.ErrWrongStatus
	}

	if userID == req.ReceiverID {
		receiverStatus = models.StatusReceiverProved
	} else if userID == req.SenderID {
		senderStatus = models.StatusSenderProved
	} else {
		return errors.ErrForbiddenForUser
	}

	req = &models.ShareRequest{
		ID:             id,
		SenderBookID:   req.SenderBookID,
		SenderStatus:   senderStatus,
		ReceiverStatus: receiverStatus,
	}

	return s.bookRepo.UpdateRequest(ctx, req)
}

func (s *service) GetRequest(ctx context.Context, id int) (*models.RequestWithFields, error) {
	req, err := s.bookRepo.GetRequest(ctx, id)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.GetUsersByIDs(ctx, []int{req.SenderID, req.ReceiverID})
	if err != nil {
		return nil, err
	}

	usersMap := lo.SliceToMap(users, func(item *models.User) (int, *models.UserWithoutPassword) {
		return item.ID, item.ToUserWithoutPassword()
	})

	books, err := s.bookRepo.GetBooksByStockIDs(ctx, []int{req.SenderBookID, req.ReceiverBookID})
	if err != nil {
		return nil, err
	}

	booksMap := lo.SliceToMap(books, func(item *models.Book) (int, *models.Book) {
		return item.ID, item
	})

	return &models.RequestWithFields{
		ID:             req.ID,
		Sender:         usersMap[req.SenderID],
		Receiver:       usersMap[req.ReceiverID],
		SenderBook:     booksMap[req.SenderBookID],
		ReceiverBook:   booksMap[req.ReceiverBookID],
		SenderStatus:   req.SenderStatus,
		ReceiverStatus: req.ReceiverStatus,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}, nil
}

func (s *service) GetRequests(ctx context.Context, userID int) ([]*models.RequestWithFields, error) {
	reqs, err := s.bookRepo.ListRequests(ctx, userID)
	if err != nil {
		return nil, err
	}

	ids := lo.Map(reqs, func(req *models.ShareRequest, _ int) int {
		if req.SenderID == userID {
			return req.ReceiverID
		}
		return req.SenderID
	})

	ids = append(ids, userID)

	rBookIDs := lo.Map(reqs, func(req *models.ShareRequest, _ int) int {
		return req.ReceiverBookID
	})

	lBookIDs := lo.Map(reqs, func(req *models.ShareRequest, _ int) int {
		return req.SenderBookID
	})

	bookIDs := lo.Union(rBookIDs, lBookIDs)

	users, err := s.userRepo.GetUsersByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	usersMap := lo.SliceToMap(users, func(item *models.User) (int, *models.UserWithoutPassword) {
		return item.ID, item.ToUserWithoutPassword()
	})

	books, err := s.bookRepo.GetBooksByStockIDs(ctx, bookIDs)
	if err != nil {
		return nil, err
	}

	booksMap := lo.SliceToMap(books, func(item *models.Book) (int, *models.Book) {
		return item.ID, item
	})

	out := lo.Map(reqs, func(req *models.ShareRequest, _ int) *models.RequestWithFields {
		return &models.RequestWithFields{
			ID:             req.ID,
			Sender:         usersMap[req.SenderID],
			Receiver:       usersMap[req.ReceiverID],
			SenderBook:     booksMap[req.SenderBookID],
			ReceiverBook:   booksMap[req.ReceiverBookID],
			SenderStatus:   req.SenderStatus,
			ReceiverStatus: req.ReceiverStatus,
			CreatedAt:      req.CreatedAt,
			UpdatedAt:      req.UpdatedAt,
		}
	})

	return out, nil
}
