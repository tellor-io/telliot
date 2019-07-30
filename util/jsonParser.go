package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func parseQueryString(_query string) (url string, args []string) {
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

type QueryStruct struct {
	Arguments map[string]interface{} `json:"-"`
}

func fetchAPI(_granularity uint, queryString string) int {

	var r QueryStruct
	r.Arguments = make(map[string]interface{})

	url, args := parseQueryString(queryString)
	resp, _ := http.Get(url)

	input, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(input, &r.Arguments)
	if err != nil {
		panic(err)
	}
	var res1 interface{}
	for i, arg2 := range args {
		if i > 0 {
			res1 = r.Arguments[arg2]
		}
	}
	s, _ := strconv.ParseFloat(fmt.Sprintf("%v", res1), 64)
	return int(s * float64(_granularity))
}
