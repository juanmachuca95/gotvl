package pkg

import "errors"

var (
	ErrTranslatorNotFound   = errors.New("cannot found translator")
	ErrLanguageNotSupported = errors.New("language not supported")
	ErrInvalidTVL           = errors.New("cannot get tvl context")
)
