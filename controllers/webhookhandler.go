package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/0xma12k/graylog-line-notify-gateway/internal/entity"
	"github.com/0xma12k/graylog-line-notify-gateway/internal/templ"
	"github.com/0xma12k/graylog-line-notify-gateway/notify/line"
	log "github.com/sirupsen/logrus"
)

// Webhook ...
func Webhook(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	templName := params.Get("template")
	lineToken := params.Get("token")

	// Decode reqeust body to Graylog struct object
	var graylog entity.GraylogJSON
	if err := json.NewDecoder(r.Body).Decode(&graylog); err != nil {
		log.Error(err)
	}

	// Execute template
	msg, err := templ.ExecuteTemplate(templName, &graylog)
	if err != nil {
		log.Error(err)

		defaultMsg, err := templ.DefaultExecute(&graylog)
		if err != nil {
			log.Fatal(err)
		}
		msg = defaultMsg
	}

	statusCode, bodyBytes := line.Notify(msg, lineToken)

	w.WriteHeader(statusCode)
	w.Write(bodyBytes)

	defer r.Body.Close()
}
