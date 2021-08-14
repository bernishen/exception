package test

import (
	"exception"
	"log"
	"testing"
)

func TestEnum(t *testing.T) {
	log.Print(exception.Information)
	log.Print(exception.Warning)
	log.Print(exception.Error)
	log.Print(exception.FatalError)
}