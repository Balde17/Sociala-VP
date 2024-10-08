package jwt

import (
	"backend/app/internals/utils"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type jwt struct {
	header  map[string]interface{}
	payload map[string]interface{}
}

func New() *jwt {
	header := make(map[string]interface{})
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	payload := make(map[string]interface{})
	return &jwt{header: header, payload: payload}
}

func (j *jwt) Sign(secret string) (string, error) {
	unsigned, err := j.toUnsigned()
	if err != nil {
		return "", err
	}

	hash := hmac.New(sha256.New, []byte(secret))
	_, err = hash.Write([]byte(unsigned))
	if err != nil {
		return "", err
	}

	signed := base64Encode(hash.Sum(nil))

	return join(unsigned, signed), nil
}

func Parse(token string) (*jwt, error) {
	splited := split(token)
	if len(splited) != 3 {
		return nil, errors.New("invalid token")
	}
	header, payload := splited[0], splited[1]

	jwt, err := decodeJWT(header, payload)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

func (j *jwt) Verify(token, secret string) error {
	compare, err := j.Sign(secret)
	if err != nil {
		return err
	}
	if compare != token {
		return errors.New("invalid token")
	}
	return nil
}

func (j *jwt) SetPayload(key string, data interface{}) {
	j.payload[key] = data
}

func (j *jwt) Payload(key string) (interface{}, bool) {
	data, ok := j.payload[key]
	return data, ok
}

func (j *jwt) toUnsigned() (string, error) {
	headerBytes, err := mapToJSONBytes(&j.header)
	if err != nil {
		return "", err
	}
	headerBase64 := base64Encode(headerBytes)

	payloadBytes, err := mapToJSONBytes(&j.payload)
	if err != nil {
		return "", err
	}
	payloadBase64 := base64Encode(payloadBytes)

	return join(headerBase64, payloadBase64), nil
}

func base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func base64Decode(str string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func mapToJSONBytes(m *map[string]interface{}) ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func mapFromJSONBytes(j []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal(j, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func decodeJWT(header, payload string) (*jwt, error) {
	payloadMap, err := decodeString(payload)
	if err != nil {
		return nil, err
	}
	jwt := New()
	jwt.payload = payloadMap
	return jwt, nil
}

func decodeString(s string) (map[string]interface{}, error) {
	b, err := base64Decode(s)
	if err != nil {
		return nil, err
	}

	d, err := mapFromJSONBytes(b)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func join(strs ...string) string {
	return strings.Join(strs, ".")
}

func split(str string) []string {
	return strings.Split(str, ".")
}

func CryptCode(password string) (res string) {
	for _, value := range password {
		v := (value + rune(len(password)+4)) % 127
		res += string(v)
	}
	return RotateWord(res)
}

func DecryptCode(encrypt string) (password string) {
	encrypted := RotateWord(encrypt)
	for _, value := range encrypted {
		v := (value - rune(len(encrypted)+4) + 127) % 127
		password += string(v)
	}
	return password
}

func RotateWord(mot string) string {
	runes := []rune(mot)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	motInverse := string(runes)
	return motInverse
}

func GetPayload(r *http.Request) map[string]interface{} {

	token := utils.GetJWT(r)

	_jwt, err := Parse(token)

	if err != nil {
		return nil
	}

	_payload, _ := _jwt.Payload("payload")

	payloadMap, ok := _payload.(map[string]interface{})
	if !ok {
		fmt.Println("Erreur lors de la conversion du payload en map.")
		return nil
	}

	return payloadMap["User"].(map[string]interface{})
}
