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
	LastFeBranch    string `json:"last_fe_branch"`
	LastBeBranch    string `json:"last_be_branch"`
}

func UpdateServerBuild(Project string, ServerId string, Name string, FeBranch string, BeBranch string) {
	_, err := DB.Query("UPDATE testing_status set status = 1, last_build_by = ?, last_build_on = ?, last_fe_branch = ?, last_be_branch = ? where project = ? and server_id = ?", Name, time.Now().Unix(), FeBranch, BeBranch, Project, ServerId)
	if err != nil {
		log.Print(err.Error())
	}
}
