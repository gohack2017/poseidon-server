package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

var (
	ErrInvalidID       = errors.New("Invalid bson object id.")
	ErrInvalidArgs     = errors.New("Invalid arguments.")
	ErrInvalidPassword = errors.New("Password is incorrect.")
	ErrNotPersisted    = errors.New("Record has not persisted.")
	ErrDupScene        = errors.New("Duplicated scene event.")
	ErrNotFound        = mgo.ErrNotFound
)
