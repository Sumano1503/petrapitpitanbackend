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

func PushNotification(c *gin.Context){
	//get id admin
	var users []models.User

	Admin := models.DB.Where("role = ?", "Admin").Find(&users)

	if Admin == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}

	var AdminId []string

	for _, user := range users {
		AdminId = append(AdminId, ("admn"+strconv.Itoa(int(user.Id))))
	}



	//get data dari request body(external id peminjam, dan tipe notifikasi)
	var reqBody struct {
		ExternalIDs string `json:"external_ids"`
		Tipe     int   `json:"tipe"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if(reqBody.Tipe == 1){
		payload, err := json.Marshal(map[string]interface{}{
			"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
			"include_external_user_ids": reqBody.ExternalIDs,
			"contents":                "Waktu Peminjaman Telah Habis Mohon Segera Mengembalikan Sepeda!!!!",
			"headings":                "Waktu Peminjaman Telah Habis",
			"subtitle":                "Anda Akan Dinonaktifkan Jika Tidak Segera Mengembalikan Sepeda",
		})
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create push notification payload"})
			return
		}
	
		client := &http.Client{}
			req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(payload))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create HTTP request for push notification"})
				return
			}
	
		req.Header.Add("Authorization", "Basic ZDFhOGQ4MmYtNzZmMy00Zjg1LTkyYzctNmIyNDI0MGJhNjU0")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
	
		resp, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send push notification client"})
				return
			}
			defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read push notification response"})
			return
		}
	
		var respBody map[string]interface{}
		if err := json.Unmarshal(body, &respBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse push notification response"})
			return
		}
	
		notificationID := respBody["id"].(string)
	
		c.JSON(http.StatusOK, gin.H{"message": "push notification sent", "notification_id": notificationID})
	}else{
		payload, err := json.Marshal(map[string]interface{}{
			"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
			"include_external_user_ids": reqBody.ExternalIDs,
			"contents":                "Harap Segera Kembali Ke Area Bersepeda",
			"headings":                "Anda Keluar Dari Area Bersepeda!!!!",
			"subtitle":                "Anda Akan Dinonaktifkan Jika Tidak Segera Kembali Ke Area Bersepeda",
		})
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create push notification payload"})
			return
		}
	
		client := &http.Client{}
			req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(payload))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create HTTP request for push notification"})
				return
			}
	
		req.Header.Add("Authorization", "Basic ZDFhOGQ4MmYtNzZmMy00Zjg1LTkyYzctNmIyNDI0MGJhNjU0")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
	
		resp, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send push notification client"})
				return
			}
			defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read push notification response"})
			return
		}
	
		var respBody map[string]interface{}
		if err := json.Unmarshal(body, &respBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse push notification response"})
			return
		}
	
		notificationID := respBody["id"].(string)
	
		c.JSON(http.StatusOK, gin.H{"message": "push notification sent", "notification_id": notificationID})
	}

	payload, err := json.Marshal(map[string]interface{}{
		"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
		"include_external_user_ids": AdminId,
		"contents":                "Telah terjadi pelanggaran yang dilakukan peminjam",
		"headings":                "Pelanggaran!!!!!",
		"subtitle":                "Mohon cek detail pelangaran",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create push notification payload"})
		return
	}

	client := &http.Client{}
		req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(payload))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create HTTP request for push notification"})
			return
		}

	req.Header.Add("Authorization", "Basic ZDFhOGQ4MmYtNzZmMy00Zjg1LTkyYzctNmIyNDI0MGJhNjU0")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send push notification client"})
			return
		}
		defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read push notification response"})
		return
	}

	var respBody map[string]interface{}
	if err := json.Unmarshal(body, &respBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse push notification response"})
		return
	}

	notificationID := respBody["id"].(string)

	c.JSON(http.StatusOK, gin.H{"message": "push notification sent", "notification_id": notificationID})
}