// Copyright 2017 FutureQuest, Inc.

package FQversion

import (
	"bytes"
	"fmt"
	"sort"
)

var (
	catalog = []registeredVersion{}
)

type registeredVersion struct {
	Name    string
	Version string
	Build   string
	String  string
}

func Catalog() []registeredVersion {
	if !sort.IsSorted(byName(catalog)) {
		sort.Sort(byName(catalog))
	}

	return catalog
}

func ShowCatalog() string {
	var buf bytes.Buffer
	for rv := range _nvb() {
		buf.WriteString(rv + "\n")
	}

	return buf.String()
}

func ShowCatalogAligned() string {
	var buf bytes.Buffer
	tw := newTabWriter(&buf)

	for rv := range _nvb() {
		fmt.Fprintln(tw, rv)
	}
	tw.Flush()

	return buf.String()
}

func _nvb() <-chan string {
	ch := make(chan string, 10)
	go func() {
		for _, rv := range Catalog() {
			ch <- nvb(rv.Name, rv.Version, rv.Build)
		}
		close(ch)
	}()
	return ch
}

type byName []registeredVersion

func (n byName) Len() int           { return len(n) }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byName) Less(i, j int) bool { return n[i].Name < n[j].Name }
