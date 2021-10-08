package http

import (
	"ecommerce-integrations/driver"
	"ecommerce-integrations/internal/app/custom_types"
	"ecommerce-integrations/internal/app/woocommerce/modules/auth/repository"
	"ecommerce-integrations/internal/app/woocommerce/modules/auth/usecase"
	"ecommerce-integrations/internal/configs"
	"ecommerce-integrations/internal/utils/app"
	"ecommerce-integrations/internal/utils/constants"
	"ecommerce-integrations/internal/utils/logger"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(db *driver.Database) AuthHandler {
	tRepo := repository.NewAuthRepository(db)
	pUseCase := usecase.NewAuthUseCase(tRepo)
	return AuthHandler{
		authUseCase: pUseCase,
	}
}

func (instance *AuthHandler) RequestAuthorizeApp(c *gin.Context) {
	app := app.Gin{C: c}
	config := configs.AppConfig

	storeUrl := "http://localhost:8080"
	url := fmt.Sprintf("%s%s", storeUrl, constants.WOOCOMMERCE_APP_AUTHORIZE_ENDPOINT)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Logger.Error(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("app_name", constants.WOOCOMMERCE_APP_NAME)
	q.Add("scope", constants.WOOCOMMERCE_APP_SCOPE)
	q.Add("user_id", "123")
	returnUrl := fmt.Sprintf("%s://%s", config.HTTP.Protocol, config.HTTP.Host)
	q.Add("return_url", returnUrl)
	callbackUrl := fmt.Sprintf("%s://%s/auth/callback", config.HTTP.Protocol, config.HTTP.Host)
	q.Add("callback_url", callbackUrl)
	req.URL.RawQuery = q.Encode()

	logger.Logger.Info(req.URL.String())

	// app.Ok(gin.H{
	// 	"message": "Ok",
	// })
	app.C.Redirect(http.StatusMovedPermanently, req.URL.String())
}

func (instance *AuthHandler) Callback(c *gin.Context) {
	app := app.Gin{C: c}

	bodyData := custom_types.AuthCallBackBody{}

	if err := c.ShouldBind(&bodyData); err != nil {
		logger.Logger.Error(err)
		app.BadRequest(gin.H{
			"message": err.Error(),
		})
		return
	}

	// body, _ := ioutil.ReadAll(c.Request.Body)
	logger.Logger.Infof("%v", bodyData)

	// TODO: Store token into database

	// logger.Logger.Info("================")
	// logger.Logger.Info(app.C.Request.Body)

	app.Ok(gin.H{
		"message": "Ok Callback",
	})
}
