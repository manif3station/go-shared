package shared_lib

import "regexp"

func Grep(list []string, code func(string) bool) []string {
	found := []string{}
	for _, item := range list {
		if code(item) {
			found = append(found, item)
		}
	}
	return found
}

func Map(list []string, code func(string) string) []string {
	found := []string{}
	for _, item := range list {
		found = append(found, code(item))
	}
	return found
}

func Split(r string, str string, n int) []string {
	re := regexp.MustCompile(r)
	return re.Split(str, n)
}

func Replace(r, from, to string) string {
	re := regexp.MustCompile(r)
	return re.ReplaceAllString(from, to)
}

func Match(r, str string) []string {
	re := regexp.MustCompile(r)
	return re.FindAllString(str, -1)
}
