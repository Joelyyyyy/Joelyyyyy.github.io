package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type EmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Gmail SMTP settings
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "j.paredes.psa@gmail.com"
	smtpPassword := os.Getenv("EMAIL_PASSWORD") // Get password from environment variable

	if smtpPassword == "" {
		log.Printf("EMAIL_PASSWORD environment variable not set")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}

	// Set up authentication
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Compose email
	to := []string{smtpUsername}
	subject := "New Contact Form Message from " + req.Name
	body := "From: " + req.Name + "\n" +
		"Email: " + req.Email + "\n\n" +
		"Message:\n" + req.Message

	msg := []byte("To: " + smtpUsername + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"\r\n" +
		body)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUsername, to, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}

func main() {
	http.HandleFunc("/api/send-email", enableCORS(sendEmail))

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
} 