package service

import (
	"github.com/kajiLabTeam/hacku-2023-back/integrations"
	"github.com/kajiLabTeam/hacku-2023-back/model"
)

func CreateShort(a string, req model.Short) error {
	var short model.Short
	var slides []model.Slide
	t,_ := integrations.GetUserByID(a)
	s := req.Slides

	for i,v := range s{
	var slide model.Slide
	b,_ := getBinary(v.Script)
	fn,_ := makeMp3File(b,t.UID)
	rp,_ := integrations.Storage(fn)
	slide.SlideText = v.SlideText
	slide.Voice = rp
	slide.PageNumber = i
	slide.Script = v.Script
	slides = append(slides, slide)
	}

	short.UserID = t.UID
	short.Slides = slides
	short.Speaker = req.Speaker
	model.InsertShort(short)

	return nil
}