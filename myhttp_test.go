package main

import (
	"os"
	"testing"
)

func TestFixUrl(t *testing.T) {
	got := fixUrl("http://www.google.es")
	if got != "http://www.google.es" {
		t.Errorf("fixUrl(\"http://www.google.es\") = %s; want \"http://www.google.es\")", got)
	}

	got = fixUrl("ali")
	if got != "http://ali" {
		t.Errorf("fixUrl(\"ali\") = %s; want \"http://ali\")", got)
	}

	got = fixUrl("www.twitter.com")
	if got != "http://www.twitter.com" {
		t.Errorf("fixUrl(\"www.twitter.com\") = %s; want \"http://www.twitter.com\")", got)
	}

	got = fixUrl("adjust.com")
	if got != "http://adjust.com" {
		t.Errorf("fixUrl(\"adjust.com\") = %s; want \"http://adjust.com\")", got)
	}

	//It's impossible to construct a proper url from this
	got = fixUrl("httpsss://www.google.com")
	if got != "httpsss://www.google.com" {
		t.Errorf("fixUrl(\"httpsss://www.google.com\") = %s; want \"httpsss://www.google.com\")", got)
	}
}

func ExampleRequest() {
	sem := make(chan int, 1)
	sem <- 1
	request("http://ali", sem)
	// Output:
	// Error requesting: 'http://ali'.
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
