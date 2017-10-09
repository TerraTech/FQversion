// Copyright 2016-2017 FutureQuest, Inc.
// @ AUTO-GENERATED @

package FQversion

import (
	"fmt"
)

var (
	VERSION string = "v6.1.0"
	BUILD   string = "20171009@02:30:25"
)

func init() {
	Register("FQgolibs", VERSION, BUILD)
}

//func (FQgolibs) Version() string {
//	return Version()
//}

func __Version() string {
	return fmt.Sprintf("FQgolibs:\t%s\t(%s)", VERSION, BUILD)
}
