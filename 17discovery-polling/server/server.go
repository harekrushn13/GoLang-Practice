package server

import (
	linux_plugin "discoverypolling/linux-plugin"
	"discoverypolling/utils/logger"
	"discoverypolling/utils/models"
	"encoding/json"
	"errors"
	"github.com/pebbe/zmq4"
	"log"
)

func RunServer(address string) error {

	responder, err := zmq4.NewSocket(zmq4.REP)

	if err != nil {
		return err
	}
	defer responder.Close()

	if err = responder.Bind(address); err != nil {
		return err
	}

	for {
		request, err := responder.Recv(0)
		if err != nil {
			log.Println("Error receiving request:", err)
			continue
		}

		log.Println("Received request:", request)

		var reqPayload models.RequestPayload

		if err := json.Unmarshal([]byte(request), &reqPayload); err != nil {
			log.Println("Error unmarshalling request:", err)
			continue
		}

		if reqPayload.PluginEngine != "linux" {
			log.Println("Invalid plugin engine:", reqPayload.PluginEngine)
			sendErrorResponse(responder, "Invalid plugin engine: "+reqPayload.PluginEngine)
			continue
		}

		switch reqPayload.EventName {
		case "discovery":
			discoveryResult := linux_plugin.Discovery(reqPayload.Host, reqPayload.Port, reqPayload.Username, reqPayload.Password)

			response := models.Response{
				Status:        discoveryResult.Status,
				StatusMessage: discoveryResult.StatusMessage,
			}

			jsonResponse, err := json.Marshal(response)

			if err != nil {
				log.Println("Error marshalling response:", err)
			}

			_, err = responder.Send(string(jsonResponse), 0)

			if err != nil {
				logger.Error("Error sending discovery response:", err)
			}

		case "poller":
			pollerResponse := linux_plugin.CollectLinuxData(reqPayload.Username, reqPayload.Password, reqPayload.Host, reqPayload.Port)

			_, err := responder.Send(pollerResponse, 0)

			if err != nil {
				logger.Error("Error sending poller response:", err)
			}

		default:
			logger.Error("Unknown event name: ", errors.New(reqPayload.EventName))
			sendErrorResponse(responder, "Unknown event name: "+reqPayload.EventName)
		}
	}
}

func sendErrorResponse(responder *zmq4.Socket, msg string) {
	response := models.Response{
		Status:        false,
		StatusMessage: msg,
		Data:          nil,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Println("Error marshalling response:", err)
	}

	_, err = responder.Send(string(jsonResponse), 0)

	if err != nil {
		log.Println("Error sending response:", err)
	}
}
