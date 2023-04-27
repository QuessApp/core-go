package tests

import "math/rand"

// BatchTest is a struct that contains a function to run before the test and the test itself.
type BatchTest struct {
	OnRun func()
}

// RunBatchTests runs a batch of tests.
// It runs the BeforeRun function before the test and the OnRun function on the test.
func RunBatchTests(t []BatchTest) {
	for _, t := range t {
		t.OnRun()
	}
}

// GenerateRandomString generates a random string with a length of s.
// Special thanks to https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GenerateRandomString(s int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, s)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)

}
