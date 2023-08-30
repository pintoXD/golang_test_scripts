package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type ThingSpeak struct {
	Api_key    string `json:"api_key"`
	Created_at string `json:"created_at,omitempty"`
	Field1     int    `json:"field1,omitempty"`
	Field2     int    `json:"field2,omitempty"`
	Field3     string `json:"field3,omitempty"`
	Field4     string `json:"field4,omitempty"`
	Field5     string `json:"field5,omitempty"`
	Latitude   string `json:"latitude,omitempty"`
	Longitude  string `json:"longitude,omitempty"`
}

func create_thingspeak_payload(temperature, humidity int) []byte {
	thingspeak_payload := ThingSpeak{
		Api_key: "my_api_key",
		Field1:  temperature,
		Field2:  humidity,
	}
	encodedThingSpeak, err := json.Marshal(thingspeak_payload)
	if err != nil {
		panic(err)
	}

	return encodedThingSpeak
}

func send_payload_to_thingspeak(payload []byte) (string, []byte) {
	thingSpeakClient := http.Client{}
	resp, err := thingSpeakClient.Post("https://api.thingspeak.com/update.json", "application/json", bytes.NewReader(payload))
	if err != nil {
		panic(err)
	}
	bytes_response, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return resp.Status, bytes_response

}

func main() {
	for int := 0; int < 10; int++ {
		random_temperature := rand.Intn(35-30) + 30
		random_humidity := rand.Intn(100-90) + 90

		payload := create_thingspeak_payload(random_temperature, random_humidity)
		status, bytes_response := send_payload_to_thingspeak(payload)

		fmt.Printf("Status code: %s\n", status)
		fmt.Println(string(bytes_response))

		time.Sleep(time.Second * 5)
	}
}
