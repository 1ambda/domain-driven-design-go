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
	UpdateMessage(message string) string

	IsBadRequestException() bool
	IsUnauthorizedException() bool
	IsForbiddenException() bool
	IsNotFoundException() bool
	IsInternalServerException() bool
}

type Failure struct {
	cause      error
	statusCode int
	timestamp  time.Time
}

func (a *Failure) UpdateMessage(message string) string {
	wrap := errors.Wrap(a.cause, message)
	a.cause = wrap
	return wrap.Error()
}

func (a *Failure) Error() string {
	return a.cause.Error()
}

func (a *Failure) Cause() error {
	return a.cause
}

func (a *Failure) StatusCode() int {
	return a.statusCode
}

func (a *Failure) Timestamp() *time.Time {
	return &a.timestamp
}

func (a *Failure) ToSwaggerError() *dto.Exception {
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

func (a *Failure) IsBadRequestException() bool {
	return a.statusCode == CodeBadRequest
}

func NewBadRequestException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &Failure{
		cause:      wrap,
		statusCode: CodeBadRequest,
		timestamp:  time.Now(),
	}
}

func (a *Failure) IsUnauthorizedException() bool {
	return a.statusCode == CodeUnauthorized
}

func NewUnauthorizedException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &Failure{
		cause:      wrap,
		statusCode: CodeUnauthorized,
		timestamp:  time.Now(),
	}
}

func (a *Failure) IsForbiddenException() bool {
	return a.statusCode == CodeForbidden
}

func NewForbiddenException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &Failure{
		cause:      wrap,
		statusCode: CodeForbidden,
		timestamp:  time.Now(),
	}
}

func (a *Failure) IsNotFoundException() bool {
	return a.statusCode == CodeNotFound
}

func NewNotFoundException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &Failure{
		cause:      wrap,
		statusCode: CodeNotFound,
		timestamp:  time.Now(),
	}
}

func (a *Failure) IsInternalServerException() bool {
	return a.statusCode == CodeInternalServer
}

func NewInternalServerException(err error, message string) Exception {
	wrap := errors.Wrap(err, message)

	return &Failure{
		cause:      wrap,
		statusCode: CodeInternalServer,
		timestamp:  time.Now(),
	}
}
