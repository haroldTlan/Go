package main

import (
    "os"
	//"io"
    "text/template"
	"flag"
	"fmt"
	//"reflect"
	
)

type DeviceList struct {
	TotalNum int
	List []Device
	
}

type Device struct {
    DevID string
    UserName string
	Password string
	DevName string
	Result string
}

type Record struct {
	DevID string
    ChannelNo string
    RecordLoop string
	RecordFull string
	RecordSpace string
	HashToken string
	Result string
	Time string
}


type RecordPolicy struct {
	DevID string
    ChannelNo int
	List []WeekPolicy
	HashToken string
	Result string
}

type WeekPolicy struct {
    PolicyName string
    StoreType string
	StreamType string
	Week string
	StartTime string 
	EndTime string
}

type RecordPolicy_1 struct {
	DevID string
    ChannelNo int
	List []WeekPolicy_1
	HashToken string
	Result string
}

type WeekPolicy_1 struct {
	PolicyID string
    PolicyName string
    StoreType string
	StreamType string
	Week string
	StartTime string 
	EndTime string
}

type RecordDel struct {
	DevID string
    ChannelNo int
	PolicyID string
	HashToken string
	Result string
}

type QueryFile struct {
	DevID string
    ChannelNo int
	StartTime string 
	EndTime string	
	PageNo string
	HashToken string
	Result string
	TotalRecordNum string
	List []File
}

type File struct {
	FileName string
    FileSize string
	StartTime string 
	EndTime string	
}



type RecordRes struct {
	Result string
	DevID string 
	Time string
}
type DeviceResult struct {
	Result string
}
func main() {
	var id string
	flag.StringVar(&id,"id", "list", "choose style")
	//file, err := os.Create("Struct.xml")
	//a := []string{{"b", "c",  "asd", "zz"},{"b", "c",  "asd", "zz"}}
    /*sweaters := Device{2, "b", "c",  "asd", "zz"} 
    //muban := "{{.Count}} items are made of {{.Material}}"
    //tmpl, err := template.New("test").Parse(muban)  //建立一个模板
	tmpl, err := template.ParseFiles("devicelist.xml")
    if err != nil {   
            panic(err)
    }   
	//err = tmpl.ExecuteTemplate(os.Stdout, "english", sweaters)
    err = tmpl.Execute(os.Stdout,sweaters) //将struct与模板合成，合成结果放到os.Stdout里
	//fmt.Println(p)
	
    if err != nil {
            panic(err)
	//REG()
    } */
	flag.Parse()
	if id == "setdev"{
		SetDeviceList()
	}else if id == "deldev"{
		DelDeviceList()
	}else if id == "getdev"{
		GetDeviceList()	
	}else if id == "setrec"{
		SetRecord()
	}else if id == "getrec"{
		GetRecord()	
	}else if id == "addrec"{
		AddRecord()	
	}else if id == "querec"{
		QueryRecord()	
	}else if id == "delrec"{
		DelRecord()			
	}else if id == "file"{
		FileQuery()			
	}else if id == "result"{
		result()
	}else if id == "recordres"{
		recordres()
	}	
}


func SetDeviceList() {
	v := &DeviceList{TotalNum: 2}
	for i:=0; i < v.TotalNum; i++{
		v.List = append(v.List, Device{"a", "as", "a", "a", "success"})
		}
	//v.List = append(v.List, Device{"c", "c", "c", "c"})
	//sweaters := DeviceList{Result:"failed",Obj:["a","b"]}
	tmpl, err := template.ParseFiles("tem/devset.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	
	
	return
}

func DelDeviceList() {
	v := DeviceList{TotalNum:2} 
	for i:=0; i < v.TotalNum; i++{
		v.List = append(v.List, Device{DevID:"a", Result:"success"})
		}	
	tmpl, err := template.ParseFiles("tem/devdel.xml")
    if err != nil {   
            panic(err)
    }
	fmt.Println(v)
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func GetDeviceList() {
	v := &DeviceList{TotalNum: 2}
	for i:=0; i < v.TotalNum; i++{
		v.List = append(v.List, Device{"a", "as", "a", "a", "success"})
		}
	tmpl, err := template.ParseFiles("tem/devget.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func SetRecord() {
	v := &Record{"a", "b", "c",  "asd",  "5", "6", "success", "60"} 
	tmpl, err := template.ParseFiles("tem/recordset.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func GetRecord() {
	v := &Record{"a", "b", "c",  "asd",  "5", "6", "success", "60"} 
	tmpl, err := template.ParseFiles("tem/recordget.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func AddRecord() {
	v := &RecordPolicy{DevID:"a", ChannelNo:3, Result:"success", HashToken:"hahah"} 
	v.List = append(v.List, WeekPolicy{"a", "as", "a", "a", "a","9:30"})
	tmpl, err := template.ParseFiles("tem/recordadd.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func QueryRecord() {
	v := &RecordPolicy_1{DevID:"a", ChannelNo:3, Result:"success", HashToken:"hahah"} 
	v.List = append(v.List, WeekPolicy_1{ "id", "a", "as", "a", "a", "a","9:30"})
	tmpl, err := template.ParseFiles("tem/recordquery.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func DelRecord() {
	v := &RecordDel{DevID:"a", ChannelNo:3, Result:"success", HashToken:"hahah", PolicyID:"id"} 
	tmpl, err := template.ParseFiles("tem/recorddel.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func FileQuery() {
	v := &QueryFile{DevID:"a", ChannelNo:3, Result:"success", HashToken:"hahah", TotalRecordNum:"id"}
	//v := &QueryFile{"a", 3, "success", "hahah", "id", "id", "id", "id"} 
	v.List = append(v.List, File{ "id", "a", "9:20", "18:23"})
	
	tmpl, err := template.ParseFiles("tem/queryfile.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,v)	
    if err != nil {
            panic(err)
	}
	return
}

func result() {
	sweaters := DeviceResult{"success"} 
	tmpl, err := template.ParseFiles("tem/result.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,sweaters)	
    if err != nil {
            panic(err)
	}
	return
}



func recordres() {
	sweaters := RecordRes{"a", "b", "c"} 
	tmpl, err := template.ParseFiles("tem/recordrec.xml")
    if err != nil {   
            panic(err)
    }
	
	err = tmpl.Execute(os.Stdout,sweaters)	
    if err != nil {
            panic(err)
	}
	return
}