package errorhelpers

// AddMessages adds messages to an error map
func AddMessages(errMap map[string][]string, context string, msg ...string) {
	if _, ok := errMap[context]; !ok {
		errMap[context] = []string{}
	}
	errMap[context] = append(errMap[context], msg...)
}
