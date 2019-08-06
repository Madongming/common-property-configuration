# common-property-configuration

Requirement berkeley db for execute the 'go test' command.

Implement the interface
```
type InterStorage interface {
        Create(key string, data []byte) error
        Update(key string, data []byte) error
        Read(key string) ([]byte, error)
        Delete(key string) ([]byte, error)
}
```
The object can:
```
import (
    p "Madongming/common-property-configuration"
)
 
err := p.Create(obj InterStorage, key, data string)
err = p.Update(obj InterStorage, key, data string)
value, err := p.Read(obj InterStorage, key string)
value, err = p.Delete(obj InterStorage, key string)
```
For CURD your property.
