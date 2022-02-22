package businesses

import "errors"

var (
	ErrEmailEmpty = errors.New("Email Empty")
	ErrPassEmpty = errors.New("Password Empty")
	ErrWrongPass = errors.New("Wrong Password")
	ErrEmailExisted = errors.New("Email Existed")
	ErrInternalServer = errors.New("something's gone wrong, contact administrator")
	ErrUsernameEmpty = errors.New("Username Empty")
	ErrPasswordEmpty = errors.New("Password Empty")
	ErrUsernameExisted = errors.New("Username Existed")
	ErrProdNameEmpty = errors.New("Product Name Empty")
	ErrQtyProduct = errors.New("Minimal Product Quantity is 1")
)