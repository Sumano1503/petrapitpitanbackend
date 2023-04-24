package notificationcontroller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func PushNotificationUser(c *gin.Context){
	//get data dari request body(external id peminjam, dan tipe notifikasi)
	var reqBody struct {
		ExternalIDs []string `json:"external_ids"`
		Message     string   `json:"message"`
		Headings   string   `json:"headings"`
		Subtitle  string   `json:"subtitle"`
	}

	// Parsing request body untuk mendapatkan external ID dari user yang ingin diberi notifikasi push
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat payload untuk notifikasi push
	payload, err := json.Marshal(map[string]interface{}{
		"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
		"include_external_user_ids": reqBody.ExternalIDs,
		"contents":                map[string]string{"en": reqBody.Message},
		"headings":                map[string]string{"en": reqBody.Headings},
		"subtitle":                map[string]string{"en": reqBody.Subtitle},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create push notification payload"})
		return
	}

	// Mengirim notifikasi push ke OneSignal
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create HTTP request for push notification"})
		return
	}
	req.Header.Add("Authorization", "Basic ZDFhOGQ4MmYtNzZmMy00Zjg1LTkyYzctNmIyNDI0MGJhNjU0")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send push notification"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read push notification response"})
			return
		}

		// Parsing response dari OneSignal
	var respBody map[string]interface{}
	if err := json.Unmarshal(body, &respBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse push notification response"})
		return
	}

	// Cek apakah ada error dari OneSignal
	if _, ok := respBody["errors"]; ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": respBody["errors"]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "push notification sent"})
}

func PushNotificationAdmin(c *gin.Context){
	var users []models.User

	err := models.DB.Where("role = ?", "Admin").Find(&users).Error

	if err == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}

	var externalIDs []string
	for _, user := range users {
		externalIDs = append(externalIDs, "admn"+strconv.Itoa(int(user.Id)))
	}

	//get data dari request body(external id peminjam, dan tipe notifikasi)
	var reqBody struct {
		Message     string   `json:"message"`
		Headings   string   `json:"headings"`
		Subtitle  string   `json:"subtitle"`
	}

	// Parsing request body untuk mendapatkan external ID dari user yang ingin diberi notifikasi push
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat payload untuk notifikasi push
	payload, err := json.Marshal(map[string]interface{}{
		"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
		"include_external_user_ids": externalIDs,
		"contents":                map[string]string{"en": reqBody.Message},
		"headings":                map[string]string{"en": reqBody.Headings},
		"subtitle":                map[string]string{"en": reqBody.Subtitle},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create push notification payload"})
		return
	}

	// Mengirim notifikasi push ke OneSignal
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create HTTP request for push notification"})
		return
	}
	req.Header.Add("Authorization", "Basic ZDFhOGQ4MmYtNzZmMy00Zjg1LTkyYzctNmIyNDI0MGJhNjU0")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send push notification"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read push notification response"})
			return
		}

		// Parsing response dari OneSignal
	var respBody map[string]interface{}
	if err := json.Unmarshal(body, &respBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse push notification response"})
		return
	}

	// Cek apakah ada error dari OneSignal
	if _, ok := respBody["errors"]; ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": respBody["errors"]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "push notification sent"})
}