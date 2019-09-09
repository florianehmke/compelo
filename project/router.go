package project

import (
	"compelo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	// Key identifies the project inside the gin.Context
	Key = "project"

	idKey   = "projectId"
	nameKey = "projectName"
)

type Router struct {
	s   *Service
	jwt *jwt.GinJWTMiddleware
}

func NewRouter(s *Service) *Router {
	return &Router{s, createMiddleware(s)}
}

func (r *Router) CreateProject(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.Bind(&body)

	var p *compelo.Project
	if err == nil {
		p, err = r.s.CreateProject(body.Name, hashAndSalt([]byte(body.Password)))
	}

	if err == nil {
		c.JSON(http.StatusCreated, p)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) SelectProject(c *gin.Context) {
	r.jwt.LoginHandler(c)
}

func (r *Router) Middleware() gin.HandlerFunc {
	return r.jwt.MiddlewareFunc()
}

func (r *Router) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, r.s.LoadProjects())
}

func createMiddleware(s *Service) *jwt.GinJWTMiddleware {
	mw, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "compelo",
		Key:         []byte("secret key"), // FIXME export
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: idKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if p, ok := data.(*compelo.Project); ok {
				return jwt.MapClaims{
					idKey:   p.ID,
					nameKey: p.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &compelo.Project{
				Name: claims[nameKey].(string),
				Model: compelo.Model{
					ID: uint(claims[idKey].(float64)),
				},
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var body struct {
				Name     string `json:"name" binding:"required"`
				Password string `json:"password" binding:"required"`
			}
			if err := c.ShouldBind(&body); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			p, err := s.LoadByName(body.Name)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			err = bcrypt.CompareHashAndPassword(p.PasswordHash, []byte(body.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &p, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if fromToken, ok := data.(*compelo.Project); ok {
				fromDB, err := s.LoadByName(fromToken.Name)
				if err != nil {
					return false
				}
				c.Set(Key, fromDB)
				return true
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

func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}
