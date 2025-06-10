package configura

import (
	"errors"
	"maps"
	"sync"
)

var ErrMissingVariable = errors.New("missing configuration variables")

type constraint interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | []byte | []rune | float32 | float64 | bool
}

type Variable[T constraint] string

// Config is an interface that defines methods for accessing configuration variables of various types.
type Config interface {
	String(key Variable[string]) string
	Int(key Variable[int]) int
	Int8(key Variable[int8]) int8
	Int16(key Variable[int16]) int16
	Int32(key Variable[int32]) int32
	Int64(key Variable[int64]) int64
	Uint(key Variable[uint]) uint
	Uint8(key Variable[uint8]) uint8
	Uint16(key Variable[uint16]) uint16
	Uint32(key Variable[uint32]) uint32
	Uint64(key Variable[uint64]) uint64
	Uintptr(key Variable[uintptr]) uintptr
	Bytes(key Variable[[]byte]) []byte
	Runes(key Variable[[]rune]) []rune
	Float32(key Variable[float32]) float32
	Float64(key Variable[float64]) float64
	Bool(key Variable[bool]) bool
	ConfigurationKeysRegistered(keys ...any) error
}

// Locks for each type of configuration variable to ensure thread-safe access.
var (
	stringLock  = sync.RWMutex{}
	intLock     = sync.RWMutex{}
	int8Lock    = sync.RWMutex{}
	int16Lock   = sync.RWMutex{}
	int32Lock   = sync.RWMutex{}
	int64Lock   = sync.RWMutex{}
	uintLock    = sync.RWMutex{}
	uint8Lock   = sync.RWMutex{}
	uint16Lock  = sync.RWMutex{}
	uint32Lock  = sync.RWMutex{}
	uint64Lock  = sync.RWMutex{}
	uintptrLock = sync.RWMutex{}
	bytesLock   = sync.RWMutex{}
	runesLock   = sync.RWMutex{}
	float32Lock = sync.RWMutex{}
	float64Lock = sync.RWMutex{}
	boolLock    = sync.RWMutex{}
)

// WriteConfiguration is a generic function that writes configuration values to the provided configuration struct.
// It uses type assertions to determine the type of the values and writes them to the appropriate map in the
// configuration struct. This function is designed to be used to Mock the configuration in tests or to set
// default values in the configuration struct. It will overwrite any existing values for the keys provided and is
// not meant to be used for runtime configuration changes.
func WriteConfiguration[T constraint](cfg Config, values map[Variable[T]]T) error {
	if cfg == nil {
		return errors.New("Config cannot be nil")
	}

	typecastCfg, ok := cfg.(*ConfigImpl) // Type assertion to *ConfigImpl
	if !ok {
		return errors.New("invalid configuration type, expected *ConfigImpl")
	}

	switch v := any(values).(type) {
	case map[Variable[string]]string:
		stringLock.Lock()
		defer stringLock.Unlock()
		typecastCfg.regString = v
	case map[Variable[int]]int:
		intLock.Lock()
		defer intLock.Unlock()
		typecastCfg.regInt = v
	case map[Variable[int8]]int8:
		int8Lock.Lock()
		defer int8Lock.Unlock()
		typecastCfg.regInt8 = v
	case map[Variable[int16]]int16:
		int16Lock.Lock()
		defer int16Lock.Unlock()
		typecastCfg.regInt16 = v
	case map[Variable[int32]]int32:
		int32Lock.Lock()
		defer int32Lock.Unlock()
		typecastCfg.regInt32 = v
	case map[Variable[int64]]int64:
		int64Lock.Lock()
		defer int64Lock.Unlock()
		typecastCfg.regInt64 = v
	case map[Variable[uint]]uint:
		uintLock.Lock()
		defer uintLock.Unlock()
		typecastCfg.regUint = v
	case map[Variable[uint8]]uint8:
		uint8Lock.Lock()
		defer uint8Lock.Unlock()
		typecastCfg.regUint8 = v
	case map[Variable[uint16]]uint16:
		uint16Lock.Lock()
		defer uint16Lock.Unlock()
		typecastCfg.regUint16 = v
	case map[Variable[uint32]]uint32:
		uint32Lock.Lock()
		defer uint32Lock.Unlock()
		typecastCfg.regUint32 = v
	case map[Variable[uint64]]uint64:
		uint64Lock.Lock()
		defer uint64Lock.Unlock()
		typecastCfg.regUint64 = v
	case map[Variable[uintptr]]uintptr:
		uintptrLock.Lock()
		defer uintptrLock.Unlock()
		typecastCfg.regUintptr = v
	case map[Variable[[]byte]][]byte:
		bytesLock.Lock()
		defer bytesLock.Unlock()
		typecastCfg.regBytes = v
	case map[Variable[[]rune]][]rune:
		runesLock.Lock()
		defer runesLock.Unlock()
		typecastCfg.regRunes = v
	case map[Variable[float32]]float32:
		float32Lock.Lock()
		defer float32Lock.Unlock()
		typecastCfg.regFloat32 = v
	case map[Variable[float64]]float64:
		float64Lock.Lock()
		defer float64Lock.Unlock()
		typecastCfg.regFloat64 = v
	case map[Variable[bool]]bool:
		boolLock.Lock()
		defer boolLock.Unlock()
		typecastCfg.regBool = v
	default:
		return errors.New("unsupported values type for WriteConfiguration")
	}

	return nil
}

