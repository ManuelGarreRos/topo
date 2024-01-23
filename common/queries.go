package common

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BaseQuery struct {
	Pager  Pager
	Sorts  []Sort
	Fields []string
}

type Pager struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Sort struct {
	Field string `json:"field"`
	Desc  bool   `json:"desc"`
}

func (q *BaseQuery) OrderBy() string {
	res := ""
	if q.Sorts != nil {
		for i := range q.Sorts {
			if q.Sorts[i].Field != "" {
				if q.Sorts[i].Desc {
					res += q.Sorts[i].Field + " desc"
				} else {
					res += q.Sorts[i].Field
				}
			}
			if i < len(q.Sorts)-1 {
				res += ","
			}
		}
	}

	return res
}

func (q *BaseQuery) ParseBase(ctx *gin.Context) error {
	q.parsePager(ctx)

	if err := q.parseSort(ctx); err != nil {
		return err
	}

	return nil
}

func (q *BaseQuery) parsePager(ctx *gin.Context) {
	limQ := ctx.Query("lim")
	offsetQ := ctx.Query("offset")

	lim, err := strconv.ParseInt(limQ, 10, 0)
	if err != nil {
		lim = 100
	}

	offset, err := strconv.ParseInt(offsetQ, 10, 0)
	if err != nil {
		offset = 0
	}

	q.Pager = Pager{
		Limit:  int(lim),
		Offset: int(offset),
	}
}

func (q *BaseQuery) parseSort(ctx *gin.Context) error {
	q.Sorts = make([]Sort, 0)

	s := ctx.Query("s")
	if s == "" {
		return nil
	}

	jn, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return NewBadRequest("incorrect query params", "", s)
	}

	if err := json.Unmarshal(jn, &q.Sorts); err != nil {
		return NewBadRequest("incorrect query params", "", jn)
	}

	return nil
}

func ParseQuery(ctx *gin.Context, q interface{}) error {
	qs := ctx.Query("q")
	if qs == "" {
		return nil
	}

	jn, err := base64.StdEncoding.DecodeString(qs)
	if err != nil {
		return NewBadRequest("incorrect query params", "", qs)
	}

	if err := json.Unmarshal(jn, &q); err != nil {
		return NewBadRequest("incorrect query params", "", jn)
	}

	return nil
}

type BoolSearch struct {
	HasValue bool `json:"hasValue"`
	Value    bool `json:"value"`
}

// NewBadRequest creates a new bad request error
func NewBadRequest(message string, field string, value interface{}) error {
	return &BadRequestError{Message: message, Field: field, Value: value}
}

// BadRequestError represents a bad request error
type BadRequestError struct {
	Message string      `json:"message"`
	Field   string      `json:"field,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}

// Error returns the error message
func (e *BadRequestError) Error() string {
	return e.Message
}
