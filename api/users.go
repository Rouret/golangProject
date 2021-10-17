package api

import (
    //"github.com/gin-gonic/gin"
)

type Users struct {
    Id   int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Name string `gorm:"not null" form:"name" json:"name"`
}