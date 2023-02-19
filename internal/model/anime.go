package model

type Anime struct {
	TeamName  string
	AnimeName string
	Season    []Season
}

type Season struct {
	Number int
	part   []Category
}

type Category struct {
	Number  int
	FileDir string
	Suffix  string
}
