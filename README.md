# golang-decode-string-json-to-int
Cannot unmarshal string into Go value of type int

An int type alias that will allow you to send int fields as string 
This will prevent you from getting the error `json: cannot unmarshal string into Go struct field Item.item_id of type int` when you try 
to convert the json into a struct.