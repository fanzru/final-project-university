package errs

import (
	"errors"
)

var (
	// App Error
	ErrNotValidToken    = errors.New("not a valid token")
	ErrNotValidUserID   = errors.New("not a valid user id")
	ErrUserVerified     = errors.New("user is already verified")
	ErrEmailUsed        = errors.New("email is already in used")
	ErrPasswordNotMatch = errors.New("wrong password")

	// Mailer Error
	ErrEmailNotSend = errors.New("email not send")

	// Usecase Error
	ErrInvalidParam = errors.New("invalid parameter")

	// Repository Error
	ErrInstanceNotFound    = errors.New("user not found")
	ErrFailSaveUserEmail   = errors.New("cannot save verification code")
	ErrFailSaveUserAccount = errors.New("cannot save user account token code")

	// Model error
	ErrEncryptPassword = errors.New("error creating encrypted passowrd")
	ErrCreateToken     = errors.New("cannot create random token")

	// Json Unmarshal Error
	ErrJsonUnmarshal = errors.New("json unmarshal error")

	// Limiter
	ErrManyWrongPassword = errors.New("your account has too many wrong passwords, please try again in 5 minutes")
)
