package trelloservice

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	trello "github.com/adlio/trello"

	domain "github.com/devpablocristo/nanlabs/domain"
)

var (
	CARD_BASE_URL        string = "https://api.trello.com/1/cards"
	ID_LIST              string = "63c490512d8e4301239ceaa0"
	ID_BOARD             string = "63c4903f1b996f00571f777b"
	ID_MEMBER            string = "5d09b19d2b2fe58a92132bad"
	ID_LABEL_MAINTENANCE string = "63e40f80b1c831ec59bd0d5c"
	ID_LABEL_RESEARCH    string = "63e40f91208e91c4f5cbb5ac"
	ID_LABEL_TEST        string = "63e40f9dcc75976d321b4e7e"
)

type TrelloService struct {
	trelloClient *trello.Client
}

func NewTrelloService(key string, token string) *TrelloService {
	c := trello.NewClient(key, token)
	return &TrelloService{
		trelloClient: c,
	}
}

func (ts *TrelloService) CreateIssueCard(ctx context.Context, task *domain.Task) error {
	log.Println("Create Issue Card")

	card := &trello.Card{
		IDList: ID_LIST,
		Name:   task.Title,
		Desc:   task.Description,
		URL:    CARD_BASE_URL,
	}

	err := ts.createCardStd(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ts.createCardAdlio(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ts *TrelloService) CreateBugCard(ctx context.Context, task *domain.Task) error {
	log.Println("Create Bug Card")

	member, err := ts.getRandomMemberFromBoard(ID_BOARD)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	card := &trello.Card{
		IDList: ID_LIST,
		Name:   task.Title,
		Desc:   task.Description,
		URL:    CARD_BASE_URL,
		IDMembers: []string{
			member.ID,
		},
	}

	err = ts.createCardStd(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = ts.createCardAdlio(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ts *TrelloService) CreateTaskCard(ctx context.Context, task *domain.Task) error {
	log.Println("Create Task Card")

	var label string
	switch task.Category {
	case "Maintenance":
		label = ID_LABEL_MAINTENANCE
	case "Test":
		label = ID_LABEL_TEST
	case "Research":
		label = ID_LABEL_RESEARCH
	default:
		return errors.New("incorrect task category")
	}

	card := &trello.Card{
		IDList: ID_LIST,
		Name:   task.Title,
		URL:    CARD_BASE_URL,
		IDLabels: []string{
			label,
		},
	}

	err := ts.createCardStd(card)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ts *TrelloService) createCardAdlio(newCard *trello.Card) error {
	list, err := ts.trelloClient.GetList(newCard.IDList, trello.Defaults())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = list.AddCard(newCard, trello.Defaults())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ts *TrelloService) createCardStd(newCard *trello.Card) error {
	params := url.Values{}
	params.Add("idList", ID_LIST)
	params.Add("key", ts.trelloClient.Key)
	params.Add("token", ts.trelloClient.Token)
	params.Add("name", newCard.Name)
	params.Add("idLabels", strings.Join(newCard.IDLabels, ","))
	params.Add("idMembers", strings.Join(newCard.IDMembers, ","))
	params.Add("desc", newCard.Desc)

	_, err := http.PostForm(newCard.URL, params)
	if err != nil {
		log.Printf("Request Failed: %s", err)
		return err
	}

	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("Request Failed: %s", err)
	// 	return err
	// }

	// bodyString := string(body)
	// log.Print(bodyString)

	return nil

}

func (ts *TrelloService) getRandomMemberFromBoard(idBoard string) (*trello.Member, error) {
	b, err := ts.trelloClient.GetBoard(idBoard, trello.Defaults())
	if err != nil {
		log.Println(err.Error())
		return &trello.Member{}, err
	}

	members, err := b.GetMembers(trello.Defaults())
	if err != nil {
		log.Println(err.Error())
		return &trello.Member{}, err
	}

	randomIndex := rand.Intn(len(members))
	pickMemember := members[randomIndex]

	return pickMemember, nil
}
