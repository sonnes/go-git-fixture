package modulea

import gogitfixture "github.com/sonnes/go-git-fixture"

// FuncA is a function in module-a
func FuncA() string {
	version := gogitfixture.Version()

	return "module-a " + version
}
