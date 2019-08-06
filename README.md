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
 
p.Create(obj InterStorage, key, data string)
p.Update(obj InterStorage, key, data string)
p.Read(obj InterStorage, key string) (string, error)
p.Delete(obj InterStorage, key string)
```
For CURD your property.
