package main

import (
	. "github.com/erkanzileli/mockhttpseerver-go"
)

func main() {
	mockServer := Server().Address(":8080")

	mockServer.
		When(
			Request().
				Path("/say-hello").
				Method("GET"),
		).
		Respond(
			Response().
				Status(200).
				Body([]byte("hello friend")),
		)

	if err := mockServer.Start(); err != nil {
		panic(err)
	}
}
