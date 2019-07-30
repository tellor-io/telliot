package util

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//QueryStruct holds query result fields as a map
type QueryStruct struct {
	Arguments map[string]interface{} `json:"-"`
}

//ParseQueryString extracts the url and any args for extract values from query result payload
func ParseQueryString(_query string) (url string, args []string) {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	url = rgx.FindStringSubmatch(_query)[1]
	result := strings.Split(_query, ")")
	result = strings.Split(result[1], ".")
	// Display all elements.
	for i := range result {
		args = append(args, result[i])
	}

	return url, args
}

//ParsePayload will extract the value from a query payload
func ParsePayload(payload []byte, _granularity uint, args []string) (int, error) {
	var r QueryStruct
	r.Arguments = make(map[string]interface{})
	err := json.Unmarshal(payload, &r.Arguments)
	if err != nil {
		return 0, err
	}

	var res1 interface{}
	for i, arg2 := range args {
		if i > 0 {
			res1 = r.Arguments[arg2]
		}
	}
	s, _ := strconv.ParseFloat(fmt.Sprintf("%v", res1), 64)
	return int(s * float64(_granularity)), nil
}
