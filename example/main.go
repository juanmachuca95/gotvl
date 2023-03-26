package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	gotvl "github.com/juanmachuca95/gotvl"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,lte=8"`
}

func main() {
	r := gin.Default()

	// Accept-Language (en or es) required
	r.Use(gotvl.SetInstancesTranslate)
	r.POST("user", func(ctx *gin.Context) {
		tvl, err := gotvl.GetTVLContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		var u UserRequest
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, tvl.L.MustLocalize(&i18n.LocalizeConfig{MessageID: "Request malformed"}))
			return
		}

		if err := tvl.V.Struct(u); err != nil {
			if err, ok := err.(validator.ValidationErrors); ok {
				ctx.JSON(http.StatusBadRequest, err.Translate(tvl.T))
				return
			}

			ctx.JSON(http.StatusUnprocessableEntity, tvl.L.MustLocalize(&i18n.LocalizeConfig{MessageID: "Error unexpected"}))
			return
		}

		ctx.JSON(http.StatusOK, tvl.L.MustLocalize(&i18n.LocalizeConfig{MessageID: "OK", PluralCount: 1})) // 🚀
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
