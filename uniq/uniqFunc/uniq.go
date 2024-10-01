package uniqFunc

import (
	"bufio"
	"strconv"
	"strings"
)

func UniqFunc(input []string, options UniqOptions) []string {
	reader := strings.NewReader(strings.Join(input, "\n"))
	in := bufio.NewScanner(reader)

	var prev string
	var result []string
	count := 1 
	var flag bool

	switch {
	case options.Count:
		for in.Scan() {
			txt := in.Text()

			if txt == prev {
				count++
			} else {
				if flag {
					result = append(result, strconv.Itoa(count) + " " + prev)
				}
				prev = txt
				count = 1
			}
			
			flag = true
		}
		if prev != "" {
			result = append(result, strconv.Itoa(count) + " " + prev)
		}
	case options.Repeated:
		for in.Scan() {
			txt := in.Text()

			if txt == prev {
				count++
			} else {
				if flag  && count > 1{
					result = append(result, prev)
				}
				prev = txt
				count = 1
			}
			
			flag = true
		}
		if flag && count > 1 {
			result = append(result, prev)
		}
	case options.Unique:
		for in.Scan() {
			txt := in.Text()

			if txt == prev {
				count++
			} else {
				if flag  && count == 1{
					result = append(result, prev)
				}
				prev = txt
				count = 1
			}
			
			flag = true
		}
		if flag && count == 1 {
			result = append(result, prev)
		}
	case options.IgnoreCase:
		for in.Scan() {
			txt := in.Text()
	        toLower := strings.ToLower(txt)

			if toLower == prev {
				continue
			}

			prev = toLower
			result = append(result, txt)
		}
	case options.SkipFields != 0:
		for in.Scan() {
			txt := in.Text()
	        fields := strings.Fields(txt)
			if options.SkipFields > len(fields) {
				result = append(result, txt)
				continue
			}

			cutString := strings.Join(fields[options.SkipFields:], " ")
			if cutString == prev {
				continue
			}

			prev = cutString
			result = append(result, txt)
		}
		fallthrough
	case options.SkipChars != 0:
		for in.Scan() {
			txt := in.Text()

			if options.SkipChars > len(txt) {
				result = append(result, txt)
				continue
			}

			cutString := txt[options.SkipChars:]

			if cutString == prev {
				continue
			}

			prev = cutString
			result = append(result, txt)
		}
	default:
		for in.Scan() {
			txt := in.Text()
	
			if txt == prev {
				continue
			}

			prev = txt
			result = append(result, txt)
		}

	}

	return result
}
