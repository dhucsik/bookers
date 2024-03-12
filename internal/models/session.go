package models

import "context"

const (
	Key     = "session"
	roleKey = "role"
)

type Session struct {
	UserID int
	Role   string
	Claims map[string]interface{}
}

func (s *Session) Exp() int64 {
	return s.Claims["exp"].(int64)
}

func (s *Session) SetInCtx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, Key, s)
	ctx = context.WithValue(ctx, roleKey, s.Role)

	return ctx
}

func GetSession(ctx context.Context) (*Session, bool) {
	session, ok := ctx.Value(Key).(*Session)
	if !ok || session == nil {
		return nil, false
	}

	return session, ok
}
