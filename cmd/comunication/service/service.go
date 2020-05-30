package serviceInterface

import (
  "dunnorobot/cmd/comunication/out" 
  "dunnorobot/cmd/comunication/dto"
  
)


func PrimaryService(mObj dtoInterface.MessageStruct) {
  outInterface.PrimaryOut(mObj)
}
