package chainetix

func (client *Client) Currencies() ([]interface{}, error) {
	dest := []interface{}{}
	_, err := client.Get(
		"https://multichain.blokhub.io/api/currencies",
		&dest,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}