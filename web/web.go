package web

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getData(url, ua string) (body []byte, err error) {
	var client = &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	if ua != "" {
		req.Header.Add("User-Agent", ua)
	}
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("获取数据失败, Code: " + strconv.Itoa(res.StatusCode))
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return
}
