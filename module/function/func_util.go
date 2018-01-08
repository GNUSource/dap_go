package function

import (
	"module/logger"
	"regexp"
	"strings"
)

// delete subStr from str
// deleStr format is : 说明书|标盒|...
func DeleteSubStr(str string, deleteStr string) string {
	// replace （...）to (...)
	str = strings.Replace(str, "（", "(", -1)
	str = strings.Replace(str, "）", ")", -1)

	delStr := strings.Split(deleteStr, "|")
	for _, value := range delStr {
		str = strings.Replace(str, value, "", -1)
	}
	return str
}

// drug subStr from str
func DrawStr(str string, regexpRule string) (string, []string) {
	// replace （...）to (...)
	str = strings.Replace(str, "（", "(", -1)
	str = strings.Replace(str, "）", ")", -1)
	logger.Trace.Printf("origin str is %v\n", str)

	// regexp \(\S+\)|(\d+(.\d+)?\S+(\*|/)+(\d+(.\d+)?)?(mg|g|ml|mm|粒|盒|袋|片|板)?)|(\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板))
	reg := regexp.MustCompile(regexpRule)
	regexpStrs := reg.FindAllString(str, -1)
	for _, v := range regexpStrs {
		str = DeleteSubStr(str, v)
	}
	logger.Trace.Printf("drawed str is %v\n", str)
	logger.Trace.Printf("matched str is %#v\n", regexpStrs)
	return str, regexpStrs
}
