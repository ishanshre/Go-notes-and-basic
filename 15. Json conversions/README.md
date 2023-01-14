# Json Conversion

## Converting JSOn string to go struct or objects
- We use json.unmarshal method to convert into go objects
- json.unmarshal comes from 'encoding/json' package


```
jsonStr = `{"NAME":"ishan","LANG":"GO"}`
err := json.Unmarshal([]byte(jsonStr), &shrestha)
if err != nil {
	fmt.Println(err)
}
```

## Converting Go struct or objects to json
- We use json.marshal method to convert go objects to json strings

```
binod := Warrior{
	Name:   "Binod",
	Weapon: "Knife",
	Level:  99,
}
convertToJson, err := json.Marshal(binod)
if err != nil {
	fmt.Println(err)
}
fmt.Println(convertToJson)
fmt.Printf("%s:\n ", convertToJson)
fmt.Printf("%T : \n", convertToJson)
```