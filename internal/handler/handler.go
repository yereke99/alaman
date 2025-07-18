package handler

import (
	"alaman/config"
	"alaman/internal/repository"
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

const (
	stateStart      string = "start"
	stateCount      string = "count"
	statePaid       string = "paid"
	stateContact    string = "contact"
	stateAdminPanel string = "admin_panel"
	stateBroadcast  string = "broadcast"
)

type Handler struct {
	cfg       *config.Config
	logger    *zap.Logger
	ctx       context.Context
	repo      *repository.UserRepository
	redisRepo *repository.RedisRepository
	bot       *bot.Bot // Add bot instance to handler
}

func NewHandler(cfg *config.Config, zapLogger *zap.Logger, ctx context.Context, repo *repository.UserRepository, redisRepo *repository.RedisRepository) *Handler {
	rand.Seed(time.Now().UnixNano())
	handle := &Handler{
		cfg:       cfg,
		logger:    zapLogger,
		ctx:       ctx,
		repo:      repo,
		redisRepo: redisRepo,
	}

	return handle
}

func (h *Handler) SetBot(b *bot.Bot) {
	h.bot = b
}

func (h *Handler) StartWebServer(ctx context.Context, b *bot.Bot) {
	h.SetBot(b)
}

// setCORSHeaders sets CORS headers for HTTP responses
func (h *Handler) setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func (h *Handler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

}
