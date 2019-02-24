package glo

import (
	"log"

	"github.com/levigross/grequests"
)

// Glo Base Struct for the library and get the paths and the token for access to the API
type Glo struct {
	baseURLApi     string
	pathBoards     string
	pathColumn     string
	pathCard       string
	pathAttachment string
	pathComment    string
	pathUser       string
	token          string
}

// NewGloClient returns a new Glo Api Client
func NewGloClient(token string) Glo {

	g := Glo{
		baseURLApi:     "https://gloapi.gitkraken.com/v1/glo",
		pathBoards:     "/boards",
		pathColumn:     "/columns",
		pathCard:       "/cards",
		pathAttachment: "/attachments",
		pathComment:    "/comments",
		pathUser:       "/user",
		token:          token,
	}

	return g
}

func (g Glo) prepareParams(fields []string) (params string) {

	for _, f := range fields {
		if params == "" {
			params = "?fields=" + f
		} else {
			params = params + "&fields=" + f
		}

	}
	return
}

// Unified method for call the Glo Api and put the authorization header with the token
func (g Glo) callAPI(method string, endpoint string, fields []string, data interface{}) (resp *grequests.Response) {
	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Authorization": "Bearer " + g.token},
	}

	if data != nil {
		ro.JSON = data
	}

	params := ""

	if method == "GET" && fields != nil {
		params = g.prepareParams(fields)
	}

	log.Println(g.baseURLApi + endpoint + params)

	resp, err := grequests.Req(method, g.baseURLApi+endpoint+params, ro)

	if err != nil {
		log.Fatal(err)
	}

	return
}

// GetBoards Get all the boards that read the token
func (g Glo) GetBoards(fields []string) (boards []Board) {
	resp := g.callAPI("GET", g.pathBoards, fields, nil)

	if resp.Ok {
		err := resp.JSON(&boards)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}
	return
}

// GetBoard Get a single board
func (g Glo) GetBoard(id string, fields []string) (board Board) {
	resp := g.callAPI("GET", g.pathBoards+"/"+id, fields, nil)

	if resp.Ok {
		err := resp.JSON(&board)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}
	return
}

// CreateBoardColumn Create a column in the board that pass with the id
func (g Glo) CreateBoardColumn(boardID string, data CreateBoardColumnParams) (column BoardColumn) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathColumn,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&column)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// EditBoardColumn Edit a column in a single board
func (g Glo) EditBoardColumn(boardID string, columnID string, data CreateBoardColumnParams) (column BoardColumn) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathColumn+"/"+columnID,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&column)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// DeleteBoardColumn Delete a column in a single board
func (g Glo) DeleteBoardColumn(boardID string, columnID string) {
	resp := g.callAPI(
		"DELETE",
		g.pathBoards+"/"+boardID+g.pathColumn+"/"+columnID,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete column")
	}
}

// GetBoardCards Get all the cards inside board
func (g Glo) GetBoardCards(boardID string, fields []string) (cards []BoardCard) {
	resp := g.callAPI(
		"GET",
		g.pathBoards+"/"+boardID+g.pathCard,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&cards)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// GetBoardCard Get a single card of a board
func (g Glo) GetBoardCard(boardID string, cardID string, fields []string) (card BoardCard) {
	resp := g.callAPI(
		"GET",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&card)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// CreateBoardCard Create a card inside board
func (g Glo) CreateBoardCard(boardID string, data BoardCardParams) (card BoardCard) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathCard,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&card)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// EditBoardCard Edit a card of board
func (g Glo) EditBoardCard(boardID string, cardID string, data BoardCardParams) (card BoardCard) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&card)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}
	return
}

// DeleteBoardCard Delete a card of board
func (g Glo) DeleteBoardCard(boardID string, cardID string) {
	resp := g.callAPI(
		"DELETE",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete card")
	}
}

// GetBoardColumnCards Get all the cards of a single column of board
func (g Glo) GetBoardColumnCards(boardID string, columnID string, fields []string) (cards []BoardCard) {
	resp := g.callAPI(
		"GET",
		g.pathBoards+"/"+boardID+g.pathColumn+"/"+columnID+"/"+g.pathCard,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&cards)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// GetBoardCardAttachments Get all the attachments of a single card of board
func (g Glo) GetBoardCardAttachments(boardID string, cardID string, fields []string) (attachments []Attachment) {
	resp := g.callAPI(
		"GET",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID+"/"+g.pathAttachment,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&attachments)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// GetBoardCardComments Get all the comments of a card
func (g Glo) GetBoardCardComments(boardID string, cardID string, fields []string) (comments []Comment) {
	resp := g.callAPI(
		"GET",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID+"/"+g.pathComment,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&comments)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// CreateBoardCardComment Add a comment to a card
func (g Glo) CreateBoardCardComment(boardID string, cardID string, data ParamComment) (comment Comment) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID+"/"+g.pathComment,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&comment)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// EditBoardCardComment Edit a comment of card
func (g Glo) EditBoardCardComment(boardID string, cardID string, commentID string, data ParamComment) (comment Comment) {
	resp := g.callAPI(
		"POST",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID+"/"+g.pathComment+"/"+commentID,
		nil,
		data,
	)

	if resp.Ok {
		err := resp.JSON(&comment)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}

// DeleteBoardCardComment Delete a comment of card
func (g Glo) DeleteBoardCardComment(boardID string, cardID string, commentID string) {
	resp := g.callAPI(
		"DELETE",
		g.pathBoards+"/"+boardID+g.pathCard+"/"+cardID+"/"+g.pathComment+"/"+commentID,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete comment")
	}
}

// GetUser Get information of the connect user
func (g Glo) GetUser(fields []string) (user User) {
	resp := g.callAPI(
		"GET",
		g.pathUser,
		fields,
		nil,
	)

	if resp.Ok {
		err := resp.JSON(&user)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}

	return
}
