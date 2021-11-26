package apiserver

import "errors"

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAutheticated          = errors.New("not authenticated")
	errEmailAlreadyExists       = errors.New("email is already registered")
	errSmallPassword            = errors.New("password needs at least 8 simbols")
	errNoApiKey                 = errors.New("tinkoff api key needed")
	errWrongName                = errors.New("no such name in db")
)
