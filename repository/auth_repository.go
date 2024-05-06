package repository

type AuthRepository struct {
}

func InitAuthRepository() AuthRepository {
	return AuthRepository{}
}

func (repo *AuthRepository) FindEmail(email string) {

}
