package database

type TokenV0 struct {
	id             string
	pkg_name       string
	pkg_path       string
	pkg_address    string
	publisher      string
	token_name     string
	token_symbol   string
	token_decimals float64
	register       bool
}
