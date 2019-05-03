package lunch

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Parameter struct {
	SubCommand string
	Value      string
}

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

	p := new(Parameter)
	p.parse(parsed.Get("text"))

	switch p.SubCommand {
	case "add":
		//TODO add の処理を追加

	case "list":
		//TODO list の処理を追加

	default:
		e := "Invalid SubCommand."
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(parsed.Get("text")))
}


func (p *Parameter) parse(text string) {
	t := strings.TrimSpace(text)
	if len(t) < 1 {
		return
	}

	s := strings.SplitN(t, " ", 2)
	p.SubCommand = s[0]

	if len(s) == 1 {
		return
	}

	p.Value = s[1]
}