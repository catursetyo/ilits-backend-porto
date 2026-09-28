package models

import "time"

type Portfolio struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	TechStack   string    `gorm:"type:varchar(255);not null" json:"tech_stack"`
	ProjectURL  string    `gorm:"type:varchar(255)" json:"project_url"`
	Status      string    `gorm:"type:varchar(50);default:'Completed'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreatePortfolioInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	TechStack   string `json:"tech_stack" binding:"required"`
	ProjectURL  string `json:"project_url" binding:"omitempty,url"`
	Status      string `json:"status" binding:"omitempty,oneof=Planned 'In Progress' Completed"`
}

type UpdatePortfolioInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	TechStack   string `json:"tech_stack"`
	ProjectURL  string `json:"project_url" binding:"omitempty,url"`
	Status      string `json:"status" binding:"omitempty,oneof=Planned 'In Progress' Completed"`
}

type ResponseFormat struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
