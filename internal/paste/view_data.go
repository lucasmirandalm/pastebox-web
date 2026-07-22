package paste

type HomePageData struct {
	Title       string
	TotalPastes int
	Pastes      []Paste
}

type EditPageData struct {
	Title string
	Paste Paste
	Error string
}
