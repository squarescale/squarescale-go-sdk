package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/squarescale/squarescale-go-sdk/models"
)

// CreateProject will call http endpoint(s) required to create
// a new Project. This call is blocking.
func (c *Client) CreateProject(project models.Project, ctx context.Context) error {
	var err error
	var payload []byte
	var req *http.Request
	var res *http.Response

	payload, err = json.Marshal(project)

	if err != nil {
		return err
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, c.address, bytes.NewReader(payload))

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "bearer "+c.apiKey)
	res, err = c.httpClient().Do(req)

	if res.StatusCode != http.StatusCreated {
		var errBody []byte
		defer res.Body.Close()
		errBody, err = ioutil.ReadAll(res.Body)

		if err != nil {
			return err
		}

		msg := fmt.Sprintf("Unexpected http code receive (%d) with the message: %s", res.StatusCode, string(errBody))
		return errors.New(msg)
	}

	return nil
}
