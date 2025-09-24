package app

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"
)

type ModelInterface interface {
	grest.ModelInterface
}

type Model struct {
	grest.Model
}

type ListSingleModel struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"timestamp"`
	Results   any       `json:"results"`

	Ctx *Ctx `json:"-"`
}

func (list *ListSingleModel) SetData(data any, query url.Values) {
	list.Results = data
	list.Status = "Success"

	switch list.Ctx.Action.Method {
	case "GET":
		list.Message = "Data retrivied successfully"
	case "POST":
		list.Message = "Data created successfully"
	case "PUT":
		list.Message = "Data updated successfully"
	case "DELETE":
		list.Message = "Data deleted successfully"
	default:
		list.Message = "Data retrivied successfully"
	}

	list.TimeStamp = time.Now()
}

type ListModel struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"timestamp"`
	Results   struct {
		Data        []map[string]any `json:"list"`
		PageContext struct {
			Page      int   `json:"page"`
			PerPage   int   `json:"limit"`
			Count     int64 `json:"total"`
			PageCount int   `json:"total_pages"`
			Links     struct {
				First    string `json:"first"`
				Previous string `json:"previous"`
				Next     string `json:"next"`
				Last     string `json:"last"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"results"`
}

func (list *ListModel) SetData(data []map[string]any, query url.Values) {
	list.Results.Data = data
}

func (list *ListModel) SetLink(c *fiber.Ctx) {
	q := Query().Parse(c.OriginalURL())
	q.Set(grest.QueryLimit, strconv.Itoa(int(list.Results.PageContext.PerPage)))

	path, _, _ := strings.Cut(c.OriginalURL(), "?")

	first := q
	first.Del(grest.QueryPage)
	first.Add(grest.QueryPage, "1")
	firstQS, _ := url.QueryUnescape(first.Encode())
	list.Results.PageContext.Links.First = c.BaseURL() + path + firstQS

	if list.Results.PageContext.Page > 1 && list.Results.PageContext.PageCount > 1 {
		previous := q
		previous.Set(grest.QueryPage, strconv.Itoa(int(list.Results.PageContext.Page-1)))
		previousQS, _ := url.QueryUnescape(previous.Encode())
		list.Results.PageContext.Links.Previous = c.BaseURL() + path + previousQS
	}

	if list.Results.PageContext.Page < list.Results.PageContext.PageCount {
		next := q
		next.Set(grest.QueryPage, strconv.Itoa(int(list.Results.PageContext.Page+1)))
		nextQS, _ := url.QueryUnescape(next.Encode())
		list.Results.PageContext.Links.Next = c.BaseURL() + path + nextQS
	}

	last := q
	last.Set(grest.QueryPage, strconv.Itoa(int(list.Results.PageContext.PageCount)))
	lastQS, _ := url.QueryUnescape(last.Encode())
	list.Results.PageContext.Links.Last = c.BaseURL() + path + lastQS

	list.Status = "Success"
	list.Message = "Data retrivied successfully"
	list.TimeStamp = time.Now()
}
func (list *ListModel) SetOpenAPISchema(m ModelInterface) map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"status":    map[string]any{"type": "string"},
			"message":   map[string]any{"type": "string"},
			"timestamp": map[string]any{"type": "string"},
			"results": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"list": map[string]any{
						"type":  "array",
						"items": m.GetOpenAPISchema(),
					},
					"pagination": map[string]any{"type": "object", "properties": map[string]any{
						"page":        map[string]any{"type": "integer"},
						"limit":       map[string]any{"type": "integer"},
						"total":       map[string]any{"type": "integer"},
						"total_pages": map[string]any{"type": "integer"},
						"links": map[string]any{"type": "object", "properties": map[string]any{
							"first":    map[string]any{"type": "string"},
							"previous": map[string]any{"type": "string"},
							"next":     map[string]any{"type": "string"},
							"last":     map[string]any{"type": "string"},
						}},
					}},
				},
			},
		},
	}
}

type Setting struct {
	Key   string `gorm:"column:key;primaryKey"`
	Value string `gorm:"column:value"`
}

func (Setting) TableName() string {
	return "settings"
}

func (Setting) KeyField() string {
	return "key"
}

func (Setting) ValueField() string {
	return "value"
}

func (Setting) MigrationKey() string {
	return "table_versions"
}

func (Setting) SeedKey() string {
	return "executed_seeds"
}
