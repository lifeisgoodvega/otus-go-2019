package top10

import (
	"regexp"
	"sort"
	"strings"
)

func filter(arr []string, predicate func(string) bool) (result []string) {
	for _, s := range arr {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return
}

func transform(arr []string, tFun func(string) string) (result []string) {
	for _, s := range arr {
		result = append(result, tFun(s))
	}
	return
}

type fPair struct {
	Word  string
	Count int
}

//Top10 - returns top 10 most frequent words in inputString
func Top10(inputString string) (result []string) {
	r := regexp.MustCompile("[^\\s:,.]+")
	unfilteredArr := r.FindAllString(inputString, -1)
	arr := transform(
		filter(unfilteredArr, func(s string) bool { return s != "-" }),
		func(s string) string { return strings.ToLower(s) })

	fMap := make(map[string]int)
	for _, s := range arr {
		fMap[s]++
	}

	var fpairArr []fPair
	for k, v := range fMap {
		fpairArr = append(fpairArr, fPair{k, v})
	}

	sort.SliceStable(fpairArr, func(i, j int) bool {
		if fpairArr[i].Count == fpairArr[j].Count {
			return fpairArr[i].Word < fpairArr[j].Word
		} 
		return fpairArr[i].Count > fpairArr[j].Count
	})

	for i, v := range fpairArr {
		if (i == 10) { break; }
		result = append(result, v.Word)
	}
	return
}
