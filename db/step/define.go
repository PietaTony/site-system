package step

import "log"

const (
	sqlType      = "mysql"
	sqlIP        = "localhost:3306"
	sqlUser      = "root"
	sqlPassword  = "123qweasd"
	sqlLogin     = sqlUser + ":" + sqlPassword + "@tcp(" + sqlIP + ")/" + sqlTableName
	sqlTableName = "site_system"
)

//time layout / default time
const (
	layout      = "2006-01-02 15:04:05"
	defaultTime = "2020-12-29 12:34:56"
)

//default step data
const (
	defaultID        = 0
	defaultReaderNum = 0
	defaultTagNum    = 0
	defaultUserNum   = 0
	defaultStyle     = 0
	defaultStart     = "2020-12-29 12:34:56"
	defaultEnd       = "2020-12-29 12:34:56"
	defaultIsRemove  = false
)

func panicError(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func printError(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}
