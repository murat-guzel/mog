# Mog Lib For MongoDB - Overview
**mog** is a mongodb middleware library written for golang. You can perform _Insert-Update-Delete-Select_ operations via this library.
Once you have created your own basic structured object, you can use mog methods to do things without thinking about session and connection operations.


# Requirements

gopkg.in/mgo.v2
gopkg.in/mgo.v2/bson

# Documentation

## Usage

```      
BuildMongo("Collection","DatabaseName")
person := NewSomething()
Insert(person)
UpdateById(person,bson.ObjectIdHex("5a86bb4ccb3652435837d604"))
UpdateByProp(person,"Name","Murat")
Delete(person,bson.ObjectIdHex("5a86bb4ccb3652435837d604"))
DeleteByProp(person,"Name","Murat")
```
