package chainetix

func (client *Client) NewAgent(label string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"label": label,
	}
	_, err := client.Post(
		"https://multichain.blokhub.io/api/agent",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}