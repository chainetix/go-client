package chainetix

func (client *Client) DeleteUsers() error {
	src := map[string]interface{}{
	}
	_, err := client.Delete(
		"https://multichain.blokhub.io/api/users",
		src,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}