package api

type Pageable struct {
	Page  int    `json:"page" form:"page"`
	Size  int    `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

func (p Pageable) GetSize() int {
	if p.Size == 0 {
		return 10
	}
	return p.Size
}

func (p Pageable) GetPage() int {
	if p.Page < 0 {
		return 0
	}
	return p.Page
}

func (p Pageable) GetOffset() int {
	return p.GetPage() * p.GetSize()
}

func (p Pageable) GetLimit() int {
	return p.GetSize()
}

type Page struct {
	Number           int         `json:"number"`
	NumberOfElements int         `json:"numberOfElements"`
	Pageable         Pageable    `json:"pageable"`
	Size             int         `json:"size"`
	TotalElements    int         `json:"totalElements"`
	TotalPages       int         `json:"totalPages"`
	Content          interface{} `json:"content"`
	First            bool        `json:"first"`
	Last             bool        `json:"last"`
	Empty            bool        `json:"empty"`
}

var emptyPage = Page{
	Empty:   true,
	First:   true,
	Last:    true,
	Content: make([]interface{}, 0),
}

func EmptyPage() Page {
	return emptyPage
}

func NewPage(page Pageable, size, dataSize int, data interface{}) interface{} {
	tp, d := div(size, page.GetSize())
	if d != 0 {
		tp = tp + 1
	}
	return Page{
		Number:           page.GetSize(),
		NumberOfElements: dataSize,
		Pageable:         page,
		Size:             page.GetSize(),
		TotalElements:    size,
		TotalPages:       tp,
		Last:             tp-1 == page.GetPage(),
		First:            page.GetPage() == 0,
		Empty:            dataSize == 0,
		Content:          data,
	}
}

func div(first, last int) (int, int) {
	return first / last, first % last
}
