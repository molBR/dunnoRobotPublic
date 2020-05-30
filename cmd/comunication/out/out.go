package outInterface

import ( 

  "fmt"
  "dunnorobot/cmd/comunication/dto"
  "net/http"
)

func PrimaryOut(mObj dtoInterface.MessageStruct) {
  
  fmt.Println(mObj.Message)
  reqBody, err ?= json.Marshal(map[string]string{
    "msg": mObj.Message
  })
  if err != nil {
    fmt.Println("ERRO!")
  }
  
}
