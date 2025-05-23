package gonextcloud

import (
	"fmt"

	req "github.com/levigross/grequests"
)

var errUnauthorized = fmt.Errorf("login first")

// Login perform login and create a session with the Nextcloud API.
func (c *client) LoginWithOptions(username string, password string, options req.RequestOptions ) error {
	c.username = username
	c.password = password
	
	if len(options.Headers) == 0 {
		options.Headers = c.headers
	}
	if len(options.Auth) == 0 {
		options.Auth = []string{c.username, c.password}
	}
	c.session = req.NewSession(&options)
	// TODO What to do with capabilities ? (other thant connection validation)
	u := c.baseURL.ResolveReference(routes.capabilities)
	res, err := c.session.Get(u.String(), nil)
	if err != nil {
		return err
	}
	var r capabilitiesResponse
	res.JSON(&r)
	// No need to check for Ocs.meta.StatusCode as capabilities are always returned
	c.capabilities = &r.Ocs.Data.Capabilities
	c.version = &r.Ocs.Data.Version
	// Check if authentication failed
	if !c.loggedIn() {
		e := APIError{Message: "authentication failed"}
		return &e
	}
	// Create webdav client
	c.webdav = newWebDav(c.baseURL.String()+"/remote.php/webdav", c.username, c.password)
	for k, v := range options.Headers {
		c.webdav.SetHeader(k, v)
	}
	return nil
}
func (c *client) Login(username string, password string) error {
	options := req.RequestOptions{
		Headers: c.headers,
		Auth:    []string{c.username, c.password},
	}
	return c.LoginWithOptions(username, password, options)
}

// Logout logs out from the Nextcloud API, close the session and delete session's cookie
func (c *client) Logout() error {
	c.session.CloseIdleConnections()
	c.session.HTTPClient.Jar = nil
	// Clear capabilities as it is used to check for valid authentication
	c.capabilities = nil
	return nil
}

func (c *client) loggedIn() bool {
	// When authentication failed, capabilities doesn't contains core information
	if c.capabilities == nil {
		return false
	}
	return c.capabilities.Core.WebdavRoot != ""
}
