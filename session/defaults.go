package session

var prodDCs = map[int]string{
	1: "149.154.175.53",
	2: "149.154.167.51",
	3: "149.154.175.100",
	4: "149.154.167.91",
	5: "149.154.171.5",
}

var testDCs = map[int]string{
	1: "149.154.175.10",
	2: "149.154.167.40",
	3: "149.154.175.117",
	4: "149.154.167.91",
	5: "149.154.171.5",
}

func defaultServerAddress(dcID int, testMode bool) string {
	if testMode {
		return testDCs[dcID]
	}
	return prodDCs[dcID]
}

func defaultPort(testMode bool) int {
	if testMode {
		return 80
	}
	return 443
}
