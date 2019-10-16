/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository_test

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"free5gc/lib/Nudr_DataRepository"
	"free5gc/lib/openapi/models"
)

// QuerySmfRegList - Retrieves the SMF registration list of a UE
func TestQuerySmfRegList(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.contextData.smfRegistrations")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "pduSessionId": 1})
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "pduSessionId": 2})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	{
		ueId := "imsi-0123456789"
		pduSessionIdInt := 1
		testData := models.SmfRegistration{
			SmfInstanceId: "SmfInstanceId_Test",
			PduSessionId:  1,
			SingleNssai: &models.Snssai{
				Sst: 1,
			},
			Dnn: "Dnn_Test",
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
		}
		insertTestData := toBsonM(testData)
		insertTestData["ueId"] = ueId
		insertTestData["pduSessionId"] = int32(pduSessionIdInt)
		collection.InsertOne(context.TODO(), insertTestData)
	}

	{
		ueId := "imsi-0123456789"
		pduSessionIdInt := 2
		testData := models.SmfRegistration{
			SmfInstanceId: "SmfInstanceId_Test2",
			PduSessionId:  2,
			SingleNssai: &models.Snssai{
				Sst: 1,
			},
			Dnn: "Dnn_Test",
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
		}
		insertTestData := toBsonM(testData)
		insertTestData["ueId"] = ueId
		insertTestData["pduSessionId"] = int32(pduSessionIdInt)
		collection.InsertOne(context.TODO(), insertTestData)
	}

	testData := []models.SmfRegistration{
		{
			SmfInstanceId: "SmfInstanceId_Test",
			PduSessionId:  1,
			SingleNssai: &models.Snssai{
				Sst: 1,
			},
			Dnn: "Dnn_Test",
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			}},
		{
			SmfInstanceId: "SmfInstanceId_Test2",
			PduSessionId:  2,
			SingleNssai: &models.Snssai{
				Sst: 1,
			},
			Dnn: "Dnn_Test",
			PlmnId: &models.PlmnId{
				Mcc: "208",
				Mnc: "93",
			},
		},
	}

	{
		// Check test data (Use RESTful GET)
		var querySmfRegListParamOpts Nudr_DataRepository.QuerySmfRegListParamOpts
		smfRegistrationArray, res, err := client.SMFRegistrationsCollectionApi.QuerySmfRegList(context.TODO(), ueId, &querySmfRegListParamOpts)
		if err != nil {
			log.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if reflect.DeepEqual(testData, smfRegistrationArray) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				smfRegistrationArray, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "pduSessionId": 1})
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "pduSessionId": 2})

	// TEST END
}