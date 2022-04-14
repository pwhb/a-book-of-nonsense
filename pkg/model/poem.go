package model

import (
	"context"
	"log"

	"github.com/pwhb/lear/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Poem struct {
	Id    int      `json:"id"`
	Lines []string `json:"lines"`
	Image string   `json:"image"`
}

var ctx = context.TODO()

var poemCol *mongo.Collection

func init() {
	c, err := config.InitializeClient()
	if err != nil {
		log.Panicln(err)
	}
	db := c.Database("a_book_of_nonsense")
	poemCol = db.Collection("poems")
}

// func GetPoems() []Poem {
// 	data, e := os.ReadFile("a book of nonsense.txt")
// 	if e != nil {
// 		log.Println(e)
// 	}
// 	dataStr := string(data)

// 	poemsArr := strings.Split(dataStr, "\n\n\n\n\n")
// 	poems := []Poem{}

// 	for idx, str := range poemsArr {
// 		if idx == 112 {
// 			break
// 		}
// 		// pstr := parsePoem(str)
// 		lines := strings.Split(str, "\n\n")
// 		id := idx + 1
// 		poem := Poem{Lines: lines, Id: id, Image: fmt.Sprintf("%d.jpg", id)}

// 		poems = append(poems, poem)

// 	}

// 	return poems
// }

// func CreatePoems() {
// 	poems := GetPoems()
// 	for i, p := range poems {
// 		_, e := poemCol.InsertOne(ctx, p)
// 		if e != nil {
// 			log.Println(e)
// 		}
// 		fmt.Printf("poem #%d created\n", i)
// 	}
// }

func GetAllPoems() ([]*Poem, error) {
	var poems []*Poem
	cur, err := poemCol.Find(ctx, bson.D{})
	if err != nil {
		return poems, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var p Poem
		err := cur.Decode(&p)
		if err != nil {
			return poems, err
		}
		poems = append(poems, &p)
	}

	return poems, nil
}

// func FixPoem() {
// 	// var poems []*Poem
// 	cur, err := poemCol.Find(ctx, bson.D{})
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 		var p Poem
// 		err := cur.Decode(&p)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		p.Lines[0] = strings.Replace(p.Lines[0], "\n", "", 1)

// 		// poems = append(poems, &p)
// 		filter := bson.M{"id": p.Id}
// 		update := bson.M{
// 			"id":    p.Id,
// 			"lines": p.Lines,
// 			"image": p.Image,
// 		}
// 		poemCol.UpdateOne(ctx, filter, bson.M{"$set": update})

// 		fmt.Println("poem", p.Id, "updated")
// 	}

// }

func GetOnePoem(id int) (*Poem, error) {
	var p *Poem
	filter := bson.M{"id": id}
	err := poemCol.FindOne(ctx, filter).Decode(&p)
	return p, err
}

// func GetRandomPoem() (*Poem, error) {
// 	var p *Poem

// }
