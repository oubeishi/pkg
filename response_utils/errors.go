package response_utils

import (
	"github.com/pkg/errors"
)

var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrBadRequest              = New400Response("ErrBadRequest")
	ErrInvalidParent           = New400Response("ErrInvalidParent")
	ErrNotAllowDeleteWithChild = New400Response("ErrNotAllowDeleteWithChild")
	ErrNotAllowDelete          = New400Response("ErrNotAllowDelete")
	ErrInvalidUserName         = New400Response("ErrInvalidUserName")
	ErrInvalidPassword         = New400Response("ErrInvalidPassword")
	ErrInvalidUser             = New400Response("ErrInvalidUser")
	ErrUserDisable             = New400Response("ErrUserDisable")
	ErrUserHasExist            = New400Response("ErrUserHasExist")
	ErrNoPerm                  = NewErrorRes(401, 401, "ErrNoPerm")
	ErrInvalidToken            = NewErrorRes(9999, 401, "ErrInvalidToken")
	ErrNotFound                = NewErrorRes(404, 404, "ErrNotFound")
	ErrMethodNotAllow          = NewErrorRes(405, 405, "ErrMethodNotAllow")
	ErrTooManyRequests         = NewErrorRes(429, 429, "ErrTooManyRequests")
	ErrInternalServer          = NewErrorRes(500, 500, "ErrInternalServer")
)
