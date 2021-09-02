package functions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewRequestCampaign(mobile_number, campaign_link string) (string, error) {
	campaign_code := strings.Replace(campaign_link, "https://gift.truemoney.com/compaign?v=", "", -1)
	redeem_url := "https://gift.truemoney.com/campaign/vouchers/" + campaign_code + "/redeem"
	payload, _ := json.Marshal(map[string]string{"mobile": mobile_number})
	reader_buffer := bytes.NewBuffer(payload)
	response_campaign, err := http.Post(redeem_url, "application/json", reader_buffer)
	if err != nil {
		return "", err
	}
	defer response_campaign.Body.Close()

	body, _ := ioutil.ReadAll(response_campaign.Body)

	return string(body), nil
}
