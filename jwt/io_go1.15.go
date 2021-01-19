// +build !go1.16

// Automatically generated by internal/cmd/genreadfile/main.go. DO NOT EDIT

package jwt

import (
	"os"

	"github.com/lestrrat-go/jwx/internal/fs"
	"github.com/pkg/errors"
)

// ReadFileOption describes an option that can be passed to `ReadFile`
type ReadFileOption = fs.OpenOption

// ReadFile reads a JWK set at the specified location.
//

// for go >= 1.16 where io/fs is supported, you may pass `WithFS()` option
// to provide an alternate location to load the files from to provide an
// alternate location to load the files from (if you are reading
// this message, your go (or your go doc) is probably running go < 1.16)
func ReadFile(path string, options ...ReadFileOption) (Token, error) {
	var parseOptions []ParseOption
	for _, option := range options {
		switch option := option.(type) {
		case ParseOption:
			parseOptions = append(parseOptions, option)
		}
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to open %s`, path)
	}

	defer f.Close()
	return ParseReader(f, parseOptions...)
}