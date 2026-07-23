package paste

type HomePageData struct {
	Title         string
	TotalPastes   int
	Pastes        []Paste
	OnlyFavorites bool
	Search        string
	AllURL        string
	FavoritesURL  string
}

type EditPageData struct {
	Title string
	Paste Paste
	Error string
}

type NewPageData struct {
	Title string
	Paste Paste
	Error string
}

type PublicPageData struct {
	Title string
	Paste Paste
}
