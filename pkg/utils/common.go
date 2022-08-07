package utils

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid2 "github.com/gofrs/uuid"
	"github.com/ztrue/tracerr"
	"os"
	"portfolio-api/graph/model"
	"regexp"
	"strings"
	"time"
)

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}
		return nil
	}
	return
}

func FormatHTML(html model.GetTranslations) *model.GetTranslations {
	liRegex := regexp.MustCompile("<li>")
	ulRegex := regexp.MustCompile("</ul>")
	brRegex := regexp.MustCompile("<br/>")

	var translations []*model.Translation
	for i := 0; i < len(html.Translations); i++ {
		var translateItem = html.Translations[i]
		translateItem.Field = liRegex.ReplaceAllString(translateItem.Field, "\n\t<li>")
		translateItem.Field = ulRegex.ReplaceAllString(translateItem.Field, "\n</ul>")
		translateItem.Field = brRegex.ReplaceAllString(translateItem.Field, "\n\t<br/>")
		translations = append(translations, translateItem)
	}

	return &model.GetTranslations{Field: html.Field, Translations: translations}
}

func ReplaceImage(oldLink, inputImg string) (newLink string, err error) {
	if oldLink != inputImg {
		err = os.Remove(oldLink)
		if err != nil {
			return "", tracerr.Errorf("Не удалось удалить изображение: ", err)
		}
		newLink, err = SaveImage(inputImg)
		if err != nil {
			return "", tracerr.Errorf("Не удалось сохранить изображение ", err)
		}
	} else {
		newLink = inputImg
	}
	return newLink, nil
}

func SaveImage(imgBase64 string) (imgLink string, err error) {
	imgUUID, err := uuid2.NewV1()
	if err != nil {
		return "", err
	}
	path := "uploaded"
	link := path + "/" + imgUUID.String() + ".png"

	b64data := imgBase64[strings.IndexByte(imgBase64, ',')+1:]
	dec, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir("uploaded", 0750)
		if err != nil {
			return "", err
		}
	}

	f, err := os.Create(link)

	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return "", err

	}
	if err := f.Sync(); err != nil {
		return "", err

	}

	return link, nil
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func ParseToken(accessToken string) (int, error) {
	jwtSecret, exists := os.LookupEnv("jwt_secret")
	if !exists {
		return 0, errors.New("")
	}
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("кк")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("ч")
	}
	return claims.UserId, nil
}
