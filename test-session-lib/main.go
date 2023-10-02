package main

import "github.com/gorilla/sessions"

func main() {
	store := sessions.NewFilesystemStore("./session.log", []byte("doantrongdat"))
	
}
