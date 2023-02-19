package model

import (
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/pkg/errors"
)

const (
	GENERAL = iota
	GUEST
	ADMIN
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`                      // unique key
	Username string `json:"username" gorm:"unique" binding:"required"` // username
	Password string `json:"password"`                                  // password
	BasePath string `json:"base_path"`                                 // base path
	Role     int    `json:"role"`                                      // user's role
	Disabled bool   `json:"disabled"`
	// Determine permissions by bit
	//   0: can see hidden files
	//   1: can access without password
	//   2: can add aria2 tasks
	//   3: can mkdir and upload
	//   4: can rename
	//   5: can move
	//   6: can copy
	//   7: can remove
	//   8: webdav read
	//   9: webdav write
	//  10: can add qbittorrent tasks
	Permission int32  `json:"permission"`
	OtpSecret  string `json:"-"`
	GithubID   int    `json:"github_id"`
}

func (u User) IsGuest() bool {
	return u.Role == GUEST
}

func (u User) IsAdmin() bool {
	return u.Role == ADMIN
}

func (u User) ValidatePassword(password string) error {
	if password == "" {
		return errors.WithStack(errs.EmptyPassword)
	}
	if u.Password != password {
		return errors.WithStack(errs.WrongPassword)
	}
	return nil
}

func (u User) CanSeeHides() bool {
	return u.IsAdmin() || u.Permission&1 == 1
}