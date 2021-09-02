package articles

type Article struct {
	Aid  int `json:"aid"`
	Title string `json:"title"`
	Desc string
	Content string
	Type string
	Ctime int
	Mtime int
	IsValid int
}

