package inputInterface


import ( 
	"net/http"
  "dunnorobot/cmd/comunication/service"
  "encoding/json"
  "dunnorobot/cmd/comunication/dto"
)



func HelloWorld (w http.ResponseWriter, r *http.Request)  {
  

  decoder := json.NewDecoder(r.Body)

  var mObj dtoInterface.MessageStruct
  err := decoder.Decode(&mObj)
  if err != nil {
    panic(err)
  }
  serviceInterface.PrimaryService(mObj)
}