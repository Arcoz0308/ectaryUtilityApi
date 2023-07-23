package error

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrPlayerNotFound    = errors.New("player not found")                                                                                // code => 0
	ErrNoSkin            = errors.New("player don't have a skin")                                                                        // code => 1
	ErrNoCape            = errors.New("player don't have a cape")                                                                        // code => 2
	ErrInvalidSkinSize   = errors.New("invalid size of skin")                                                                            // code => 3
	ErrDecompress        = errors.New("something wrong with decompress of skin")                                                         // code => 4
	ErrDatabase          = errors.New("something wrong with database")                                                                   // code => 5
	ErrPngEncoding       = errors.New("something wrong with encoding of the image")                                                      // code => 6
	ErrInvalidBorderSize = errors.New("invalid size of border, allowed size range : [0-100]")                                            // code => 7
	ErrInvalidImageSize  = errors.New("invalid size of image resolution, allowed value : [0, 8, 16, 32, 64, 128, 256, 512, 1024, 2048]") // code => 8
	ErrUnauthorized      = errors.New("this path require authentication (with header 'Authorization' or parameter '?token=yourtoken?'")  // code => 9
	ErrInvalidToken      = errors.New("your token are invalid")
	ErrServerQuery       = func(adr string, err error) error {
		return errors.New(fmt.Sprintf("someting whrong with query %s, error = %s", adr, err.Error()))
	}
)

const (
	CodeErrPlayerNotFound = iota
	CodeErrNoSkin
	CodeErrNoCape
	CodeErrInvalidSkinSize
	CodeErrDecompress
	CodeErrDatabase
	CodeErrPngEncoding
	CodeErrInvalidBorderSize
	CodeErrInvalidImageSize
	CodeErrUnauthorized
	CodeErrInvalidToken
	CodeErrServerQuery
)

type ApiError struct {
	Code           int
	AdditionalData map[string]interface{}
	Err            error
}

func (a *ApiError) Error() string {
	return a.Err.Error()
}
func (a *ApiError) ErrorWithDebug() string {
	if len(a.AdditionalData) == 0 {
		return a.Err.Error()
	}
	var d []string
	for k, v := range a.AdditionalData {
		d = append(d, fmt.Sprintf("\n %s: %#v", k, v))
	}
	return a.Err.Error() + strings.Join(d, "")
}
