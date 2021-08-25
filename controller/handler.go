package controller

import (
	"html/template"
	"net/http"

	m "opus/model"
)

type musician struct {
	Name       string
	Instrument string
	Text       template.HTML
	Quote      string
	From       string
	Path       string
}

var melanie = musician{
	Name:       "Melanie Reinhard",
	Instrument: "Piano",
	Text:       "The Canadian pianist, who has many prizes and awards to her credit. Having performed extensively on both sides of the Atlantic, she has been recognised as a versatile and artistic musician who is sought after by noted instrumentalists and singers. Melanie received her ARCT (Performance) and LPRCM (Concert Diploma) from the Royal Conservatory of Music, Toronto, studying with James Anagnoson and Leon Fleisher. In master classes, she has worked with John Perry, Richard Goode, Marek Jablonski and Monika Leonhard.<br><br>Having lived and worked in Germany and Switzerland for a decade, performing in many places including Frankfurt, Munich, Konstanz, Zürich and the Gewandhaus Leipzig, where her concert was broadcast on Mitteldeutscher Rundfunk, Melanie moved to the UK in 2004 and is active here as a teacher and performer. As the pianist of the Opus 3 piano trio, she has performed in venues such as St. James’ Picadilly, Charterhouse, St. George in the East and for the Beethoven Society of Europe to critical acclaim.",
	Quote:      "delicately and subtly Melanie Reinhard accompanied on the piano, so salient were the creative impulses she chose when she took over the lead melody onto the piano",
	From:       "Badische Zeitung",
	Path:       "img/melanie.jpg",
}

var chris = musician{
	Name:       "Christopher White",
	Instrument: "Violin",
	Text:       "The London born violinist studied at the Royal College of Music, London, where he won many prizes and left with one of the most coveted awards. Further awards enabled him to continue his studies at the Banff Centre in Canada and the Royal Conservatory of Music, Toronto. He has studied with many distinguished violinists both in North America and England, including Josef Gingold, Jaime Laredo, Thomas Brandis, Rodney Friend and Erich Gruenberg.<br><br> Christopher has been concertmaster with a number of orchestras both in England and Europe, alongside a varied performing career as soloist and chamber musician. His solo CD was released in 1999, featuring solo violin works by Bach, Bartok and Ysaye. In 2004 he was appointed Head of Strings at Uppingham School. In 2010, Christopher left his post at Uppingham to pursue a more varied career and alongside performing, teaches at Rugby and Loughborough Endowed Schools.<br><br>Christopher plays on a Giuseppe Guarneri violin, dated 1707.",
	Quote:      "Mature perfect touch and bow technique - always in the service of high musicality and sensitive tone production",
	From:       "Südkurier",
	Path:       "img/chris.jpg",
}

var jonas = musician{
	Name:       "Jonas Seeberg",
	Instrument: "Cello",
	Text:       "Born in Marburg, Germany, Jonas Seeberg began his study at the age of sixteen in Munich with the Romanian Horatiu Cenariu, who influenced him significantly. Six year later he travelled to the United States to continue his studies with Bernard Greenhouse, the cellist and co-founder of the legendary Beaux-Arts-Trio. This period was particularly formative for his artistic development. Greenhouse shared with him an ideal of cello playing and musicianship - almost forgotten today - with roots in the legacy of Pablo Casals and Emanuel Feuermann.<br><br>Once back in Germany, Seeberg completed his soloist exam at the Berliner Musikhochschule and embarked on a career as a principal cellist in various European orchestras – in Zürich, in Meiningen, and since 2001 in Maastricht. In the last few years, he has been increasingly appreciated and sought after as a soloist. The Six Suites for solo cello by Johann Sebastian Bach are an essential component of his repertoire – a repertoire that also extends to all epochs of music history. These are the starting point of his signature both as a musician and cellist, and a daily source of inspiration.<br><br> Jonas plays a cello made by Matteo Gofriller (Venice, 1698) and a bow by Francois Tourte (Paris, 1805).",
	Quote:      "Seeberg exhibits total effortlessness and deep emotion, playing with great virtuosity, but also playfully, with crystal clarity and sure intonation.",
	From:       "Rheinische Post",
	Path:       "img/jonas.jpg",
}

// HomeHandler :
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates := m.GetTemplates("./view/home.html")
	t := template.Must(template.ParseFiles(templates...))
	t.Execute(w, nil)
}

// MusicianHandler :
func MusicianHandler(w http.ResponseWriter, r *http.Request) {
	templates := m.GetTemplates("./view/musician.html")
	t := template.Must(template.ParseFiles(templates...))
	name := r.URL.Query().Get("name")
	if name == "melanie" {
		t.Execute(w, melanie)
	} else if name == "chris" {
		t.Execute(w, chris)

	} else if name == "jonas" {
		t.Execute(w, jonas)

	} else {
		t.Execute(w, nil)
	}
}
