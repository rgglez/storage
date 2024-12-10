package main

import (
	"github.com/kr/pretty"
	"github.com/rgglez/storage/storage"
)

func main() {
	cnn := "oss://test/?credential=hmac:Test123:Test123&endpoint=http://localhost:9090&name=test"
	s := storage.NewStorage(cnn)

	r, err := s.ReadWithSignedURL("2e1cdd85-e594-415e-9b84-70b9718fb15a/2e1cdd85-e594-415e-9b84-70b9718fb15a/facafe9b-1bb7-4c03-bea4-0e72f0b67f3e/9d7710b7-5f2d-4823-bdaf-94b35d454912", 1800000)
	if err != nil {
		panic(err)
	}
	pretty.Print(r)

	/*
	err := s.Read("2e1cdd85-e594-415e-9b84-70b9718fb15a/2e1cdd85-e594-415e-9b84-70b9718fb15a/facafe9b-1bb7-4c03-bea4-0e72f0b67f3e/9d7710b7-5f2d-4823-bdaf-94b35d454912", "./e.zip")
	if err != nil {
		panic(err)
	}*/
}
