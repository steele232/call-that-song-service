package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createSongsTableIfNotExists() error {

	createTableStr :=
		`CREATE TABLE IF NOT EXISTS songs ( id BIGSERIAL PRIMARY KEY, name VARCHAR NOT NULL, url VARCHAR NOT NULL,  originalViews INTEGER NOT NULL, latestViews INTEGER NOT NULL DEFAULT 0 );`

	if _, err := db.Exec(createTableStr); err != nil {
		return err
	}
	return nil
}

func getSongs(c *gin.Context) {

	err := createSongsTableIfNotExists()
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading ticks: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
	}
}

type InsertSongReq struct {
	auth          string
	name          string
	url           string
	originalViews int
}

func addSong(c *gin.Context) {

	// TODO GET STUFF FROM FORM...
	// var bodyBytes []byte
	// if context.Request().Body != nil {
	// 	bodyBytes, _ = ioutil.ReadAll(context.Request().Body)
	// }
	req := InsertSongReq{}
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error decoding json in req body: %q", err))
		return
	}

	songName := "Jonathan"

	err = createSongsTableIfNotExists()
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
	}

	_, err = db.Exec("INSERT INTO songs(name, url, originalviews) VALUES ($1, $2, $3);",
		req.name, req.url, req.originalViews)
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error inserting song: %q", err))
		return
	}
	c.String(
		http.StatusCreated,
		fmt.Sprintf("Success inserting song: %q", songName),
	)
}

func updateSong(c *gin.Context) {
	err := createSongsTableIfNotExists()
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading ticks: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
	}
}

func deleteSong(c *gin.Context) {
	err := createSongsTableIfNotExists()
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error creating database table: %q", err))
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error incrementing tick: %q", err))
		return
	}

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading ticks: %q", err))
		return
	}

	defer rows.Close()
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning ticks: %q", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
	}
}
