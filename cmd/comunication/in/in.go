package inputInterface


import ( 
	"net/http"
  "dunnorobot/cmd/comunication/service"
  "encoding/json"
  "dunnorobot/cmd/comunication/dto"
)



func HelloWorld (w http.ResponseWriter, r *http.Request) {
  
  decoder := json.NewDecoder(r.Body)
  encoder := json.NewEncoder(w)
  var mObj dtoInterface.MessageStruct
  err := decoder.Decode(&mObj)
  if err != nil {
    panic(err)
  }
  serviceResponse:=serviceInterface.PrimaryService(mObj)
  var response dtoInterface.MessageResponseStruct
  response.Response = serviceResponse
  encoder.Encode(response)
  return 
  
}