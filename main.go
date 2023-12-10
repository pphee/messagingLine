package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type TextEvent struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

type StickerEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64  `json:"timestamp"`
		Mode      string `json:"mode"`
		Message   struct {
			Type                string `json:"type"`
			ID                  string `json:"id"`
			StickerID           string `json:"stickerId"`
			PackageID           string `json:"packageId"`
			StickerResourceType string `json:"stickerResourceType"`
		} `json:"message"`
	} `json:"events"`
	Destination string `json:"destination"`
}

type LocationEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64  `json:"timestamp"`
		Mode      string `json:"mode"`
		Message   struct {
			Type      string  `json:"type"`
			ID        string  `json:"id"`
			Address   string  `json:"address"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"message"`
	} `json:"events"`
	Destination string `json:"destination"`
}

type ImageEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64  `json:"timestamp"`
		Mode      string `json:"mode"`
		Message   struct {
			Type            string `json:"type"`
			ID              string `json:"id"`
			ContentProvider struct {
				Type string `json:"type"`
			} `json:"contentProvider"`
		} `json:"message"`
	} `json:"events"`
	Destination string `json:"destination"`
}

type AudioEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64  `json:"timestamp"`
		Mode      string `json:"mode"`
		Message   struct {
			Type            string `json:"type"`
			ID              string `json:"id"`
			ContentProvider struct {
				Type string `json:"type"`
			} `json:"contentProvider"`
			Duration int `json:"duration"`
		} `json:"message"`
	} `json:"events"`
	Destination string `json:"destination"`
}

type VideoEvent struct {
	Events []struct {
		Type       string `json:"type"`
		ReplyToken string `json:"replyToken"`
		Source     struct {
			UserID string `json:"userId"`
			Type   string `json:"type"`
		} `json:"source"`
		Timestamp int64  `json:"timestamp"`
		Mode      string `json:"mode"`
		Message   struct {
			Type            string `json:"type"`
			ID              string `json:"id"`
			ContentProvider struct {
				Type string `json:"type"`
			}
			Duration int `json:"duration"`
		} `json:"message"`
	} `json:"events"`
	Destination string `json:"destination"`
}

type TextMessage struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type StickerMessage struct {
	ID        string `json:"id"`
	StickerID string `json:"stickerId"`
	PackageID string `json:"packageId"`
}

type LocationMessage struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ImageMessage struct {
	ID string `json:"id"`
}

type AudioMessage struct {
	ID       string `json:"id"`
	Duration int    `json:"duration"`
}

type VideoMessage struct {
	Type               string `json:"type"`
	OriginalContentUrl string `json:"originalContentUrl"`
	PreviewImageUrl    string `json:"previewImageUrl"`
	TrackingId         string `json:"trackingId"`
}

type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}

type ReplyMessageSticker struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

type ReplyMessageLocation struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

type ReplyMessageImage struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

type ReplyAudio struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

type ReplyMessageVideo struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

