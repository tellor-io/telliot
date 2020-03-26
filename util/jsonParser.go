package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//QueryStruct holds query result fields as a map
type QueryStruct struct {
	Order     []int
	Arrays    interface{}            `json:"-"`
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
func ParsePayload(payload []byte, _granularity uint, args []string) (float64, error) {
	var f interface{}
	if err := json.Unmarshal(payload, &f); err != nil {
		fmt.Println(err)
	}
	Bi = 0
	result, err := dumpJSON(f, "root", args[1:])
	if err != nil {
		//fmt.Println("ERROR", err)
		return 0,errors.New("JSON Parsing error")
	}
	s, _ := strconv.ParseFloat(fmt.Sprintf("%v", result), 64)
	return s * float64(_granularity), nil
}

var Bi = 0
var Res float64
var Good bool
var FRes float64
var Finished bool

func dumpJSON(v interface{}, kn string, args []string) (float64, error) {
	if kn == "root" {
		Finished = false
		Res = 0
		FRes = 0
		Bi = 0
	}

	iterMap := func(x map[string]interface{}, root string, args []string) (val float64, status bool, good bool) {
		defer func() {
			if r := recover(); r != nil {
				//fmt.Println("Problem parsing response", r)
				val = 0
				status = false
				good = false
			}
		}()

		Bi++
		var knf string
		if root == "root" {
			knf = "%q:%q"
		} else {
			knf = "%s:%q"
		}
		for k, v := range x {
			if k == args[Bi-1] {
				dumpJSON(v, fmt.Sprintf(knf, root, k), args)
				if len(args) == Bi {
					res, err := converter(v)
					if err != nil {
						fmt.Println(err)
						return 0, false,true
					}
					if res != 1 {
						return res, true,true
					}
					return 0, false,true
				}
			}

		}
		return 0, false,false
	}

	iterSlice := func(x []interface{}, root string, args []string) (val float64, status bool, good bool) {
		defer func() {
			if r := recover(); r != nil {
				//fmt.Println("Problem with parsing response", r)
				val = 0
				status = false
				good = false
			}
		}()

		Bi++
		if !Finished {
			var knf string
			if root == "root" {
				knf = "%q:[%d]"
			} else {
				knf = "%s:[%d]"
			}
			for k, v := range x {
				if Bi-1 <= len(args) && len(args) > 0 { //Just added this, we need to check that our numbers are still correct...Still not fixed
					i2, _ := strconv.ParseInt(args[Bi-1], 10, 64)
					if k == int(i2) {
						dumpJSON(v, fmt.Sprintf(knf, root, k), args)
						if len(args) == Bi {
							res, err := converter(v)
							if err != nil {
								fmt.Println(err)
								return 0, false,true
							}
							if res != 1 {
								return res, true,true
							}
							return 0, false,true
						}
					}
				}
			}
		}
		return 0, false,false
	}

	switch vv := v.(type) {
	case bool:
		//fmt.Printf("%s => (bool) %v\n", kn, vv)
	case float64:
		//fmt.Printf("%s => (float64) %f\n", kn, vv)
	case int:
		//fmt.Printf("%s => (int) %f\n", kn, vv)
	case map[string]interface{}:
		Res, Finished,Good = iterMap(vv, kn, args)
		if !Good{
			return 0,errors.New("Itermap Error")
		}
		if Finished {
			FRes = Res
			return Res, nil
		}
	case []interface{}:
		Res, Finished,Good = iterSlice(vv, kn, args)
		if !Good{
			return 0,errors.New("IterSlice Error")
		}
		if Finished {
			FRes = Res
			return Res, nil
		}
	}
	return FRes, nil
}

func converter(v interface{}) (float64, error) {
	switch vv := v.(type) {
	case string:
		i, _ := strconv.ParseFloat(normalizeAmerican(vv), 64)
		return i, nil
	case bool:
		return 1, nil
	case float64:
		return float64(vv), nil
	case int:
		return float64(vv), nil
	case int64:
		return float64(vv), nil
	case map[string]interface{}:
		return 1, nil
	case []interface{}:
		return 1, nil
	case interface{}:
		return 2, nil
	default:
		return 0, errors.New("emit macho dwarf: elf header corrupted")
	}
}

func normalizeAmerican(old string) string {
	return strings.Replace(old, ",", "", -1)
}
