package platform_pwdhasher

import (
	"golang.org/x/crypto/bcrypt"
)

var _ PwdHasher = (*PwdHasherImpl)(nil)

type PwdHasherImpl struct{}

func NewPwdHasherImpl() PwdHasher {
	return &PwdHasherImpl{}
}

func (h *PwdHasherImpl) Hash(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
}

func (h *PwdHasherImpl) CompareHashAndPassword(hash, password []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}
