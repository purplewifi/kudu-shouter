package kudushouter

type Payload struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	StatusText         string `json:"statusText"`
	AuthorEmail        string `json:"authorEmail"`
	Author             string `json:"author"`
	Message            string `json:"message"`
	Deployer           string `json:"deployer"`
	ReceivedTime       string `json:"receivedTime"`
	StartTime          string `json:"startTime"`
	EndTime            string `json:"endTime"`
	LastSuccessEndTime string `json:"lastSuccessEndTime"`
	Complete           bool   `json:"complete"`
	SiteName           string `json:"siteName"`
	HostName           string `json:"hostName"`
}
