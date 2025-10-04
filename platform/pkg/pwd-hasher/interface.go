package platform_pwdhasher

type PwdHasher interface {
	Hash(data []byte) ([]byte, error)
	CompareHashAndPassword(hash, password []byte) bool
}
