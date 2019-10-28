module github.com/winlinvip/mod_replace_app

go 1.13

require (
	github.com/pkg/errors v0.8.1
	github.com/winlinvip/errors v1.0.2
)

replace github.com/pkg/errors v0.8.1 => github.com/winlinvip/errors v1.0.2
