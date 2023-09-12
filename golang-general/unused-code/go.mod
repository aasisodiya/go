module github.com/aasisodiya/go-general/unused-code

go 1.21.1

replace pkg1 => ./pkg1

replace pkg2 => ./pkg2

require (
	pkg1 v0.0.0-00010101000000-000000000000
	pkg2 v0.0.0-00010101000000-000000000000
)
