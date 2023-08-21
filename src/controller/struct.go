package controller

type Short struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Slide  string `json:"slide"`
	Views  int    `json:"views"`
	Poster string `json:"poster"`
}

type Slide struct {
	Script   string `json:"script"`
	Content  string `json:"slide"`
	VoiceURL string `json:"voiceURL"`
}

type ReactionRequest struct {
	Reaction string `json:"reaction"`
}

type Reaction struct {
	Count   int  `json:"count"`
	Reacted bool `json:"reacted"`
}
type Reactions struct {
	Heart Reaction `json:"heart"`
	Good  Reaction `json:"good"`
	Smile Reaction `json:"smile"`
}

type Achievement struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Genre struct {
	Name       string `json:"name"`
	Color      string `json:"color"`
	DailyViews []int  `json:"dailyViews"`
}

type Report struct {
	Dates  []string `json:"dates"`
	Genres []Genre  `json:"genres"`
}
type Data struct {
	Achievements []Achievement `json:"achievements"`
	Report       Report        `json:"report"`
}

type Presentation struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Speaker   string    `json:"speaker"`
	Slides    []Slide   `json:"slides"`
	Tags      []string  `json:"tags"`
	Genre     string    `json:"genre"`
	Views     int       `json:"views"`
	Poster    string    `json:"poster"`
	CreatedAt string    `json:"createdAt"`
	Reactions Reactions `json:"reactions"`
}

type ShortOutPut struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Speaker   string    `json:"speaker"`
	Slides    []Slide   `json:"slides"`
	Tags      []string  `json:"tags"`
	Genre     string    `json:"genre"`
	Views     int       `json:"views"`
	Poster    string    `json:"poster"`
	CreatedAt string    `json:"createdAt"`
	Reactions Reactions `json:"reactions"`
}
