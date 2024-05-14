package common

import "errors"

type Credential struct {
	ak string
	sk string
}

func (c Credential) Ak() string {
	return c.ak
}

func (c Credential) Sk() string {
	return c.sk
}

const (
	AkLen = 32
	SkLen = 32
)

// CheckAk 校验ak
func CheckAk(ak string) error {
	if len(ak) != AkLen {
		return errors.New("ak长度必须为32")
	}
	return nil
}

// CheckSk 校验sk
func CheckSk(sk string) error {
	if len(sk) != SkLen {
		return errors.New("sk长度必须为32")
	}
	return nil
}

func NewCredential(ak string, sk string) (*Credential, error) {
	err := CheckAk(ak)
	if err != nil {
		return nil, err
	}

	err = CheckAk(sk)
	if err != nil {
		return nil, err
	}
	return &Credential{ak: ak, sk: sk}, nil
}
