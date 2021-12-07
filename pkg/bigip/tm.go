package bigip

func (c *client) GTM() *GTM {
	return &GTM{c: c}
}
