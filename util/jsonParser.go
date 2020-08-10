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
func ParseQueryString(_query string) (url string, args [][]string) {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	url = rgx.FindStringSubmatch(_query)[1]
	result := strings.Split(_query, ")")
	groups := strings.Split(result[1], ",")
	for _,group := range groups {
		args = append(args, strings.Split(group, ".")[1:])
	}

	return url, args
}

//ParsePayload will extract the value from a query payload
func ParsePayload(payload []byte, args [][]string) ([]float64, error) {
	var f interface{}
	if err := json.Unmarshal(payload, &f); err != nil {
		fmt.Println(err)
	}
	results := make([]float64, len(args))
	for i,argGroup := range args {
		result, err := dumpJSON(f, "root", argGroup, &dumpJsonState{})
		if err != nil {
			//fmt.Println("ERROR", err)
			return nil,errors.New("JSON Parsing error")
		}
		val, _ := strconv.ParseFloat(fmt.Sprintf("%v", result), 64)
		results[i] = val

	}
	return results, nil
}

type dumpJsonState struct {
	Bi int
	Res float64
	Good bool
	FRes float64
	Finished bool
}

func dumpJSON(v interface{}, kn string, args []string, s *dumpJsonState) (float64, error) {

	iterMap := func(x map[string]interface{}, root string, args []string) (val float64, status bool, good bool) {
		defer func() {
			if r := recover(); r != nil {
				//fmt.Println("Problem parsing response", r)
				val = 0
				status = false
				good = false
			}
		}()

		s.Bi++
		var knf string
		if root == "root" {
			knf = "%q:%q"
		} else {
			knf = "%s:%q"
		}
		for k, v := range x {
			if k == args[s.Bi-1] {
				dumpJSON(v, fmt.Sprintf(knf, root, k), args, s)
				if len(args) == s.Bi {
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

		s.Bi++
		if !s.Finished {
			var knf string
			if root == "root" {
				knf = "%q:[%d]"
			} else {
				knf = "%s:[%d]"
			}
			for k, v := range x {
				if s.Bi-1 <= len(args) && len(args) > 0 { //Just added this, we need to check that our numbers are still correct...Still not fixed
					i2, _ := strconv.ParseInt(args[s.Bi-1], 10, 64)
					if k == int(i2) {
						dumpJSON(v, fmt.Sprintf(knf, root, k), args, s)
						if len(args) == s.Bi {
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
		s.Res, s.Finished,s.Good = iterMap(vv, kn, args)
		if !s.Good{
			return 0,errors.New("Itermap Error")
		}
		if s.Finished {
			s.FRes = s.Res
			return s.Res, nil
		}
	case []interface{}:
		s.Res, s.Finished,s.Good = iterSlice(vv, kn, args)
		if !s.Good{
			return 0,errors.New("IterSlice Error")
		}
		if s.Finished {
			s.FRes = s.Res
			return s.Res, nil
		}
	}
	return s.FRes, nil
}

func converter(v interface{}) (float64, error) {
	switch vv := v.(type) {
	case string:
		i, _ := strconv.ParseFloat(normalizeAmerican(vv), 64)
		return i, nil
	case bool:
		return 1, nil
	case float64:
		return vv, nil
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
