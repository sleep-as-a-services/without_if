package main

import (
	"testing"
)

func TestCaptcha_ShouldDoLeft_Operation(t *testing.T) {
	expected := "1 - one"
	acutal := g(0, 1, 1, 1)
	if expected != acutal {
		t.Errorf("captcha process should return %s but got %s", expected, acutal)
	}
}

func TestCaptcha_ShouldDoRight_Operation(t *testing.T) {
	expected := "one - 1"
	acutal := g(1, 1, 1, 1)
	if expected != acutal {
		t.Errorf("captcha process should return %s but got %s", expected, acutal)
	}
}
