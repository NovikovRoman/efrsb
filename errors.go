package efrsb

import (
	"fmt"
)

type errService struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrUnauthorized не авторизован
type ErrUnauthorized struct{}

func (e ErrUnauthorized) Error() string {
	return "Не авторизован"
}

type ErrNotFound struct{}

func (e ErrNotFound) Error() string {
	return "Не найдено"
}

// ErrTokenExpired истек срок действия токена
type ErrTokenExpired struct{}

func (e ErrTokenExpired) Error() string {
	return "Истек срок действия токена"
}

// ErrRequiredParam не заполнен обязательный параметр запроса
type ErrRequiredParam struct {
	message string
}

func (e ErrRequiredParam) Error() string {
	return e.message
}

// ErrUserNotFound пользователь не найден
type ErrUserNotFound struct {
	message string
}

func (e ErrUserNotFound) Error() string {
	return e.message
}

// ErrAccountExpired срок действия учетной записи истек или не наступил
type ErrAccountExpired struct {
	message string
}

func (e ErrAccountExpired) Error() string {
	return e.message
}

// ErrUserBlocked пользователь заблокирован
type ErrUserBlocked struct {
	message string
}

func (e ErrUserBlocked) Error() string {
	return e.message
}

// ErrParamLimit в параметре limit может быть указано значение меньше или равно 500
type ErrParamLimit struct {
	message string
}

func (e ErrParamLimit) Error() string {
	return e.message
}

func NewErrParamLimit() ErrParamLimit {
	return ErrParamLimit{message: "Превышен лимит запроса - 500"}
}

// DateIntervalErr разница между datePublishEnd и datePublishBegin не может быть больше 31 дня.
type ErrMaxPeriod struct {
	message string
}

func (e ErrMaxPeriod) Error() string {
	return e.message
}

// ErrParamNotFound не найден код параметра.
type ErrParamNotFound struct {
	message string
}

func (e ErrParamNotFound) Error() string {
	return e.message
}

// ErrInvalidParam в параметре указано некорректное значение.
type ErrInvalidParam struct {
	message string
}

func (e ErrInvalidParam) Error() string {
	return e.message
}

func (e errService) Error() error {
	switch e.Code {
	case 1000:
		return ErrRequiredParam{message: e.Message}
	case 1001:
		return ErrParamNotFound{message: e.Message}
	case 1002:
		return ErrInvalidParam{message: e.Message}

	case 1009:
		return ErrParamLimit{message: e.Message}
	case 1010:
		return ErrMaxPeriod{message: e.Message}

	case 2000:
		return ErrUserNotFound{message: e.Message}
	case 2001:
		return ErrAccountExpired{message: e.Message}
	case 2002:
		return ErrUserBlocked{message: e.Message}
	}
	return fmt.Errorf(" %d: %s", e.Code, e.Message)
}
