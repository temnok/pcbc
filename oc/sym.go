package oc

const (
	symOpen   = '{'
	symClose  = '}'
	symEscape = '\\'
	symSpace  = ' '
)

var (
	sliceOpen   = []byte{symOpen}
	sliceClose  = []byte{symClose}
	sliceEscape = []byte{symEscape}
	sliceSpace  = []byte{symSpace}
)
