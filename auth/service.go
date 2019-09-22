package auth

import (
	"log"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"compelo/project"
)

const (
	projectIDKey   = "projectId"
	projectNameKey = "projectName"
)

type Config struct {
	Secret        string
	Realm         string
	TimeoutSec    int
	MaxRefreshSec int
}

func DefaultConfig() Config {
	return Config{
		Realm:         "compelo",
		TimeoutSec:    60 * 60,
		MaxRefreshSec: 60 * 60,
	}
}

func (c Config) WithSecret(secret string) Config {
	c.Secret = secret
	return c
}

type Service struct {
	config Config
	ps     *project.Service
}

func NewService(config Config, ps *project.Service) *Service {
	return &Service{
		config: config,
		ps:     ps,
	}
}

type TokenData struct {
	ProjectID   uint   `json:"projectId"`
	ProjectName string `json:"projectName"`
}

type AuthRequest struct {
	ProjectName string `json:"projectName"`
	Password    string `json:"password" binding:"required"`
}

func (s *Service) Middleware() *jwt.GinJWTMiddleware {
	mw, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       s.config.Realm,
		Key:         []byte(s.config.Secret),
		Timeout:     time.Second * time.Duration(s.config.TimeoutSec),
		MaxRefresh:  time.Second * time.Duration(s.config.MaxRefreshSec),
		IdentityKey: projectIDKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if d, ok := data.(TokenData); ok {
				return jwt.MapClaims{
					projectIDKey:   d.ProjectID,
					projectNameKey: d.ProjectName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &TokenData{
				ProjectID:   uint(claims[projectIDKey].(float64)),
				ProjectName: claims[projectNameKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var param AuthRequest
			if err := c.ShouldBind(&param); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			// If project name is set, auth for project.
			if param.ProjectName != "" && param.Password != "" {
				p, err := s.ps.AuthorizeProject(param.ProjectName, param.Password)
				if err != nil {
					return nil, jwt.ErrFailedAuthentication
				}
				return TokenData{ProjectName: p.Name, ProjectID: p.ID}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if data, ok := data.(*TokenData); ok {
				// Just load the project into the gin context.
				// If we have a valid token no further authorization is necessary.
				if p, err := s.ps.LoadByName(data.ProjectName); err == nil {
					c.Set(project.Key, p)
					return true
				}
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal(err)
	}
	return mw
}
