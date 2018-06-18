package chainetix

func (client *Client) Agents() ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.chainetix.com/api/agents",
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}