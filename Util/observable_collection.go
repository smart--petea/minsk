package Util

import (
    "fmt"
)

type ObservableCollection struct {
    Collection []string
    Listeners []func(interface{})
}

func (o *ObservableCollection) CollectionChanged(listener func(interface{}) ) {
    o.Listeners = append(o.Listeners, listener)
}

func (o *ObservableCollection) fireCollectionChanged() {
    //log.Printf("fireCollectionChanged")
    for _, l := range o.Listeners {
        l(nil)
    }
}

func (o *ObservableCollection) Clear() {
    //log.Printf("ObservableCollection.Clear")
    o.Collection = []string{}
    o.fireCollectionChanged()
}

func (o *ObservableCollection) Insert(index int, e string) {
    //log.Printf("ObservableCollection.Insert")
    if index <= len(o.Collection) - 1 {
        collection := o.Collection
        o.Collection = append(collection[:index], e)
        o.Collection = append(o.Collection, o.Collection[:index]...)
    } else if index == len(o.Collection) {
        o.Collection = append(o.Collection, e)
    } else {
        panic(fmt.Sprintf("ObservableCollection.Insert index(%d) > len(collection)(%d)", index, len(o.Collection)))
    }

    o.fireCollectionChanged()
}

func (o *ObservableCollection) Add(e string) {
    //log.Printf("ObservableCollection.Add")
    o.Collection = append(o.Collection, e)
    o.fireCollectionChanged()
}

func (o *ObservableCollection) RemoveAt(index int) {
    if index < len(o.Collection)  - 1 {
        o.Collection = append(o.Collection[:index], o.Collection[index + 1:]...)
    } else if index == len(o.Collection)  - 1 {
        o.Collection = o.Collection[:index]
    }
}

func (o *ObservableCollection) Get(index int) string {
    return o.Collection[index]
}

func (o *ObservableCollection) Set(index int, val string) {
    //log.Printf("ObservableCollection.Set")
    o.Collection[index] = val
    o.fireCollectionChanged()
}

func (o *ObservableCollection) Count() int {
    return len(o.Collection)
}

func NewObservableCollection(collection... string) *ObservableCollection {
    return &ObservableCollection{
        Collection: collection,
    }
}

