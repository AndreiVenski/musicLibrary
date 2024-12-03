package httpErrors

import "github.com/pkg/errors"

var (
	NotFoundSongOrVerseError     = errors.New("song or verse not found")
	NotFoundSongError            = errors.New("song not found")
	ExistedSongError             = errors.New("song exists in database")
	IncorrectRequestToAPIError   = errors.New("incorrect request")
	APINotWorkError              = errors.New("API server doesn't work")
	UnkownStatusCodeFromAPIError = errors.New("unknown response status code")
)

func IsServiceError(err error) bool {
	switch {
	case errors.Is(err, NotFoundSongError),
		errors.Is(err, NotFoundSongOrVerseError),
		errors.Is(err, ExistedSongError),
		errors.Is(err, IncorrectRequestToAPIError),
		errors.Is(err, APINotWorkError),
		errors.Is(err, UnkownStatusCodeFromAPIError):
		return false

	default:
		return true
	}
}
