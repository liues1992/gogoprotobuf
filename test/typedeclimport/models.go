package typedeclimport

import subpkg "github.com/liues1992/gogoprotobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
