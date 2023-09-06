package models

import (
	"log"
	"time"
)

type TestingStatus struct {
	Id              string `json:"id"`
	Project         string `json:"project"`
	ServerId        string `json:"server_id"`
	LastBuildBy     string `json:"last_build_by"`
	LastBuildOn     int    `json:"last_build_on"`
	Status          bool   `json:"status"`
	StatusChangedBy string `json:"status_changed_by"`
	StatusChangedOn int    `json:"status_changed_on"`
}

func UpdateServerBuild(Project string, ServerId string, Name string) {
	_, err := DB.Query("UPDATE testing_status set status = 1, last_build_by = ?, last_build_on = ? where project = ? and server_id = ?", Name, time.Now().Unix(), Project, ServerId)
	if err != nil {
		log.Print(err.Error())
	}
}
