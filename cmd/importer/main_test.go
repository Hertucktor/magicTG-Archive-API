package main

import "testing"

func TestImportCardInfo(t *testing.T) {
	/*responseCodes := []struct{
		name string
		code int
		codeDescription string
	}{
		{"Bad Request", 400, "We could not process that action"},
		{"Forbidden", 403, "You exceeded the rate limit"},
		{"Not Found", 404, "The requested resource could not be found"},
		{"Internal Server Error", 500, "We had a problem with our server. Please try again later"},
		{"Service Unavailable", 503, "We are temporarily offline for maintenance. Please try again later"},

	}*/
	codes := []int{400, 403,404, 500, 503}
	descriptions := []string{"We could not process that action","You exceeded the rate limit","The requested resource could not be found","We had a problem with our server. Please try again later", "We are temporarily offline for maintenance. Please try again later"}

	response := ImportCardInfo
	if !containsIntInSliceInts(codes, response.Code){
		t.Fatalf("Expected one of those http codes: %v, got: %v")
	}
	if !containsStringInSliceStrings(descriptions, response.Description){
		t.Fatalf("Expected one of those descriptions: %v\n, got: %v")
	}
}

func containsIntInSliceInts(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsStringInSliceStrings(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}