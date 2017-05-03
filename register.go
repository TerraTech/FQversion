// Copyright 2017 FutureQuest, Inc.

package FQversion

func Register(name, version, build string) {
	sv := nvb(name, version, build)
	catalog = append(catalog, registeredVersion{Name: name, Version: version, Build: build, String: sv})
}
