package api

import (
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/repository"
)

type Server struct {
	Repository *repository.RepoPg
}

func NewServer(repo *repository.RepoPg) *Server {
	return &Server{Repository: repo}
}
