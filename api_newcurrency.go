package chainetix

func (client *Client) NewCurrency(name string, units float64) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"name": name,
		"units": units,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/currency",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}