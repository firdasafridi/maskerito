package maskerito

import "reflect"

// Marshal must input a interface{}/json, add tag json on struct fields, return a pointer interface{} of input type and it will be masked with the tag format type
//
// Example:
//
//
//	 type Sample struct {
//	 	Mobile []string `json:"mobile"`
//	 }
//
//	 func main() {
//
//	 	sample := Sample{
//	 		Mobile: []string{"081313131313"},
//	 	}
//
//	 	m, err := maskerito.New(maskerito.DefaultConfig())
//	 	fmt.Println("is err?", err)
//
//	 	c, err := m.Struct(sample)
//
//	 	output, _ := json.Marshal(c)
//	 	fmt.Println(string(output), sample, err)
//	 }
func (m *Maskerito) Marshal(s interface{}) (interface{}, error) {
	if s == nil {
		return nil, ErrInputNil
	}

	var valElement, valPtr reflect.Value

	st := reflect.TypeOf(s)

	if st.Kind() == reflect.Ptr {
		valPtr = reflect.New(st.Elem())
		valElement = reflect.ValueOf(s).Elem()
	} else {
		valPtr = reflect.New(st)
		valElement = reflect.ValueOf(s)
	}

	for i := 0; i < valElement.NumField(); i++ {
		jsonTag := valElement.Type().Field(i).Tag.Get("json")

		if jsonTag == "" {
			jsonTag = valElement.Type().Field(i).Name
		}

		mtag := dictionary[jsonTag]

		switch valElement.Field(i).Type().Kind() {
		default:
			valPtr.Elem().Field(i).Set(valElement.Field(i))
		case reflect.String:
			valPtr.Elem().Field(i).SetString(m.String(mtag, valElement.Field(i).String()))
		case reflect.Struct:
			if mtag == MStruct {
				_t, err := m.Struct(valElement.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				valPtr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())
			}
		case reflect.Ptr:
			if valElement.Field(i).IsNil() {
				continue
			}
			if mtag == MStruct {
				_t, err := m.Struct(valElement.Field(i).Interface())
				if err != nil {
					return nil, err
				}
				valPtr.Elem().Field(i).Set(reflect.ValueOf(_t))
			}
		case reflect.Slice:
			if valElement.Field(i).IsNil() {
				continue
			}
			if valElement.Field(i).Type().Elem().Kind() == reflect.String {
				orgval := valElement.Field(i).Interface().([]string)
				newval := make([]string, len(orgval))
				for i, val := range valElement.Field(i).Interface().([]string) {
					newval[i] = m.String(mtag, val)
				}
				valPtr.Elem().Field(i).Set(reflect.ValueOf(newval))
				continue
			}
			if valElement.Field(i).Type().Elem().Kind() == reflect.Struct || mtag == MStruct {
				newval := reflect.MakeSlice(valElement.Field(i).Type(), 0, valElement.Field(i).Len())
				for j, l := 0, valElement.Field(i).Len(); j < l; j++ {
					_n, err := m.Struct(valElement.Field(i).Index(j).Interface())
					if err != nil {
						return nil, err
					}
					newval = reflect.Append(newval, reflect.ValueOf(_n).Elem())
				}
				valPtr.Elem().Field(i).Set(newval)
				continue
			}
			if valElement.Field(i).Type().Elem().Kind() == reflect.Ptr || mtag == MStruct {
				newval := reflect.MakeSlice(valElement.Field(i).Type(), 0, valElement.Field(i).Len())
				for j, l := 0, valElement.Field(i).Len(); j < l; j++ {
					_n, err := m.Struct(valElement.Field(i).Index(j).Interface())
					if err != nil {
						return nil, err
					}
					newval = reflect.Append(newval, reflect.ValueOf(_n))
				}
				valPtr.Elem().Field(i).Set(newval)
				continue
			}
			if valElement.Field(i).Type().Elem().Kind() == reflect.Interface && mtag == MStruct {
				newval := reflect.MakeSlice(valElement.Field(i).Type(), 0, valElement.Field(i).Len())
				for j, l := 0, valElement.Field(i).Len(); j < l; j++ {
					_n, err := m.Struct(valElement.Field(i).Index(j).Interface())
					if err != nil {
						return nil, err
					}
					if reflect.TypeOf(valElement.Field(i).Index(j).Interface()).Kind() != reflect.Ptr {
						newval = reflect.Append(newval, reflect.ValueOf(_n).Elem())
					} else {
						newval = reflect.Append(newval, reflect.ValueOf(_n))
					}
				}
				valPtr.Elem().Field(i).Set(newval)
				continue
			}
		case reflect.Interface:
			if valElement.Field(i).IsNil() {
				continue
			}
			if mtag != MStruct {
				continue
			}
			_t, err := m.Struct(valElement.Field(i).Interface())
			if err != nil {
				return nil, err
			}
			if reflect.TypeOf(valElement.Field(i).Interface()).Kind() != reflect.Ptr {
				valPtr.Elem().Field(i).Set(reflect.ValueOf(_t).Elem())
			} else {
				valPtr.Elem().Field(i).Set(reflect.ValueOf(_t))
			}
		}
	}

	return valPtr.Interface(), nil
}
