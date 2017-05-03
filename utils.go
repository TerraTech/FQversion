// Copyright 2017 FutureQuest, Inc.

package FQversion

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func nvb(name, version, build string) string {
	return fmt.Sprintf("%s:\t%s\t(%s)", name, version, build)
}

func newTabWriter(b *bytes.Buffer) *tabwriter.Writer {
	return tabwriter.NewWriter(b, 0, 0, 1, '.', 0)
}
