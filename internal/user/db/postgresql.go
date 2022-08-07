package user

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/ztrue/tracerr"
	"os"
	"portfolio-api/graph/model"
	"portfolio-api/infrastructure/postgresql"
	"portfolio-api/internal/user"
	"portfolio-api/pkg/utils"
	"time"
)

type repository struct {
	client postgres.Client
}

func (r *repository) Auth(ctx context.Context, input model.UserInput) (model.UserOutput, error) {
	if input.Login == "" {
		return nil, tracerr.Errorf("Логин не может быть пустым")
	}
	if input.Password == "" {
		return nil, tracerr.Errorf("Пароль не может быть пустым")
	}

	q := `

		SELECT 
			id, login, password

		FROM 
			public.user

		WHERE 
			login = $1

		`

	var usr model.User
	usrRow := r.client.QueryRow(ctx, q, input.Login)

	err := usrRow.Scan(&usr.ID, &usr.Login, &usr.Password)
	if err != nil {
		return model.NotFoundError{Message: "Пользователь не найден"}, nil
	}

	if usr.Login == input.Login && usr.Password != input.Password {
		return model.WrongPassword{Message: "Неверный пароль"}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.TokenClaims{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		Subject:   usr.Login,
	}, UserId: usr.ID})

	jwtSecret, exists := os.LookupEnv("jwt_secret")

	if exists {
		signingString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			return nil, err
		}
		usr.AccessToken = signingString
	}

	return usr, nil
}
func (r *repository) GetOne(ctx context.Context, id int) (model.User, error) {
	qUser := `

				SELECT 
					id, login, role
				
				FROM 
					public.user 
				
				WHERE 
					id = $1

			 `

	var usr model.User
	err := r.client.QueryRow(ctx, qUser, id).Scan(&usr.ID, &usr.Login, &usr.RoleID)
	if err != nil {
		return model.User{}, tracerr.Errorf("Не найти пользователя: %s", err)
	}

	qRole := `

				SELECT 
					 name
				
				FROM 
					public.role 
				
				WHERE 
					id = $1

			 `

	err = r.client.QueryRow(ctx, qRole, &usr.RoleID).Scan(&usr.Role)
	if err != nil {
		return model.User{}, tracerr.Errorf("Не найти роль пользователя: %s", err)
	}
	return usr, nil
}
func NewRepository(client postgres.Client) user.Repository {
	return &repository{
		client: client,
	}
}
