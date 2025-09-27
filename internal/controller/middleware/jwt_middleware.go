package middleware

import (
	"log/slog"
	"time"

	"github.com/Haya372/go-samples/internal/domain/user"
	"github.com/Haya372/go-samples/internal/domain/vo"
	"github.com/Haya372/go-samples/internal/usecase/login"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	identityKey = "userId"
)

type JwtConfig struct {
	Realm          string
	SecretKey      string
	Timeout        time.Duration
	RefreshTimeout time.Duration
}

func NewGinJwtMiddleware(conf JwtConfig, loginUsecase login.LoginUsecase) *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       conf.Realm,
		Key:         []byte(conf.SecretKey),
		Timeout:     conf.Timeout,
		IdentityKey: identityKey,
		MaxRefresh:  conf.RefreshTimeout,

		TokenHeadName: "Bearer",
		TokenLookup:   "header: Authorization, cookie: jwt",

		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var loginParam struct {
				Id       string `binding:"required" form:"id"       json:"id"`
				Password string `binding:"required" form:"password" json:"password"`
			}
			if err := ctx.ShouldBind(&loginParam); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			user, err := loginUsecase.Execute(ctx, login.LoginParam{
				UserId:      vo.NewUserId(loginParam.Id),
				RawPassword: loginParam.Password,
			})

			if err != nil {
				slog.Error("authentication failed", "error", err)

				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					identityKey: user.Id,
				}
			}

			return jwt.MapClaims{}
		},
	})
	if err != nil {
		panic(err)
	}

	return middleware
}
