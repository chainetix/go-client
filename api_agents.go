package chainetix

func (client *Client) Agents() ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/agents",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}