package settings

type Settings struct {
	SiteTitle   string
	Description string
	Dashboard   int
	Frontend    int
	Categories  int
	Playlist    int
	Thumb       string
	Date        string
}

func Setup() Settings {
	settings := Settings{
		SiteTitle:   "ដំណឹង​ល្អ",
		Description: "description",
		Dashboard:   10,
		Frontend:    20,
		Categories:  20,
		Playlist:    20,
		Thumb:       "",
		Date:        "",
	}

	return settings
}
