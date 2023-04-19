package moduleb

import (
	gogitfixture "github.com/sonnes/go-git-fixture"
	moddulec "github.com/sonnes/go-git-fixture/module-a/module-c"
)

// FuncC is a function in module-c
func FuncC() string {
	version := gogitfixture.Version()

	dependency := moddulec.FuncC()

	return dependency + " / module-b " + version
}
