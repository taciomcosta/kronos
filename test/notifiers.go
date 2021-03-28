package features

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/interfaces/rest"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// NotifiersFeature contains BDD steps related to notifiers feature
type NotifiersFeature struct {
	responseFindNotifiers *httptest.ResponseRecorder
	inputNotifier         uc.CreateNotifierRequest
}

// IProvideValidDataForNotifierCreation represents a BDD step
func (n *NotifiersFeature) IProvideValidDataForNotifierCreation() error {
	n.inputNotifier = uc.CreateNotifierRequest{
		Name: "myslack",
		Type: entities.SlackNotifierType,
		Metadata: map[string]string{
			"auth_token":  "123",
			"channel_ids": "1,2,3",
		},
	}
	return nil
}

// IProvideInvalidDataForNotifierCreation represents a BDD step
func (n *NotifiersFeature) IProvideInvalidDataForNotifierCreation() error {
	n.inputNotifier = uc.CreateNotifierRequest{}
	return nil
}

// ICreateANewNotifier represents a BDD step
func (n *NotifiersFeature) ICreateANewNotifier() error {
	request, err := newRequest(n.inputNotifier)
	n.responseFindNotifiers = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.CreateNotifier(n.responseFindNotifiers, request, ps)
	return err
}

// IListTheExistingNotifiers represents a BDD step
func (n *NotifiersFeature) IListTheExistingNotifiers() error {
	request, err := http.NewRequest("GET", "", nil)
	n.responseFindNotifiers = httptest.NewRecorder()
	ps := httprouter.Params{}
	rest.FindNotifiers(n.responseFindNotifiers, request, ps)
	return err
}

// AnErrorMessageIsShownForNotifier represents a BDD step
func (n *NotifiersFeature) AnErrorMessageIsShownForNotifier() error {
	var errorMsg rest.ErrorMessage
	err := rest.ReadJSON(n.responseFindNotifiers.Body, &errorMsg)
	if errorMsg.Msg == "" {
		return errors.New("no error message")
	}
	return err
}

// TheNewNotifierIsListed represents a BDD step
func (n *NotifiersFeature) TheNewNotifierIsListed() error {
	var findNotifiersResponse uc.FindNotifiersResponse
	err := rest.ReadJSON(n.responseFindNotifiers.Body, &findNotifiersResponse)
	if err != nil {
		return err
	}
	notifier := findNotifierByName(findNotifiersResponse, "list")
	if notifier == nil {
		return errors.New("notifier not listed when it should")
	}
	return nil
}

func findNotifierByName(response uc.FindNotifiersResponse, name string) *uc.NotifierDTO {
	for _, n := range response.Notifiers {
		if n.Name == "myslack" {
			return &n
		}
	}
	return nil
}

// IDeleteTheNewNotifier represents a BDD step
func (n *NotifiersFeature) IDeleteTheNewNotifier() error {
	request, err := http.NewRequest("DELETE", "", nil)
	response := httptest.NewRecorder()
	name := httprouter.Param{Key: "name", Value: n.inputNotifier.Name}
	params := httprouter.Params{name}
	rest.DeleteNotifier(response, request, params)
	return err
}

// TheNewNotifierIsNotListed represents a BDD step
func (n *NotifiersFeature) TheNewNotifierIsNotListed() error {
	var response uc.FindNotifiersResponse
	err := rest.ReadJSON(n.responseFindNotifiers.Body, &response)
	if err != nil {
		return err
	}
	notifier := findNotifierByName(response, n.inputNotifier.Name)
	if notifier != nil {
		return errors.New("notifier was listed when it should not")
	}
	return nil
}