// LoadEnvironment is a generic function that loads an environment variable into the provided configuration,
// using the specified key and fallback value. It uses type assertions to determine the type of the key
// and fallback value, and registers the variable in the appropriate map of the configuration struct.
func LoadEnvironment[T constraint](config *ConfigImpl, key Variable[T], fallback T) {
	switch any(key).(type) {
	case Variable[string]:
		stringLock.Lock()
		defer stringLock.Unlock()
		config.regString[any(key).(Variable[string])] = String(any(key).(Variable[string]), any(fallback).(string))
	case Variable[int]:
		intLock.Lock()
		defer intLock.Unlock()
		config.regInt[any(key).(Variable[int])] = Int(any(key).(Variable[int]), any(fallback).(int))
	case Variable[int8]:
		int8Lock.Lock()
		defer int8Lock.Unlock()
		config.regInt8[any(key).(Variable[int8])] = Int8(any(key).(Variable[int8]), any(fallback).(int8))
	case Variable[int16]:
		int16Lock.Lock()
		defer int16Lock.Unlock()
		config.regInt16[any(key).(Variable[int16])] = Int16(any(key).(Variable[int16]), any(fallback).(int16))
	case Variable[int32]:
		int32Lock.Lock()
		defer int32Lock.Unlock()
		config.regInt32[any(key).(Variable[int32])] = Int32(any(key).(Variable[int32]), any(fallback).(int32))
	case Variable[int64]:
		int64Lock.Lock()
		defer int64Lock.Unlock()
		config.regInt64[any(key).(Variable[int64])] = Int64(any(key).(Variable[int64]), any(fallback).(int64))
	case Variable[uint]:
		uintLock.Lock()
		defer uintLock.Unlock()
		config.regUint[any(key).(Variable[uint])] = Uint(any(key).(Variable[uint]), any(fallback).(uint))
	case Variable[uint8]:
		uint8Lock.Lock()
		defer uint8Lock.Unlock()
		config.regUint8[any(key).(Variable[uint8])] = Uint8(any(key).(Variable[uint8]), any(fallback).(uint8))
	case Variable[uint16]:
		uint16Lock.Lock()
		defer uint16Lock.Unlock()
		config.regUint16[any(key).(Variable[uint16])] = Uint16(any(key).(Variable[uint16]), any(fallback).(uint16))
	case Variable[uint32]:
		uint32Lock.Lock()
		defer uint32Lock.Unlock()
		config.regUint32[any(key).(Variable[uint32])] = Uint32(any(key).(Variable[uint32]), any(fallback).(uint32))
	case Variable[uint64]:
		uint64Lock.Lock()
		defer uint64Lock.Unlock()
		config.regUint64[any(key).(Variable[uint64])] = Uint64(any(key).(Variable[uint64]), any(fallback).(uint64))
	case Variable[uintptr]:
		uintptrLock.Lock()
		defer uintptrLock.Unlock()
		config.regUintptr[any(key).(Variable[uintptr])] = Uintptr(any(key).(Variable[uintptr]), any(fallback).(uintptr))
	case Variable[[]byte]:
		bytesLock.Lock()
		defer bytesLock.Unlock()
		config.regBytes[any(key).(Variable[[]byte])] = Bytes(any(key).(Variable[[]byte]), any(fallback).([]byte))
	case Variable[[]rune]:
		runesLock.Lock()
		defer runesLock.Unlock()
		config.regRunes[any(key).(Variable[[]rune])] = Runes(any(key).(Variable[[]rune]), any(fallback).([]rune))
	case Variable[float32]:
		float32Lock.Lock()
		defer float32Lock.Unlock()
		config.regFloat32[any(key).(Variable[float32])] = Float32(any(key).(Variable[float32]), any(fallback).(float32))
	case Variable[float64]:
		float64Lock.Lock()
		defer float64Lock.Unlock()
		config.regFloat64[any(key).(Variable[float64])] = Float64(any(key).(Variable[float64]), any(fallback).(float64))
	case Variable[bool]:
		boolLock.Lock()
		defer boolLock.Unlock()
		config.regBool[any(key).(Variable[bool])] = Bool(any(key).(Variable[bool]), any(fallback).(bool))
	}
}

