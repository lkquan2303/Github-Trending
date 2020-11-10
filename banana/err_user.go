package banana

import "errors"

var (
	UserConflict = errors.New("Nguoi dung da ton tai")
	SignUpFails  = errors.New("That bai")
	UserNotFound = errors.New("Khong ton tai email")
)
