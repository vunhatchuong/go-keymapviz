package dactyl_manuform5x6

import _ "embed"

var (
	//go:embed ascii.tmpl
	ascii string
	//go:embed fancy.tmpl
	fancy string

	Layout = map[string]string{
		"ascii": ascii,
		"fancy": fancy,
	}
)
