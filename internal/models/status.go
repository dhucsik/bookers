package models

type ShareRequest struct {
	ID             int    `json:"id"`
	SenderID       int    `json:"sender_id"`
	ReceiverID     int    `json:"receiver_id"`
	SenderBookID   int    `json:"sender_book_id"`
	ReceiverBookID int    `json:"receiver_book_id"`
	SenderStatus   string `json:"sender_status"`
	ReceiverStatus string `json:"receiver_status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type RequestWithFields struct {
	ID             int                  `json:"id"`
	Sender         *UserWithoutPassword `json:"sender"`
	Receiver       *UserWithoutPassword `json:"receiver"`
	SenderBook     *Book                `json:"sender_book"`
	ReceiverBook   *Book                `json:"receiver_book"`
	SenderStatus   string               `json:"sender_status"`
	ReceiverStatus string               `json:"receiver_status"`
	CreatedAt      string               `json:"created_at"`
	UpdatedAt      string               `json:"updated_at"`
}

const (
	StatusCreated           = "created"
	StatusCanceled          = "canceled"
	StatusReceiverRequested = "receiver_requested"
	StatusSenderAccepted    = "sender_accepted"
	StatusSenderProved      = "sender_proved"
	StatusReceiverProved    = "receiver_proved"
)
