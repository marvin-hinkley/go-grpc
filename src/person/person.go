package person

import (
	"log"

	"golang.org/x/net/context"
)

//Person represents a person.
type Person struct {
	name      string
	eyeColor  string
	hairColor string
	height    uint32
}

//Server is a tie-in point for grpc server
type Server struct {
}

var pete = Person{"Pete", "hazel", "brown", 82}

//Speak returns a string stating name and hair color.
func (s *Server) Speak(ctx context.Context, void *Void) (*SpeakMessage, error) {
	log.Println("Got a request for Pete to speak.")
	return &SpeakMessage{"Hi. I'm " + pete.name + ". My hair color is " + pete.hairColor + "."}, nil
}

//DyeHair changes hair color.
func (s *Server) DyeHair(ctx context.Context, color *ColorMessage) (*Void, error) {
	log.Println("Got a request to dye Pete's hair.")
	pete.hairColor = color.Color
	return &Void{}, nil
}

//GetPerson returns the internally defined person
func (s *Server) GetPerson(ctx context.Context, void *Void) (*PersonMessage, error) {
	log.Println("Got a request to send Pete along.")
	return &PersonMessage{pete.name, pete.eyeColor, pete.hairColor, pete.height}, nil
}
