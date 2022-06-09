package service

import (
	"github.com/stretchr/testify/require"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"testing"
)

func TestNew(t *testing.T) {
	new := New()
	res := &videoService{
		videos: []entity.Video{},
	}
	require.Equal(t, new, res)

}
func TestSave(t *testing.T) {
	person := entity.Person{
		Firstname: "Samet",
		Lastname:  "Avcı",
		Age:       21,
		Email:     "sametavc05@gmail.com",
	}
	video := entity.Video{
		Author:      person,
		Title:       "Samet Avcı , Yazılım Mühendisi Adayıdır",
		URL:         "www.linkedin.com/in/samet-avci/",
		Description: "Samet Avcı bir Yazılım Mühendisi Adayıdır. Fırat Universitesinde öğrenim görmektedsir"}
	service := videoService{}
	service.videos = append(service.videos, video)
	len := len(service.videos)
	require.Equal(t, service.videos[len-1], video)
}
