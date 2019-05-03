package lunch

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

 func Lunch(w http.ResponseWriter, r *http.Request) {
	 if r.Method != "POST" {
		 e := "Method Not Allowed."
		 log.Println(e)
		 w.WriteHeader(http.StatusMethodNotAllowed)
		 w.Write([]byte(e))
		 return
	 }

	 b, err := ioutil.ReadAll(r.Body)
	 if err != nil {
		 log.Printf("ReadAllError: %v\n", err)
		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write([]byte(err.Error()))
		 return
	 }

	 parsed, err := url.ParseQuery(string(b))
	 if err != nil {
		 log.Printf("ParseQueryError: %v\n", err)
		 w.WriteHeader(http.StatusInternalServerError)
		 w.Write([]byte(err.Error()))
		 return
	 }

	 if parsed.Get("token") != os.Getenv("SLACK_TOKEN") {
		 e := "Unauthorized Token."
		 log.Printf(e)
		 w.WriteHeader(http.StatusUnauthorized)
		 w.Write([]byte(e))
		 return
	 }

	 w.WriteHeader(http.StatusOK)
	 w.Write([]byte(parsed.Get("text")))
 }