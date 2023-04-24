package notificationcontroller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PushNotification(c *gin.Context){
	//get data dari request body(external id peminjam, dan tipe notifikasi)
	var reqBody struct {
		ExternalIDs []string `json:"external_ids"`
		Tipe     int  `json:"tipe"`
	}

	// Parsing request body untuk mendapatkan external ID dari user yang ingin diberi notifikasi push
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat payload untuk notifikasi push
	if(reqBody.Tipe == 1){
		payload, err := json.Marshal(map[string]interface{}{
			"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
			"include_external_user_ids": reqBody.ExternalIDs,
			"contents":                "Anda Keluar Dari Area Bersepeda Harap Kembali Ke Area Bersepeda",
			"headings":                "Anda Keluar Dari Area Bersepeda!!!!!",
			"subtitle":                "Anda Akan Di Blokir Jika Tidak Segera Kembali Ke Area Bersepeda",
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
	}else{
		payload, err := json.Marshal(map[string]interface{}{
			"app_id":                  "59865fb1-ab37-4f3a-9f21-e41e33194070",
			"include_external_user_ids": reqBody.ExternalIDs,
			"contents":                "Waktu Bersepeda Anda Telah Habis Harap Mengembalikan Sepeda Di Halte Tujuan Anda",
			"headings":                "Waktu Habis!!!!!",
			"subtitle":                "Anda Akan Di Blokir Jika Tidak Segera Mengembalikan Sepeda Ke Halte Tujuan Anda",
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


}