package importer

import "testing"

func TestRequestCardInfo(t *testing.T) {
	testCases := []struct{
		name string
		url string
		filter string
		error error
	}{
		{"functional", "https://api.magicthegathering.io/v1/cards/", "386616", nil},
	}

	for _, tc := range testCases{
		err := RequestCardInfo(tc.url, tc.filter)
		if err != nil {
			t.Logf("Testcase: %v expected error: %v got: %v", tc.name, tc.error, err)
		}
	}
}