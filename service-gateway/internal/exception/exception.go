package exception

import (
	"time"

	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"strings"
	"github.com/pkg/errors"
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
	Error() string
	Wrap(message string) string

	IsBadRequestException() bool
	IsUnauthorizedException() bool
	IsForbiddenException() bool
	IsNotFoundException() bool
	IsInternalServerException() bool
}

type failure struct {
	cause      error
	statusCode int
	timestamp  time.Time
}

func (a *failure) Wrap(message string) string {
	wrap := errors.Wrap(a.cause, message)
	a.cause = wrap
	return wrap.Error()
}

func (a *failure) Error() string {
	return a.cause.Error()
}

func (a *failure) Cause() error {
	return a.cause
}

func (a *failure) StatusCode() int {
	return a.statusCode
}

func (a *failure) Timestamp() *time.Time {
	return &a.timestamp
}

func (a *failure) ToSwaggerError() *dto.Exception {
	errorType := ""

	message := strings.Split(a.cause.Error(), ":")[0]
	// message = errors.Cause(a.cause).Error()

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
		Message:   message,
		Timestamp: a.timestamp.UTC().String(),
		Type:      errorType,
	}
}

func (a *failure) IsBadRequestException() bool {
	return a.statusCode == CodeBadRequest
}

func NewBadRequestException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &failure{
		cause:      wrap,
		statusCode: CodeBadRequest,
		timestamp:  time.Now(),
	}
}

func (a *failure) IsUnauthorizedException() bool {
	return a.statusCode == CodeUnauthorized
}

func NewUnauthorizedException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &failure{
		cause:      wrap,
		statusCode: CodeUnauthorized,
		timestamp:  time.Now(),
	}
}

func (a *failure) IsForbiddenException() bool {
	return a.statusCode == CodeForbidden
}

func NewForbiddenException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &failure{
		cause:      wrap,
		statusCode: CodeForbidden,
		timestamp:  time.Now(),
	}
}

func (a *failure) IsNotFoundException() bool {
	return a.statusCode == CodeNotFound
}

func NewNotFoundException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &failure{
		cause:      wrap,
		statusCode: CodeNotFound,
		timestamp:  time.Now(),
	}
}

func (a *failure) IsInternalServerException() bool {
	return a.statusCode == CodeInternalServer
}

func NewInternalServerException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &failure{
		cause:      wrap,
		statusCode: CodeInternalServer,
		timestamp:  time.Now(),
	}
}
