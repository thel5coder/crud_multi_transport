package usecase

import (
	"crud_multi_transport/db/repositories/actions"
	"crud_multi_transport/helpers/hashing"
	"crud_multi_transport/helpers/messages"
	"crud_multi_transport/http_server/request"
	"crud_multi_transport/usecase/viewmodel"
	"errors"
	uuid "github.com/satori/go.uuid"
	"os"
)

type AuthUseCase struct {
	*UcContract
}

func (uc AuthUseCase) UpdateSessionLogin(ID string) (res string, err error) {
	value := uuid.NewV4().String()
	exp := os.Getenv("SESSION_EXP")
	key := "session-" + ID
	resSession := viewmodel.UserSessionVm{}
	resSession.Session = value

	uc.StoreToRedistWithExpired(key, resSession, exp)

	return value, err
}

func (uc AuthUseCase) GenerateJwtToken(ID, email, session string) (token, refreshToken, expTokenAt, expRefreshTokenAt string, err error) {
	token, expTokenAt, err = uc.JwtCred.GetToken(session, email, ID)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	refreshToken, expRefreshTokenAt, err = uc.JwtCred.GetRefreshToken(session, email, ID)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	return token, refreshToken, expTokenAt, expRefreshTokenAt, err
}

func (uc AuthUseCase) Login(input *request.LoginRequest) (res viewmodel.UserJwtTokenVm,err error){
	repository := actions.NewUserRepository(uc.DB)

	user,err := repository.ReadBy("email",input.Email)
	if err != nil {
		return res,errors.New(messages.CredentialDoNotMatch)
	}

	isPasswordMatch := hashing.CheckHashString(input.Password,user.Password)
	if !isPasswordMatch {
		return res,errors.New(messages.CredentialDoNotMatch)
	}

	IDEncrypted, _ := uc.Jwe.GenerateJwePayload(user.ID)
	emailEncrypted, _ := uc.Jwe.GenerateJwePayload(user.Email)
	session, _ := uc.UpdateSessionLogin(user.ID)

	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(IDEncrypted, emailEncrypted, session)

	res = viewmodel.UserJwtTokenVm{
		Token:           token,
		ExpTime:         tokenExpiredAt,
		RefreshToken:    refreshToken,
		ExpRefreshToken: refreshTokenExpiredAt,
	}

	return res, err
}
