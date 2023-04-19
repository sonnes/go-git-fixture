package modulec

import (
	gogitfixture "github.com/sonnes/go-git-fixture"
	moddulea "github.com/sonnes/go-git-fixture/module-a"
)

// FuncC is a function in module-c
func FuncC() string {
	version := gogitfixture.Version()

	parent := moddulea.FuncA()

	return parent + " / module-c " + version
}
