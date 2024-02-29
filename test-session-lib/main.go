package main

import "github.com/gorilla/sessions"

func main() {
	_ = sessions.NewFilesystemStore("./session.log", []byte("doantrongdat"))

}
