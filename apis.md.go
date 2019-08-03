// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// execGetAccessToken 获取access_token
func (c *WorkwxApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}

	return resp, nil
}

// execUserGet 读取成员
func (c *WorkwxApp) execUserGet(req reqUserGet) (respUserGet, error) {
	var resp respUserGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return respUserGet{}, err
	}

	return resp, nil
}

// execDeptList 获取部门列表
func (c *WorkwxApp) execDeptList(req reqDeptList) (respDeptList, error) {
	var resp respDeptList
	err := c.executeQyapiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return respDeptList{}, err
	}

	return resp, nil
}

// execMessageSend 发送应用消息
func (c *WorkwxApp) execMessageSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}

	return resp, nil
}

// execAppchatSend 应用推送消息
func (c *WorkwxApp) execAppchatSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}

	return resp, nil
}
