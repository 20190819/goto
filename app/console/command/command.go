package command

type command interface {
	Intro() string
	Handle()
}