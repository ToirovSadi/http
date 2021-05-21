package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) handleGetAllBanners(w http.ResponseWriter, r *http.Request) {
	items, err := s.bannersSvc.All(r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	err = writeJson(w, items)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleGetBannerById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	item, err := s.bannersSvc.ByID(r.Context(), id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = writeJson(w, item)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleSaveBanner(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	title := r.URL.Query().Get("title")
	content := r.URL.Query().Get("content")
	button := r.URL.Query().Get("button")
	link := r.URL.Query().Get("link")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	banner, err := s.bannersSvc.ByID(r.Context(), id)
	if err == nil {
		if title != "" {
			banner.Title = title
		}
		if content != "" {
			banner.Content = content
		}
		if button != "" {
			banner.Button = button
		}
		if link != "" {
			banner.Link = link
		}
	}
	updBanner, err := s.bannersSvc.Save(r.Context(), banner)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = writeJson(w, updBanner)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
func (s *Server) handleRemoveById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	delBanner, err := s.bannersSvc.RemoveByID(r.Context(), id)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = writeJson(w, delBanner)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
func writeJson(w http.ResponseWriter, item interface{}) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "appication/json")
	_, err = w.Write(data)
	return err
}
