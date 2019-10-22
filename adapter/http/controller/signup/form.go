package signup

import "net/http"

func (s *Signup) Form(w http.ResponseWriter, r *http.Request) {
	if err := s.re.Form(w); err != nil {
		panic(err)
	}
}
