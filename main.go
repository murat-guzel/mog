package main

import(

	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)


var (
    mgoSession *mgo.Session
    databaseName  = ""
    collection = ""
)

type Person struct {
	Id  bson.ObjectId `json:"_id" bson:"_id"`
	Name string `json:"Name" bson:"Name"`
	Surname string  `json:"Surname" bson:"Surname"`

}

func main(){

	BuildMongo("Collection","DatabaseName")
	person := NewSomething()
	//Insert(person)
	//UpdateById(person,bson.ObjectIdHex("5a86bb4ccb3652435837d604"))
	//UpdateByProp(person,"Name","Murat")
	//Delete(person,bson.ObjectIdHex("5a86bb4ccb3652435837d604"))
	DeleteByProp(person,"Name","Murat")
	fmt.Println("Success")
}

//Bootstrap Mongo
func BuildMongo(_collection string,_databaseName string){

    collection = _collection
    databaseName =_databaseName

}

//Create a struct
func NewSomething() Person {
   something := Person{}
   something.Name = "UpdatedMurat"
   something.Surname = "UpdatedGuzel"
   return something
}

//Get Session
func getSession () *mgo.Session {
    if mgoSession == nil {
        var err error
        mgoSession, err = mgo.Dial("127.0.0.1")
        if err != nil {
             panic(err) // no, not really
        }
    }
    return mgoSession.Clone()
}

//Create Connection
func withCollection(s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(databaseName).C(collection)
    return s(c)
}

//Insert Struct
func Insert (_person Person) (IsSuccess bool) {
  query := func(c *mgo.Collection) error {
	  err := c.Insert(&Person{Id:bson.NewObjectId(),Name: _person.Name, Surname:  _person.Surname})

	  return err
    }
    insert := func() error {
        return withCollection(query)
    }
    err := insert()
    if err != nil {
         panic(err)
    }

	return
}

//UpdateById struct
func UpdateById (_person Person,_structId bson.ObjectId) (IsSuccess bool) {
  query := func(c *mgo.Collection) error {

        colQuerier := bson.M{"_id": _structId}

	change := bson.M{"$set": bson.M{"Name": _person.Name, "Surname": _person.Surname}}
	err := c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}



	return err
    }
    insert := func() error {
        return withCollection(query)
    }
    err := insert()
    if err != nil {
         panic(err)
    }

	return
}

//UpdateByProp struct
func UpdateByProp (_person Person,_prop string,_propValue string) (IsSuccess bool) {
  query := func(c *mgo.Collection) error {

        colQuerier := bson.M{_prop :_propValue}

	change := bson.M{"$set": bson.M{"Name": _person.Name, "Surname": _person.Surname}}
	err := c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}



	return err
    }
    insert := func() error {
        return withCollection(query)
    }
    err := insert()
    if err != nil {
         panic(err)
    }

	return
}
//DeleteById Struct
func DeleteById (_person Person,_structId bson.ObjectId) (IsSuccess bool) {
  query := func(c *mgo.Collection) error {

	colQuerier := bson.M{"_id": _structId}
	err := c.Remove(colQuerier)
	if err != nil {
		panic(err)
	}

	return err
    }
    insert := func() error {
        return withCollection(query)
    }
    err := insert()
    if err != nil {
         panic(err)
    }

	return
}

//DeleteByProp Struct
func DeleteByProp (_person Person,_prop string,_propValue string) (IsSuccess bool) {
  query := func(c *mgo.Collection) error {

	colQuerier := bson.M{_prop :_propValue}
	err := c.Remove(colQuerier)
	if err != nil {
		panic(err)
	}

	return err
    }
    insert := func() error {
        return withCollection(query)
    }
    err := insert()
    if err != nil {
         panic(err)
    }

	return
}

