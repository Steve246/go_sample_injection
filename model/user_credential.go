package model

type UserCredential struct {
	Id           uint
	Username     string
	IsBlocked    bool
	UserPassword string
}