// ConfigImpl is a concrete implementation of the Config interface, holding maps for each type of configuration
// variable. It provides methods to retrieve values for each type and checks if all required keys are registered.
type ConfigImpl struct {
	regString  map[Variable[string]]string
	regInt     map[Variable[int]]int
	regInt8    map[Variable[int8]]int8
	regInt16   map[Variable[int16]]int16
	regInt32   map[Variable[int32]]int32
	regInt64   map[Variable[int64]]int64
	regUint    map[Variable[uint]]uint
	regUint8   map[Variable[uint8]]uint8
	regUint16  map[Variable[uint16]]uint16
	regUint32  map[Variable[uint32]]uint32
	regUint64  map[Variable[uint64]]uint64
	regUintptr map[Variable[uintptr]]uintptr
	regBytes   map[Variable[[]byte]][]byte
	regRunes   map[Variable[[]rune]][]rune
	regFloat32 map[Variable[float32]]float32
	regFloat64 map[Variable[float64]]float64
	regBool    map[Variable[bool]]bool
}

func NewConfigImpl() *ConfigImpl {
	return &ConfigImpl{
		regString:  make(map[Variable[string]]string),
		regInt:     make(map[Variable[int]]int),
		regInt8:    make(map[Variable[int8]]int8),
		regInt16:   make(map[Variable[int16]]int16),
		regInt32:   make(map[Variable[int32]]int32),
		regInt64:   make(map[Variable[int64]]int64),
		regUint:    make(map[Variable[uint]]uint),
		regUint8:   make(map[Variable[uint8]]uint8),
		regUint16:  make(map[Variable[uint16]]uint16),
		regUint32:  make(map[Variable[uint32]]uint32),
		regUint64:  make(map[Variable[uint64]]uint64),
		regUintptr: make(map[Variable[uintptr]]uintptr),
		regBytes:   make(map[Variable[[]byte]][]byte),
		regRunes:   make(map[Variable[[]rune]][]rune),
		regFloat32: make(map[Variable[float32]]float32),
		regFloat64: make(map[Variable[float64]]float64),
		regBool:    make(map[Variable[bool]]bool),
	}
}

var _ Config = (*ConfigImpl)(nil)

func (c *ConfigImpl) String(key Variable[string]) string {
	stringLock.RLock()
	defer stringLock.RUnlock()
	if value, exists := c.regString[key]; exists {
		return value
	}
	return ""
}

