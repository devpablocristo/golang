package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

var jokes = []Joke{
	{ID: 1, Likes: 0, Joke: "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	{ID: 2, Likes: 0, Joke: "What do you call a fake noodle? An Impasta."},
	{ID: 3, Likes: 0, Joke: "How many apples grow on a tree? All of them."},
	{ID: 4, Likes: 0, Joke: "Want to hear a joke about paper? Nevermind it's tearable."},
	{ID: 5, Likes: 0, Joke: "I just watched a program about beavers. It was the best dam program I've ever seen."},
	{ID: 6, Likes: 0, Joke: "Why did the coffee file a police report? It got mugged."},
	{ID: 7, Likes: 0, Joke: "How does a penguin build its house? Igloos it together."},
	{ID: 8, Likes: 0, Joke: "Dad, did you get a haircut? No I got them all cut."},
	{ID: 9, Likes: 0, Joke: "What do you call a Mexican who has lost his car? Carlos."},
	{ID: 10, Likes: 0, Joke: "Dad, can you put my shoes on? No, I don't think they'll fit me."},
	{ID: 11, Likes: 0, Joke: "Why did the scarecrow win an award? Because he was outstanding in his field."},
	{ID: 12, Likes: 0, Joke: "Why don't skeletons ever go trick or treating? Because they have no body to go with."},
}

var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			aud := os.Getenv("AUTH0_API_AUDIENCE")
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAudience {
				return nil, errors.New("invalid audience")
			}
			iss := os.Getenv("AUTH0_DOMAIN")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return nil, errors.New("invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				return nil, fmt.Errorf("could not get cert: %v", err)
			}

			return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	jwtMiddleWare = jwtMiddleware

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", authMiddleware(), JokeHandler)
		api.POST("/jokes/like/:jokeID", authMiddleware(), LikeJoke)
	}

	log.Fatal(router.Run(":3000"))
}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks Jwks
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return cert, err
	}

	for _, key := range jwks.Keys {
		if token.Header["kid"] == key.Kid {
			return fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", key.X5c[0]), nil
		}
	}

	return cert, errors.New("unable to find appropriate key")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtMiddleWare.CheckJWT(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, Response{Message: "Unauthorized"})
			return
		}
		c.Next()
	}
}

func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

func LikeJoke(c *gin.Context) {
	jokeID, err := strconv.Atoi(c.Param("jokeID"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range jokes {
		if jokes[i].ID == jokeID {
			jokes[i].Likes++
			c.JSON(http.StatusOK, jokes[i])
			return
		}
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func getJokeByID(id int) (*Joke, error) {
	for _, joke := range jokes {
		if joke.ID == id {
			return &joke, nil
		}
	}
	return nil, errors.New("joke not found")
}
