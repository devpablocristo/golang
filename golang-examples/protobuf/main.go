package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	p := &Person{
		Name: "Clark",
		Age:  35,
		SocialFollowers: &SocialFollowers{
			Ig: 15000,
			Fb: 230000,
		},
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newP := &Person{}
	err = proto.Unmarshal(data, newP)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newP.GetName())
	fmt.Println(newP.GetAge())
	fmt.Println(newP.GetSocialFollowers())

}
