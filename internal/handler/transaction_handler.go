package handler

import (
	"encoding/json"
	"net/http"
	"payment-system/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type TransactionHandler struct {
	Repo  *repository.TransactionRepository
	Queue *asynq.Client
}

func NewTransactionHandler(repo *repository.TransactionRepository, queue *asynq.Client) *TransactionHandler {
	return &TransactionHandler{
		Repo:  repo,
		Queue: queue,
	}
}

type CreateTransactionRequest struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txID, err := h.Repo.Create(req.UserID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction"})
		return
	}

	// payload
	payload, _ := json.Marshal(map[string]int{
		"transaction_id": txID,
	})

	// push ke queue
	task := asynq.NewTask("process:transaction", payload)
	_, err = h.Queue.Enqueue(task,
		asynq.MaxRetry(5),             // 🔥 retry max 5x
		asynq.Timeout(10*time.Second), // timeout per job
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue task"})
		return
	}

}

func (h *TransactionHandler) List(c *gin.Context) {
	data, err := h.Repo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch transactions"})
		return
	}

	c.JSON(200, data)
}
