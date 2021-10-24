package user

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"service-tutorial/shared/proto"
	"time"
)

type UserRegistration struct {
	Email string
}

func CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	var reg UserRegistration
	err := json.NewDecoder(r.Body).Decode(&reg)
	fmt.Println("registering email for", reg.Email)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("notifications:4040", opts...)
	if err != nil {
		log.Fatalln("fail to dial: %w", err)
	}
	defer conn.Close()
	client := proto.NewNotificationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SendRegistrationEmail(ctx, &proto.RegistrationEmailRequest{
		Email: reg.Email,
	})
	if !resp.Success {
		log.Fatalf("error sending email notification: %s", resp.Error)
	} else {
		log.Printf("successfully sent email user email notification for: %s", reg.Email)
	}
}
