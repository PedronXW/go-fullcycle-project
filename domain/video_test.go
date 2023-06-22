package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideo(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "abc"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceID = "abc"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}
