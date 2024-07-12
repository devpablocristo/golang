map1 := map[string]bool{"Interview": true, "Bit": true}
map2 := make(map[string]bool)
for key, value := range map1 {
	map2[key] = value
}
From this code, we are iterating the contents of map1 and then adding the values to map2 to the corresponding key.

If we want to copy just the description and not the content of the map, we can again use the = operator as shown below:

map1 := map[string]bool{"Interview": true, "Bit": true}
map2 := map[string]bool{"Interview": true, "Questions": true}
map3 := map1
map1 = map2    //copy description
fmt.Println(map1, map2, map3)
The output of the below code would be:

map[Interview:true Questions:true] map[Interview:true Questions:true] map[Interview:true Bit:true]
