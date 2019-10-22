
package firebase

import "fmt"

// ErrorResponse FirebaseのError Responseの受け口
// https://firebase.google.com/docs/reference/rest/auth/#section-error-format
type ErrorResponse struct {
	Error *Error `json:"error"`
}

// Error FirebaseAPIを格納
type Error struct {
	Errors  []Detail `json:"errors"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
}

// Error Firebaseのエラー内容をerrorインターフェースに適応させる
func (e *Error) Error() string {
	var mes string
	for _, v := range e.Errors {
		mes = fmt.Sprintf("%sdomain: %s, reason: %s, message: %s\n", mes, v.Domain, v.Reason, v.Message)
	}
	return mes
}

// Detail Firebase Erros単体
type Detail struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}