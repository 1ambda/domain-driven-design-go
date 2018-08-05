package exception

import (
	"time"

	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
)

const CodeBadRequest = 400
const CodeUnauthorized = 401
const CodeForbidden = 403
const CodeNotFound = 404
const CodeInternalServer = 500

type Exception interface {
	Cause() error
	StatusCode() int
	Timestamp() *time.Time
	ToSwaggerError() *dto.Exception

	IsBadRequestException() bool
	IsUnauthorizedException() bool
	IsForbiddenException() bool
	IsNotFoundException() bool
	IsInternalServerException() bool
}

type appException struct {
	cause      error
	statusCode int
	timestamp  time.Time
}

func (a *appException) Cause() error {
	return a.cause
}

func (a *appException) StatusCode() int {
	return a.statusCode
}

func (a *appException) Timestamp() *time.Time {
	return &a.timestamp
}

func (a *appException) ToSwaggerError() *dto.Exception {
	errorType := ""

	switch status := a.statusCode; status {
	case CodeBadRequest:
		errorType = dto.ExceptionTypeBadRequest

	case CodeForbidden:
		errorType = dto.ExceptionTypeForbidden

	case CodeUnauthorized:
		errorType = dto.ExceptionTypeUnauthorized

	case CodeNotFound:
		errorType = dto.ExceptionTypeNotFound

	default:
		errorType = dto.ExceptionTypeInternalServer
	}

	return &dto.Exception{
		Code:      int64(a.statusCode),
		Message:   a.cause.Error(),
		Timestamp: a.timestamp.UTC().String(),
		Type:      errorType,
	}
}

func (a *appException) IsBadRequestException() bool {
	return a.statusCode == CodeBadRequest
}

func NewBadRequestException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeBadRequest,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsUnauthorizedException() bool {
	return a.statusCode == CodeUnauthorized
}

func NewUnauthorizedException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeUnauthorized,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsForbiddenException() bool {
	return a.statusCode == CodeForbidden
}

func NewForbiddenException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeForbidden,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsNotFoundException() bool {
	return a.statusCode == CodeNotFound
}

func NewNotFoundException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeNotFound,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsInternalServerException() bool {
	return a.statusCode == CodeInternalServer
}

func NewInternalServerException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeInternalServer,
		timestamp:  time.Now(),
	}
}
