module github.com/sonnes/go-git-fixture/module-a/module-c

go 1.20

replace (
	github.com/sonnes/go-git-fixture => ../../
	github.com/sonnes/go-git-fixture/module-a => ../
)

require (
	github.com/sonnes/go-git-fixture v0.0.0-00010101000000-000000000000
	github.com/sonnes/go-git-fixture/module-a v0.0.0-00010101000000-000000000000
)
