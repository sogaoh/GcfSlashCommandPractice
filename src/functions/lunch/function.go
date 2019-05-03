package lunch

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Lunch(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	log.Printf("%v", string(b))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