func (c *ConfigImpl) Int(key Variable[int]) int {
	intLock.RLock()
	defer intLock.RUnlock()
	if value, exists := c.regInt[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Int8(key Variable[int8]) int8 {
	int8Lock.RLock()
	defer int8Lock.RUnlock()
	if value, exists := c.regInt8[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Int16(key Variable[int16]) int16 {
	int16Lock.RLock()
	defer int16Lock.RUnlock()
	if value, exists := c.regInt16[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Int32(key Variable[int32]) int32 {
	int32Lock.RLock()
	defer int32Lock.RUnlock()
	if value, exists := c.regInt32[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Int64(key Variable[int64]) int64 {
	int64Lock.RLock()
	defer int64Lock.RUnlock()
	if value, exists := c.regInt64[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uint(key Variable[uint]) uint {
	uintLock.RLock()
	defer uintLock.RUnlock()
	if value, exists := c.regUint[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uint8(key Variable[uint8]) uint8 {
	uint8Lock.RLock()
	defer uint8Lock.RUnlock()
	if value, exists := c.regUint8[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uint16(key Variable[uint16]) uint16 {
	uint16Lock.RLock()
	defer uint16Lock.RUnlock()
	if value, exists := c.regUint16[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uint32(key Variable[uint32]) uint32 {
	uint32Lock.RLock()
	defer uint32Lock.RUnlock()
	if value, exists := c.regUint32[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uint64(key Variable[uint64]) uint64 {
	uint64Lock.RLock()
	defer uint64Lock.RUnlock()
	if value, exists := c.regUint64[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Uintptr(key Variable[uintptr]) uintptr {
	uintptrLock.RLock()
	defer uintptrLock.RUnlock()
	if value, exists := c.regUintptr[key]; exists {
		return value
	}
	return 0
}

func (c *ConfigImpl) Bytes(key Variable[[]byte]) []byte {
	bytesLock.RLock()
	defer bytesLock.RUnlock()
	if value, exists := c.regBytes[key]; exists {
		return value
	}
	return nil
}

func (c *ConfigImpl) Runes(key Variable[[]rune]) []rune {
	runesLock.RLock()
	defer runesLock.RUnlock()
	if value, exists := c.regRunes[key]; exists {
		return value
	}
	return nil
}

func (c *ConfigImpl) Float32(key Variable[float32]) float32 {
	float32Lock.RLock()
	defer float32Lock.RUnlock()
	if value, exists := c.regFloat32[key]; exists {
		return value
	}
	return 0.0
}

func (c *ConfigImpl) Float64(key Variable[float64]) float64 {
	float64Lock.RLock()
	defer float64Lock.RUnlock()
	if value, exists := c.regFloat64[key]; exists {
		return value
	}
	return 0.0
}

func (c *ConfigImpl) Bool(key Variable[bool]) bool {
	boolLock.RLock()
	defer boolLock.RUnlock()
	if value, exists := c.regBool[key]; exists {
		return value
	}
	return false
}

// missingVariableError is an error type that holds a list of missing configuration variable keys.
type missingVariableError struct {
	Keys []string
}

// Error implements the error interface for missingVariableError.
func (e missingVariableError) Error() string {
	return "missing configuration variables: " + formatKeys(e.Keys)
}

// Unwrap implements the Unwrap method for the error interface, allowing the error to be unwrapped to ErrMissingVariable.
func (e missingVariableError) Unwrap() error {
	return ErrMissingVariable
}

// formatKeys formats the keys into a string for error messages. If no keys are provided, it returns "none".
func formatKeys(keys []string) string {
	if len(keys) == 0 {
		return "none"
	}
	result := ""
	for i, key := range keys {
		if i > 0 {
			result += ", "
		}
		result += string(key)
	}
	return result
}

var _ error = (*missingVariableError)(nil)

// checkKey checks if the provided key exists in the configuration. It uses type assertion to determine the type of the
// key and checks the corresponding map in the configuration struct.
func (c *ConfigImpl) checkKey(key any) (string, bool) {
	var exists bool
	var keyName string
	switch k := key.(type) {
	case Variable[string]:
		stringLock.RLock()
		defer stringLock.RUnlock()
		_, exists = c.regString[k]
		keyName = string(k)
	case Variable[int]:
		intLock.RLock()
		defer intLock.RUnlock()
		_, exists = c.regInt[k]
		keyName = string(k)
	case Variable[int8]:
		int8Lock.RLock()
		defer int8Lock.RUnlock()
		_, exists = c.regInt8[k]
		keyName = string(k)
	case Variable[int16]:
		int16Lock.RLock()
		defer int16Lock.RUnlock()
		_, exists = c.regInt16[k]
		keyName = string(k)
	case Variable[int32]:
		int32Lock.RLock()
		defer int32Lock.RUnlock()
		_, exists = c.regInt32[k]
		keyName = string(k)
	case Variable[int64]:
		int64Lock.RLock()
		defer int64Lock.RUnlock()
		_, exists = c.regInt64[k]
		keyName = string(k)
	case Variable[uint]:
		uintLock.RLock()
		defer uintLock.RUnlock()
		_, exists = c.regUint[k]
		keyName = string(k)
	case Variable[uint8]:
		uint8Lock.RLock()
		defer uint8Lock.RUnlock()
		_, exists = c.regUint8[k]
		keyName = string(k)
	case Variable[uint16]:
		uint16Lock.RLock()
		defer uint16Lock.RUnlock()
		_, exists = c.regUint16[k]
		keyName = string(k)
	case Variable[uint32]:
		uint32Lock.RLock()
		defer uint32Lock.RUnlock()
		_, exists = c.regUint32[k]
		keyName = string(k)
	case Variable[uint64]:
		uint64Lock.RLock()
		defer uint64Lock.RUnlock()
		_, exists = c.regUint64[k]
		keyName = string(k)
	case Variable[uintptr]:
		uintptrLock.RLock()
		defer uintptrLock.RUnlock()
		_, exists = c.regUintptr[k]
		keyName = string(k)
	case Variable[[]byte]:
		bytesLock.RLock()
		defer bytesLock.RUnlock()
		_, exists = c.regBytes[k]
		keyName = string(k)
	case Variable[[]rune]:
		runesLock.RLock()
		defer runesLock.RUnlock()
		_, exists = c.regRunes[k]
		keyName = string(k)
	case Variable[float32]:
		float32Lock.RLock()
		defer float32Lock.RUnlock()
		_, exists = c.regFloat32[k]
		keyName = string(k)
	case Variable[float64]:
		float64Lock.RLock()
		defer float64Lock.RUnlock()
		_, exists = c.regFloat64[k]
		keyName = string(k)
	case Variable[bool]:
		boolLock.RLock()
		defer boolLock.RUnlock()
		_, exists = c.regBool[k]
		keyName = string(k)
	}

	return keyName, exists
}

// ConfigurationKeysRegistered checks if all provided keys are registered in the configuration. To ensure that the
// client of the package have taken all required keys into consideration when building the configuration object.
func (c *ConfigImpl) ConfigurationKeysRegistered(keys ...any) error {
	var missingKeys []string
	for _, key := range keys {
		if keyName, ok := c.checkKey(key); !ok {
			missingKeys = append(missingKeys, keyName)
		}
	}

	if len(missingKeys) > 0 {
		return missingVariableError{Keys: missingKeys}
	}

	return nil
}

// Fallback is a helper function that returns the fallback value if the provided value is empty.
// Only works on comparable types, which includes basic types like int, string, bool, etc.
func Fallback[T comparable](value T, fallback T) T {
	var emptyValue T
	if value == emptyValue {
		return fallback
	}
	return value
}

// Merge combines multiple Config instances into a single Config instance.
// To ensure a consistent view of the source configurations, it locks all
// configuration types for reading during the merge operation.
func Merge(cfgs ...Config) Config {
	stringLock.RLock()
	defer stringLock.RUnlock()
	intLock.RLock()
	defer intLock.RUnlock()
	int8Lock.RLock()
	defer int8Lock.RUnlock()
	int16Lock.RLock()
	defer int16Lock.RUnlock()
	int32Lock.RLock()
	defer int32Lock.RUnlock()
	int64Lock.RLock()
	defer int64Lock.RUnlock()
	uintLock.RLock()
	defer uintLock.RUnlock()
	uint8Lock.RLock()
	defer uint8Lock.RUnlock()
	uint16Lock.RLock()
	defer uint16Lock.RUnlock()
	uint32Lock.RLock()
	defer uint32Lock.RUnlock()
	uint64Lock.RLock()
	defer uint64Lock.RUnlock()
	uintptrLock.RLock()
	defer uintptrLock.RUnlock()
	bytesLock.RLock()
	defer bytesLock.RUnlock()
	runesLock.RLock()
	defer runesLock.RUnlock()
	float32Lock.RLock()
	defer float32Lock.RUnlock()
	float64Lock.RLock()
	defer float64Lock.RUnlock()
	boolLock.RLock()
	defer boolLock.RUnlock()

	merged := NewConfigImpl()

	for _, cfg := range cfgs {
		if c, ok := cfg.(*ConfigImpl); ok {
			maps.Copy(merged.regString, c.regString)
			maps.Copy(merged.regInt, c.regInt)
			maps.Copy(merged.regInt8, c.regInt8)
			maps.Copy(merged.regInt16, c.regInt16)
			maps.Copy(merged.regInt32, c.regInt32)
			maps.Copy(merged.regInt64, c.regInt64)
			maps.Copy(merged.regUint, c.regUint)
			maps.Copy(merged.regUint8, c.regUint8)
			maps.Copy(merged.regUint16, c.regUint16)
			maps.Copy(merged.regUint32, c.regUint32)
			maps.Copy(merged.regUint64, c.regUint64)
			maps.Copy(merged.regUintptr, c.regUintptr)
			maps.Copy(merged.regBytes, c.regBytes)
			maps.Copy(merged.regRunes, c.regRunes)
			maps.Copy(merged.regFloat32, c.regFloat32)
			maps.Copy(merged.regFloat64, c.regFloat64)
			maps.Copy(merged.regBool, c.regBool)
		} else {
			panic("unsupported config type")
		}
	}
	return merged
}
