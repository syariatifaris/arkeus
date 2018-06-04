package panics

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime/debug"

	"github.com/syariatifaris/arkeus/core/log/tokolog"
)

var (
	cfg *SlackConfig
	wg  sync.WaitGroup
)

func InitSlack(slackCfg *SlackConfig) {
	cfg = slackCfg
}

type SlackConfig struct {
	Enabled     bool
	Channel     string
	WebHookURL  string
	EnabledEnvs []string
}

//Restore the panic
func Restore() {
	if x := recover(); x != nil {
		tokolog.ERROR.Println("[Panic][Restore]", x)
		stackTrace := string(debug.Stack())
		tokolog.ERROR.Println("[Panic][Restore]", stackTrace)
		if cfg != nil {
			if cfg.Enabled {
				wg.Add(1)
				go postToSlack(fmt.Sprint("Panic:", x), stackTrace)
				wg.Wait()
			}
		}
	}
}

//postToSlack posts the panic capture stack trace to slack
func postToSlack(text, snip string) {
	defer func() {
		wg.Done()
	}()
	payload := map[string]interface{}{
		"text": text,
		//Enable slack to parse mention @<someone>
		"link_names": 1,
		"attachments": []map[string]interface{}{
			map[string]interface{}{
				"text":      snip,
				"color":     "#e50606",
				"title":     "Stack Trace",
				"mrkdwn_in": []string{"text"},
			},
		},
	}
	if cfg.Channel != "" {
		payload["channel"] = cfg.Channel
	}
	b, err := json.Marshal(payload)
	if err != nil {
		tokolog.INFO.Println("[panics] marshal err", err, text, snip)
		return
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(cfg.WebHookURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		tokolog.INFO.Printf("[panics] error on capturing error : %s %s %s\n", err.Error(), text, snip)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			tokolog.INFO.Printf("[panics] error on capturing error : %s %s %s\n", err, text, snip)
			return
		}
		tokolog.INFO.Printf("[panics] error on capturing error : %s %s %s\n", string(b), text, snip)
	}
}
