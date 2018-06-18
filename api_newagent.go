package chainetix

func (client *Client) NewAgent(label string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"label": label,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/agent",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}