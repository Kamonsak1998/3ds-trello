package models

import (
	"encoding/json"
	"github.com/adlio/trello"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

var Args = trello.Defaults()

func SetMyBoard(api string, token string) ([]*trello.Board, []error) {
	var board []*trello.Board
	a, _, err := gorequest.New().Get("https://api.trello.com/1/members/me/boards?filter=all&fields=all&lists=none&memberships=none&organization=false&organization_fields=name%2CdisplayName&limit=1000&key="+ api +"&token=" + token).End()
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(a.Body)
	json.Unmarshal(body, &board)
	return board, nil
}

func SetList(api string, idBoard string, token string) ([]*trello.List, []error) {
	var list []*trello.List
	a, _, err := gorequest.New().Get("https://api.trello.com/1/boards/" + idBoard + "/lists?cards=none&card_fields=all&limit=1000&filter=open&fields=all&key=" + api + "&token=" + token).End()
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(a.Body)
	json.Unmarshal(body, &list)
	return list, nil
}

func SetCardFinish(api string, token string, idBoard string) ([]*trello.Card, []error) {
	var card []*trello.Card
	a, _, err := gorequest.New().Get("https://api.trello.com/1/boards/" + idBoard + "/cards/?limit=1000&key=" + api + "&token=" + token).End()
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(a.Body)
	json.Unmarshal(body, &card)
	return card, nil
}

func SetMember(api string, idBoard string, token string) ([]*trello.Member, []error) {
	var member []*trello.Member
	a, _, err := gorequest.New().Get("https://api.trello.com/1/boards/" + idBoard + "/members?fields=all&key=" + api + "&token=" + token).End()
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(a.Body)
	json.Unmarshal(body, &member)
	return member, err
}

func SetStartTimeListFinish(api string, idBoard string, token string) ([]*trello.Action, []error) {
	var action []*trello.Action
	a, _, err := gorequest.New().Get("https://api.trello.com/1/boards/" + idBoard + "/actions/?filter=commentCard,updateCard:idList,createCard&limit=1000&key=" + api + "&token=" + token).End()
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(a.Body)
	json.Unmarshal(body, &action)
	return action, err
}
