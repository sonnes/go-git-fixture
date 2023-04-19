module github.com/sonnes/go-git-fixture/module-b

go 1.20

replace (
	github.com/sonnes/go-git-fixture => ../
	github.com/sonnes/go-git-fixture/module-a => ../module-a
	github.com/sonnes/go-git-fixture/module-a/module-c => ../module-a/module-c
)

require (
	github.com/sonnes/go-git-fixture v0.0.0-00010101000000-000000000000
	github.com/sonnes/go-git-fixture/module-a/module-c v0.0.0-00010101000000-000000000000
)

require github.com/sonnes/go-git-fixture/module-a v0.0.0-00010101000000-000000000000 // indirect
