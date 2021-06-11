package cgo

import "entryTask/main/session"

var globalSession *session.Manager

func init() {
	globalSession, _ = session.NewManager("memory", "GSESSIONID", 3600)
}

func GlobalSession() *session.Manager {
	return globalSession
}
