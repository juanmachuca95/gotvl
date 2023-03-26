package pkg

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TVLContext struct {
	T ut.Translator
	V *validator.Validate
	L *i18n.Localizer
}

func NewTVLContext(
	t ut.Translator,
	v *validator.Validate,
	l *i18n.Localizer,
) TVLContext {
	return TVLContext{T: t, V: v, L: l}
}

func GetTVLContext(c *gin.Context) (TVLContext, error) {
	tvl, ok := c.MustGet("tvl").(TVLContext)
	if !ok {
		return TVLContext{}, ErrInvalidTVL
	}

	return tvl, nil
}
