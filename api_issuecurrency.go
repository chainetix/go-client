package chainetix

func (client *Client) IssueCurrency(currency string, quantity float64, recipientAddress string) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	src := map[string]interface{}{
		"quantity": quantity,
		"recipientAddress": recipientAddress,
	}
	_, err := client.Post(
		"https://multichain.chainetix.com/api/currency/" + currency + "/issue",
		src,
		&dest,
		client.headers,
	)
	if err != nil {
		return nil, err
	}
	return dest, nil
}