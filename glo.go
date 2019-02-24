package glo

import (
	"github.com/levigross/grequests"
	"log"
)

type Glo struct {
	baseUrlApi     string
	pathBoards     string
	pathColumn     string
	pathCard       string
	pathAttachment string
	pathComment    string
	pathUser	   string
	token          string
}

// NewGloClient returns a new Glo Api Client
func NewGloClient(token string) Glo {

	g := Glo{
		baseUrlApi:     "https://gloapi.gitkraken.com/v1/glo",
		pathBoards:     "/boards",
		pathColumn:     "/columns",
		pathCard:       "/cards",
		pathAttachment: "/attachments",
		pathComment:    "/comments",
		pathUser:		"/user",
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
func (g Glo) callApi(method string, endpoint string, fields []string, data interface{}) (resp *grequests.Response) {
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

	log.Println(g.baseUrlApi + endpoint + params)

	resp, err := grequests.Req(method, g.baseUrlApi+endpoint+params, ro)

	if err != nil {
		log.Fatal(err)
	}

	return
}

// Get all the boards that read the token
func (g Glo) GetBoards(fields []string) (boards []Board) {
	resp := g.callApi("GET", g.pathBoards, fields, nil)

	if resp.Ok {
		err := resp.JSON(&boards)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}
	return
}

// Get a single board
func (g Glo) GetBoard(id string, fields []string) (board Board) {
	resp := g.callApi("GET", g.pathBoards+"/"+id, fields, nil)

	if resp.Ok {
		err := resp.JSON(&board)
		if err != nil {
			log.Fatal("Unmarshal error")
		}
	}
	return
}

// Create a column in the board that pass with the id
func (g Glo) CreateBoardColumn(boardId string, data CreateBoardColumnParams) (column BoardColumn) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathColumn,
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

// Edit a column in a single board
func (g Glo) EditBoardColumn(boardId string, columnId string, data CreateBoardColumnParams) (column BoardColumn) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathColumn+"/"+columnId,
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

// Delete a column in a single board
func (g Glo) DeleteBoardColumn(boardId string, columnId string) {
	resp := g.callApi(
		"DELETE",
		g.pathBoards+"/"+boardId+g.pathColumn+"/"+columnId,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete column")
	}
}

// Get all the cards inside board
func (g Glo) GetBoardCards(boardId string, fields []string) (cards []BoardCard) {
	resp := g.callApi(
		"GET",
		g.pathBoards+"/"+boardId+g.pathCard,
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

// Get a single card of a board
func (g Glo) GetBoardCard(boardId string, cardId string, fields []string) (card BoardCard) {
	resp := g.callApi(
		"GET",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId,
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

// Create a card inside board
func (g Glo) CreateBoardCard(boardId string, cardId string, data BoardCardParams) (card BoardCard) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathCard,
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

// Edit a card of board
func (g Glo) EditBoardCard(boardId string, cardId string, data BoardCardParams) (card BoardCard) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId,
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

// Delete a card of board
func (g Glo) DeleteBoardCard(boardId string, cardId string) {
	resp := g.callApi(
		"DELETE",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete card")
	}
}

// Get all the cards of a single column of board
func (g Glo) GetBoardColumnCards(boardId string, columnId string, fields []string) (cards []BoardCard) {
	resp := g.callApi(
		"GET",
		g.pathBoards+"/"+boardId+g.pathColumn+"/"+columnId+"/"+g.pathCard,
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

// Get all the attachments of a single card of board
func (g Glo) GetBoardCardAttachments(boardId string, cardId string, fields []string) (attachments []Attachment) {
	resp := g.callApi(
		"GET",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId+"/"+g.pathAttachment,
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

// Get all the comments of a card
func (g Glo) GetBoardCardComments(boardId string, cardId string, fields []string) (comments []Comment) {
	resp := g.callApi(
		"GET",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId+"/"+g.pathComment,
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

// Add a comment to a card
func (g Glo) CreateBoardCardComment(boardId string, cardId string, data ParamComment) (comment Comment) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId+"/"+g.pathComment,
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

// Edit a comment of card
func (g Glo) EditBoardCardComment(boardId string, cardId string, commentId string, data ParamComment) (comment Comment) {
	resp := g.callApi(
		"POST",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId+"/"+g.pathComment+"/"+commentId,
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

// Delete a comment of card
func (g Glo) DeleteBoardCardComment(boardId string, cardId string, commentId string) {
	resp := g.callApi(
		"DELETE",
		g.pathBoards+"/"+boardId+g.pathCard+"/"+cardId+"/"+g.pathComment+"/"+commentId,
		nil,
		nil,
	)

	if !resp.Ok {
		log.Fatal("Error at delete comment")
	}
}

// Get information of the connect user
func (g Glo) GetUser(fields []string) (user User){
	resp := g.callApi(
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