// Text struct represents the text message structure
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type MessageRequest struct {
	To       string `json:"to"`
	Messages []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"messages"`
}

func replyMessageLine(message ReplyMessage) error {
	fmt.Println("message", message)
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response Status: %s, Body: %s", resp.Status, string(body))
	return nil
}

func replyStickerLine(message ReplyMessageSticker) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

func replyLocationLine(message ReplyMessageLocation) error {
	fmt.Println("message", message)
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

func replyImageLine(message ReplyMessageImage) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

func replyAudioLine(message ReplyAudio) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

func replyVideoLine(message ReplyMessageVideo) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	url := "https://api.line.me/v2/bot/message/reply"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

func sendMessageLine(c *gin.Context) {
	var msgRequest MessageRequest
	if err := c.ShouldBindJSON(&msgRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("msgRequest", msgRequest)

	if err := sendToLineAPI("https://api.line.me/v2/bot/message/push", msgRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendToLineAPI(url string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("ACCESS_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response from LINE API: %s, Body: %s", resp.Status, string(body))
	return nil
}

// handle
func handleTextMessage(message TextMessage, replyToken string) {
	fmt.Printf("Received text message: %s, Message ID: %s\n", message.Text, message.ID)

	replyText := "You said: " + message.Text
	reply := ReplyMessage{
		ReplyToken: replyToken,
		Messages: []Text{
			{
				Type: "text",
				Text: replyText,
			},
		},
	}

	if err := replyMessageLine(reply); err != nil {
		log.Printf("Error in replying to message: %v\n", err)
	}
}

func handleStickerMessage(message StickerMessage, replyToken string) {
	fmt.Printf("Received sticker message: %+v\n", message)

	reply := ReplyMessageSticker{
		ReplyToken: replyToken,
		Messages: []interface{}{
			map[string]interface{}{
				"type":      "sticker",
				"packageId": message.PackageID,
				"stickerId": message.StickerID,
			},
		},
	}

	if err := replyStickerLine(reply); err != nil {
		log.Printf("Error in replying to message: %v\n", err)
	}
}

func handleLocationMessage(message LocationMessage, replyToken string) {
	fmt.Printf("Received location message: %+v\n", message)

	// Truncate the address if it's longer than 100 characters
	address := message.Address
	if len(address) > 100 {
		address = address[:100]
	}

	// Provide a default title if the title is empty
	title := message.Title
	if title == "" {
		title = "Location"
	}

	reply := ReplyMessageLocation{
		ReplyToken: replyToken,
		Messages: []interface{}{
			map[string]interface{}{
				"type":      "location",
				"title":     title,
				"address":   address,
				"latitude":  message.Latitude,
				"longitude": message.Longitude,
			},
		},
	}

	if err := replyLocationLine(reply); err != nil {
		log.Printf("Error in replying to message: %v\n", err)
	}
}

func handleImageMessage(message ImageMessage, replyToken string) {
	fmt.Printf("Received image message: %+v\n", message)

	originalContentUrl := "https://upload.wikimedia.org/wikipedia/en/a/a9/Example.jpg"
	previewImageUrl := "https://upload.wikimedia.org/wikipedia/commons/3/3a/Cat03.jpg"

	reply := ReplyMessageImage{
		ReplyToken: replyToken,
		Messages: []interface{}{
			map[string]interface{}{
				"type":               "image",
				"originalContentUrl": originalContentUrl,
				"previewImageUrl":    previewImageUrl,
			},
		},
	}

	if err := replyImageLine(reply); err != nil {
		log.Printf("Error in replying to message: %v\n", err)
	}
}

func handleAudioMessage(message AudioMessage, replyToken string) {
	fmt.Printf("Received image message: %+v\n", message)

	reply := ReplyAudio{
		ReplyToken: replyToken,
		Messages: []interface{}{
			map[string]interface{}{
				"type":               "audio",
				"originalContentUrl": "https://drive.google.com/file/d/1w-0-_hJdDKwiLrXWFoqyelKwHE932uBJ/view",
				"duration":           240000,
			},
		},
	}

	if err := replyAudioLine(reply); err != nil {
		log.Printf("Error in replying to message: %v\n", err)
	}
}

func handleVideoReply(replyToken string) {
	videoReply := ReplyMessageVideo{
		ReplyToken: replyToken,
		Messages: []interface{}{
			VideoMessage{
				Type:               "video",
				OriginalContentUrl: "https://storage.googleapis.com/gtv-videos-bucket/sample/ForBiggerEscapes.mp4",
				PreviewImageUrl:    "https://storage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
				TrackingId:         "track-id",
			},
		},
	}

	if err := replyVideoLine(videoReply); err != nil {
		log.Printf("Error replying with video: %v\n", err)
	}
}

func mapEventMessage(source interface{}, target interface{}) error {
	data, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

// send
func sendMessage(c *gin.Context) {
	var line TextEvent

	if err := c.BindJSON(&line); err != nil {
		log.Printf("Error in binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in request"})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", line)

	if len(line.Events) == 0 {
		log.Println("No events to process")
		c.JSON(http.StatusOK, gin.H{"message": "No events to process"})
		return
	}

	for _, event := range line.Events {
		if event.Type == "message" {
			if event.Message.Type == "text" {
				var message TextMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleTextMessage(message, event.ReplyToken)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func sendStickerMessage(c *gin.Context) {
	var Sticker StickerEvent
	if err := c.ShouldBindJSON(&Sticker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", Sticker)

	for _, event := range Sticker.Events {
		if event.ReplyToken == "" {
			log.Println("Invalid or empty reply token")
			continue
		}
		if event.Type == "message" {
			if event.Message.Type == "sticker" {
				var message StickerMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleStickerMessage(message, event.ReplyToken)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendLocationMessage(c *gin.Context) {
	var Location LocationEvent
	if err := c.ShouldBindJSON(&Location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", Location)

	for _, event := range Location.Events {
		if event.Type == "message" {
			if event.Message.Type == "location" {
				var message LocationMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleLocationMessage(message, event.ReplyToken)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendImageMessage(c *gin.Context) {
	var Image ImageEvent
	if err := c.ShouldBindJSON(&Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", Image)

	for _, event := range Image.Events {
		if event.Type == "message" {
			if event.Message.Type == "image" {
				var message ImageMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleImageMessage(message, event.ReplyToken)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendAudioMessage(c *gin.Context) {
	var Audio AudioEvent
	if err := c.ShouldBindJSON(&Audio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", Audio)

	for _, event := range Audio.Events {
		if event.Type == "message" {
			if event.Message.Type == "audio" {
				var message AudioMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleAudioMessage(message, event.ReplyToken)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendVideoMessage(c *gin.Context) {
	var Video VideoEvent
	if err := c.ShouldBindJSON(&Video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("------------------------Request------------------------ %+v\n", Video)

	for _, event := range Video.Events {
		if event.Type == "message" {
			if event.Message.Type == "video" {
				var message VideoMessage
				if err := mapEventMessage(event.Message, &message); err != nil {
					log.Printf("Error in mapping event message: %v\n", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request"})
					return
				}
				handleVideoReply(event.ReplyToken) // Call handleVideoReply with only the reply token
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Video message processed successfully"})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running")
	})

	r.POST("/message", sendMessage)
	r.POST("/sticker", sendStickerMessage)
	r.POST("/location", sendLocationMessage)
	r.POST("/image", sendImageMessage)
	r.POST("/audio", sendAudioMessage)
	r.POST("/video", sendVideoMessage)

	r.POST("/send", sendMessageLine)

	log.Fatal(r.Run(":1323")) // Start the server on port 1323
}
