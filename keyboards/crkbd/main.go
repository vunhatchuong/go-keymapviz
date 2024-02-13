package crkbd

import _ "embed"

var (
	//go:embed ascii.tmpl
	ascii string

	Layout = map[string]string{
		"ascii": ascii,
	}
)
