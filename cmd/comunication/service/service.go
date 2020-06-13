package serviceInterface

import (
  "context"
  "dunnorobot/cmd/comunication/dto"
  "fmt"
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
  "github.com/nu7hatch/gouuid"
)


func PrimaryService(mObj dtoInterface.MessageStruct) string {
  fmt.Println(mObj.Message)
  ctx := context.Background()
  sessionClient, err := dialogflow.NewSessionsClient(ctx)
  if err != nil {
    fmt.Println(err)
    return "erro"
  }
  defer sessionClient.Close()
  var projectId string
  var sessionId string
  var language string
  uuidGen, _ := uuid.NewV4()
  projectId = "small-talk-svaxoa"
  sessionId = uuidGen.String()
  language = "pt-br"
  sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectId, sessionId)
  
  textInput := dialogflowpb.TextInput{
    Text: mObj.Message,
    LanguageCode: language,
  }
  queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
  queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
  request := dialogflowpb.DetectIntentRequest{Session: sessionPath,
    QueryInput: &queryInput}
  response, err := sessionClient.DetectIntent(ctx, &request)
  if err != nil {
    return "erro"
  }  
  queryResult := response.GetQueryResult()
  fufillmentText := queryResult.GetFulfillmentText()
  return fufillmentText
}
