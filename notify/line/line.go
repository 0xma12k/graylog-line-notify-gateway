package line

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

// Notify ...
func Notify(message, token string) (int, []byte) {
	endpoint := "https://notify-api.line.me/api/notify"
	tokenBearer := "Bearer " + token
	data := url.Values{}
	data.Set("message", message)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		logrus.Error(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", tokenBearer)

	resp, err := client.Do(r)
	if err != nil {
		logrus.Error(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	resMsg := string(bodyBytes)

	logrus.WithFields(
		logrus.Fields{
			"status": resp.StatusCode,
			"resp":   resMsg,
		}).Debug("called service line notify service")

	defer resp.Body.Close()

	return resp.StatusCode, bodyBytes
}
