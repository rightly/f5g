package main

func (c *client) GTM() *GTM {
	return &GTM{c: c}
}
