package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"backend/pkg/model"
)

const (
	AuthorizationHeader = "Authorization"
	labMarkToken        = "lab-mark-token"
)

type externalService struct {
}

func NewExternalService() *externalService {
	return &externalService{}
}

func (s *externalService) SendLabMark(ctx context.Context, userId, labId, percentage int) error {
	url := fmt.Sprintf("%s/lab", os.Getenv("EXTERNAL_APP_HOST"))

	labMark := model.UserLabMark{
		UserId:     userId,
		LabId:      labId,
		Percentage: percentage,
	}

	data, err := json.Marshal(labMark)
	if err != nil {
		log.Error(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Error(err)
		return err
	}

	req.Header.Add(labMarkToken, os.Getenv("EXTERNAL_AUTH_HEADER"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer res.Body.Close()

	if _, err := ioutil.ReadAll(res.Body); err != nil {
		log.Error(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("lab mark status not 200: %d", res.StatusCode)
	}

	return nil
}

func (s *externalService) GetUserId(ctx context.Context, token string) (int, error) {
	url := fmt.Sprintf("%s/lab", os.Getenv("EXTERNAL_APP_HOST"))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return 0, err
	}

	req.Header.Add(labMarkToken, os.Getenv("EXTERNAL_AUTH_HEADER"))
	req.Header.Add(AuthorizationHeader, token)
	req.Header.Add("lab-mark-token", os.Getenv("EXTERNAL_AUTH_HEADER"))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var userId model.UserId
	if err := json.Unmarshal(data, &userId); err != nil {
		return 0, err
	}

	if userId.UserId == 0 {
		return 0, fmt.Errorf("not exist userId")
	}

	return userId.UserId, nil
}
