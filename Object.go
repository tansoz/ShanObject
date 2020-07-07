package ShanObject

import (
	"fmt"
	"strconv"
	"time"
)

type ObjectImpl struct {
	Object
	Data interface{}
}

func NewObject(v interface{})Object{
	return &ObjectImpl{
		Data: v,
	}
}
func (this *ObjectImpl)IsNil()bool{
	if this.Data == nil{
		return true
	}
	return false
}
func (this *ObjectImpl)Interface()interface{}{
	return this.Data
}

func (this *ObjectImpl)String()string{
	if this.Data == nil{
		return ""
	}
	return fmt.Sprint(this.Data)
}

func (this *ObjectImpl)Int()int{

	switch this.Data.(type) {

	case int:
		return this.Data.(int)
	case string:
		if i,err := strconv.ParseInt(this.Data.(string),10,0);err == nil{
			return int(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Int8()int8{
	switch this.Data.(type) {

	case int8:
		return this.Data.(int8)
	case string:
		if i,err := strconv.ParseInt(this.Data.(string),10,8);err == nil{
			return int8(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Int16()int16{
	switch this.Data.(type) {

	case int16:
		return this.Data.(int16)
	case string:
		if i,err := strconv.ParseInt(this.Data.(string),10,16);err == nil{
			return int16(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Int32()int32{
	switch this.Data.(type) {

	case int32:
		return this.Data.(int32)
	case string:
		if i,err := strconv.ParseInt(this.Data.(string),10,32);err == nil{
			return int32(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Int64()int64{
	switch this.Data.(type) {

	case int64:
		return this.Data.(int64)
	case string:
		if i,err := strconv.ParseInt(this.Data.(string),10,64);err == nil{
			return int64(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Uint()uint{
	switch this.Data.(type) {

	case uint:
		return this.Data.(uint)
	case string:
		if i,err := strconv.ParseUint(this.Data.(string),10,0);err == nil{
			return uint(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Uint8()uint8{
	switch this.Data.(type) {

	case uint8:
		return this.Data.(uint8)
	case string:
		if i,err := strconv.ParseUint(this.Data.(string),10,8);err == nil{
			return uint8(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Uint16()uint16{
	switch this.Data.(type) {

	case uint16:
		return this.Data.(uint16)
	case string:
		if i,err := strconv.ParseUint(this.Data.(string),10,16);err == nil{
			return uint16(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Uint32()uint32{
	switch this.Data.(type) {

	case uint32:
		return this.Data.(uint32)
	case string:
		if i,err := strconv.ParseUint(this.Data.(string),10,32);err == nil{
			return uint32(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Uint64()uint64{
	switch this.Data.(type) {

	case uint64:
		return this.Data.(uint64)
	case string:
		if i,err := strconv.ParseUint(this.Data.(string),10,64);err == nil{
			return uint64(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Float32()float32{
	switch this.Data.(type) {

	case float32:
		return this.Data.(float32)
	case string:
		if i,err := strconv.ParseFloat(this.Data.(string),32);err == nil{
			return float32(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Float64()float64{
	switch this.Data.(type) {

	case float64:
		return this.Data.(float64)
	case string:
		if i,err := strconv.ParseFloat(this.Data.(string),64);err == nil{
			return float64(i)
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Byte()byte{
	switch this.Data.(type) {

	case byte:
		return this.Data.(byte)
	case string:
		if i:=[]byte(this.Data.(string));len(i) == 1{
			return i[0]
		}else{
			return 0
		}

	}
	return 0
}
func (this *ObjectImpl)Bool()bool{
	switch this.Data.(type) {

	case bool:
		return this.Data.(bool)
	case string:
		if i,err := strconv.ParseBool(this.Data.(string));err == nil{
			return i
		}else{
			return false
		}

	}
	return false
}
func (this *ObjectImpl)Date()time.Time{

	switch this.Data.(type) {

	case time.Time:
		return this.Data.(time.Time)
	case string:

		layout := []string{
			"2006-01-02",
			"2006-01-2",
			"2006-1-2",
			"2006-1-02",
			"6-1-2",
			"6-01-2",
			"6-1-02",
			"6-01-02",
			"06-01-02",
		}

		for _,i := range layout{
			if i,err := time.Parse(i,this.Data.(string));err == nil{
				return i
			}
		}

	}
	return time.Now()
}
func (this *ObjectImpl)Datetime()time.Time{
	switch this.Data.(type) {

	case time.Time:
		return this.Data.(time.Time)
	case string:

		layout := []string{
			"2006-01-02 15:04:05",
			"2006-01-02 15:04:5",
			"2006-01-02 15:4:05",
			"2006-01-2 15:04:05",
			"2006-1-2 15:04:05",
			"2006-1-02 15:04:05",
			"6-1-2 15:4:5",
			"6-01-2 15:4:5",
			"6-1-2 15:4:05",
			"6-01-2 15:4:05",
			"6-1-2 15:04:05",
			"6-1-2 15:04:5",
			"6-1-02 15:04:05",
			"6-01-02 15:04:05",
			"06-01-02 15:04:05",
		}

		for _,i := range layout{
			if i,err := time.Parse(i,this.Data.(string));err == nil{
				return i
			}
		}

	}
	return time.Now()
}

