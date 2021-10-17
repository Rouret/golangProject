package api

import (
    //"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
)

func InitDb() *gorm.DB {
    // Ouverture du fichier
    db, err := gorm.Open("sqlite3", "./data.db")
    db.LogMode(true) //Affiche les requêtes effectuées

    // Création de la table, si elle n'existe pas
    if !db.HasTable(&Users{}) {
        db.CreateTable(&Users{})
        db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
    }

    // Erreur de chargement
    if err != nil {
        panic(err)
    }

    return db
}