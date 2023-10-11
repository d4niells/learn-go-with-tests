package concurrency

import (
	"reflect"
	"testing"
)

// Runs go test -race to check for race condition issues

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds" 
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want  := map[string]bool{
		"http://google.com": 					true,
		"http://blog.gypsydave5.com":	true,
		"waat://furhurterwe.geds": 		false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}