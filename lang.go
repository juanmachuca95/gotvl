package pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

type Lang struct {
	Tag  language.Tag
	Lang string
	Path string
}

func NewLang(tag language.Tag) Lang {
	path := fmt.Sprintf("active.%s.toml", strings.Split(tag.String(), "-")[0])
	return Lang{
		Tag:  tag,
		Lang: tag.String(),
		Path: path,
	}
}

const (
	EN = "en"
	ES = "es"
)

var (
	UniversalTranslator *ut.UniversalTranslator
)

// Este es una traductor por defecto para validaciones de gin
func (l Lang) NewLanguageTranslator() (ut.Translator, *validator.Validate, error) {
	// Note: que el primer parametro es la alternativa cuando no se encuentra el traductor buscado
	// aún asi es necesario colorcarlo en después
	UniversalTranslator = ut.New(en.New(), en.New(), es.New())
	v := validator.New()

	switch l.Lang {
	case EN:
		t, ok := UniversalTranslator.GetTranslator(EN) // en
		if !ok {
			return nil, v, ErrTranslatorNotFound
		}
		en_translations.RegisterDefaultTranslations(v, t)
		return t, v, nil
	case ES:
		t, ok := UniversalTranslator.GetTranslator(ES) // es
		if !ok {
			return nil, v, ErrTranslatorNotFound
		}
		es_translations.RegisterDefaultTranslations(v, t)
		return t, v, nil
	default:
		return nil, v, ErrLanguageNotSupported
	}
}

// Esto puede ser un archivo .toml o json. Para este ejemplo usaremos toml
func (l Lang) NewCustomTranslatorI18n() (*i18n.Localizer, error) {
	bundle := i18n.NewBundle(l.Tag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// Acá podriamos cargar todos los .toml pero me parece mejor utilizar solo el que se es requerido
	_, err := bundle.LoadMessageFile(l.Path)
	if err != nil {
		return nil, err
	}

	return i18n.NewLocalizer(bundle, l.Lang), nil
}

func SetInstancesTranslate(c *gin.Context) {
	acceptLanguage := c.GetHeader("Accept-Language")
	lang, err := language.Parse(acceptLanguage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newLang := NewLang(lang)
	l, err := newLang.NewCustomTranslatorI18n()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t, v, err := newLang.NewLanguageTranslator()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tvl := NewTVLContext(t, v, l)
	c.Set("tvl", tvl)
	c.Next()
}
