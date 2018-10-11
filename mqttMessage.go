package mqttconnector

// {"op":"remove","data":{"_id":"twitter_88372467"}}

type Message struct {
	Op    string `json:"op"`
	Topic string `json:"-"`
	Data  Data   `json:"data"`
}

type Data struct {
	ID         string                 `json:"_id"`
	SocialType string                 `json:"social_type"`
	Zone       string                 `json:"zone"`
	Interval   int64                  `json:"interval"`
	AddBy      map[string]interface{} `json:"add_by"`
	Info       UserInfo               `json:"info"`
	AccountID  []string               `json:"account_id"`
	V          interface{}            `json:"__v"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Picture  string `json:"picture"`
	SocialID string `json:"social_id"`
}
