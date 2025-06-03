package configura

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// --- Test Suite Definitions ---

// ConfigSuite tests the Config interface and ConfigImpl methods
type ConfigSuite struct {
	suite.Suite
	config *ConfigImpl
}

// LoadEnvironmentSuite tests the LoadEnvironment function
type LoadEnvironmentSuite struct {
	suite.Suite
}

// FormatKeysSuite tests the formatKeys function
type FormatKeysSuite struct {
	suite.Suite
}

// CheckKeySuite tests the checkKey method of ConfigImpl
type CheckKeySuite struct {
	suite.Suite
}

// ConfigurationKeysRegisteredSuite tests the ConfigurationKeysRegistered method
type ConfigurationKeysRegisteredSuite struct {
	suite.Suite
}

// --- Setup Methods ---

func (s *ConfigSuite) SetupTest() {
	s.config = NewConfigImpl()
}

// --- Test Methods for ConfigSuite ---

func (s *ConfigSuite) TestString() {
	key := Variable[string]("TEST_STRING")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), "", s.config.String(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegString[key] = "hello"
		assert.Equal(s.T(), "hello", s.config.String(key))
	})
}

func (s *ConfigSuite) TestInt() {
	key := Variable[int]("TEST_INT")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), 0, s.config.Int(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegInt[key] = 123
		assert.Equal(s.T(), 123, s.config.Int(key))
	})
}

func (s *ConfigSuite) TestInt8() {
	key := Variable[int8]("TEST_INT8")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int8(0), s.config.Int8(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegInt8[key] = int8(12)
		assert.Equal(s.T(), int8(12), s.config.Int8(key))
	})
}

func (s *ConfigSuite) TestInt16() {
	key := Variable[int16]("TEST_INT16")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int16(0), s.config.Int16(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegInt16[key] = int16(1234)
		assert.Equal(s.T(), int16(1234), s.config.Int16(key))
	})
}

func (s *ConfigSuite) TestInt32() {
	key := Variable[int32]("TEST_INT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int32(0), s.config.Int32(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegInt32[key] = int32(123456)
		assert.Equal(s.T(), int32(123456), s.config.Int32(key))
	})
}

func (s *ConfigSuite) TestInt64() {
	key := Variable[int64]("TEST_INT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), int64(0), s.config.Int64(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegInt64[key] = int64(1234567890)
		assert.Equal(s.T(), int64(1234567890), s.config.Int64(key))
	})
}

func (s *ConfigSuite) TestUint() {
	key := Variable[uint]("TEST_UINT")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint(0), s.config.Uint(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUint[key] = uint(123)
		assert.Equal(s.T(), uint(123), s.config.Uint(key))
	})
}

func (s *ConfigSuite) TestUint8() {
	key := Variable[uint8]("TEST_UINT8")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint8(0), s.config.Uint8(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUint8[key] = uint8(12)
		assert.Equal(s.T(), uint8(12), s.config.Uint8(key))
	})
}

func (s *ConfigSuite) TestUint16() {
	key := Variable[uint16]("TEST_UINT16")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint16(0), s.config.Uint16(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUint16[key] = uint16(1234)
		assert.Equal(s.T(), uint16(1234), s.config.Uint16(key))
	})
}

func (s *ConfigSuite) TestUint32() {
	key := Variable[uint32]("TEST_UINT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint32(0), s.config.Uint32(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUint32[key] = uint32(123456)
		assert.Equal(s.T(), uint32(123456), s.config.Uint32(key))
	})
}

func (s *ConfigSuite) TestUint64() {
	key := Variable[uint64]("TEST_UINT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uint64(0), s.config.Uint64(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUint64[key] = uint64(1234567890)
		assert.Equal(s.T(), uint64(1234567890), s.config.Uint64(key))
	})
}

func (s *ConfigSuite) TestUintptr() {
	key := Variable[uintptr]("TEST_UINTPTR")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), uintptr(0), s.config.Uintptr(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegUintptr[key] = uintptr(0xdeadbeef)
		assert.Equal(s.T(), uintptr(0xdeadbeef), s.config.Uintptr(key))
	})
}

func (s *ConfigSuite) TestBytes() {
	key := Variable[[]byte]("TEST_BYTES")
	s.Run("KeyNotExists", func() {
		assert.Nil(s.T(), s.config.Bytes(key))
	})
	s.Run("KeyExists", func() {
		val := []byte("hello")
		s.config.RegBytes[key] = val
		assert.Equal(s.T(), val, s.config.Bytes(key))
	})
}

func (s *ConfigSuite) TestRunes() {
	key := Variable[[]rune]("TEST_RUNES")
	s.Run("KeyNotExists", func() {
		assert.Nil(s.T(), s.config.Runes(key))
	})
	s.Run("KeyExists", func() {
		val := []rune("hello😊")
		s.config.RegRunes[key] = val
		assert.Equal(s.T(), val, s.config.Runes(key))
	})
}

func (s *ConfigSuite) TestFloat32() {
	key := Variable[float32]("TEST_FLOAT32")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), float32(0.0), s.config.Float32(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegFloat32[key] = float32(3.14)
		assert.Equal(s.T(), float32(3.14), s.config.Float32(key))
	})
}

func (s *ConfigSuite) TestFloat64() {
	key := Variable[float64]("TEST_FLOAT64")
	s.Run("KeyNotExists", func() {
		assert.Equal(s.T(), float64(0.0), s.config.Float64(key))
	})
	s.Run("KeyExists", func() {
		s.config.RegFloat64[key] = float64(3.14159)
		assert.Equal(s.T(), float64(3.14159), s.config.Float64(key))
	})
}

func (s *ConfigSuite) TestBool() {
	key := Variable[bool]("TEST_BOOL")
	s.Run("KeyNotExists", func() {
		assert.False(s.T(), s.config.Bool(key))
	})
	s.Run("KeyExistsTrue", func() {
		s.config.RegBool[key] = true
		assert.True(s.T(), s.config.Bool(key))
	})
	s.Run("KeyExistsFalse", func() {
		falseKey := Variable[bool]("TEST_BOOL_FALSE_VAL")
		s.config.RegBool[falseKey] = false
		assert.False(s.T(), s.config.Bool(falseKey))
	})
}

// --- Test Methods for LoadEnvironmentSuite ---

func (s *LoadEnvironmentSuite) setEnvVar(key string, value string) {
	err := os.Setenv(key, value)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		os.Unsetenv(key)
	})
}

func (s *LoadEnvironmentSuite) unsetEnvVar(key string) {
	err := os.Unsetenv(key)
	// Allow unset to fail if var doesn't exist, as that's fine for test setup.
	if err != nil && !errors.Is(err, os.ErrNotExist) && !strings.Contains(err.Error(), "unsetenv: EINVAL: Invalid argument") { // macOS specific error for empty key
		s.Require().NoError(err) // Fail for other errors
	}
}

func (s *LoadEnvironmentSuite) TestLoadString() {
	cfg := NewConfigImpl()
	key := Variable[string]("ENV_STR")
	fallback := "fb_str"
	envVal := "env_str_val"

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), envVal)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegString[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))                        // Ensure it's unset for this specific sub-test
		LoadEnvironment(cfg, key, fallback)               // Use fresh config or reset
		assert.Equal(s.T(), fallback, cfg.RegString[key]) // This line would fail if cfg is not reset or key re-added

		// Corrected approach for EnvVarNotSet
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegString[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt() {
	cfg := NewConfigImpl()
	key := Variable[int]("ENV_INT")
	fallback := 100
	envValStr := "200"
	envValInt := 200

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValInt, cfg.RegInt[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt8() {
	cfg := NewConfigImpl()
	key := Variable[int8]("ENV_INT8")
	fallback := int8(10)
	envValStr := "20"
	envValInt8 := int8(20)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValInt8, cfg.RegInt8[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt8[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int8")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt8[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "129") // Out of bounds for int8
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt8[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt16() {
	cfg := NewConfigImpl()
	key := Variable[int16]("ENV_INT16")
	fallback := int16(1000)
	envValStr := "2000"
	envValInt16 := int16(2000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValInt16, cfg.RegInt16[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt16[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int16")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt16[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "32768") // Out of bounds for int16
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt16[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt32() {
	cfg := NewConfigImpl()
	key := Variable[int32]("ENV_INT32")
	fallback := int32(100000)
	envValStr := "200000"
	envValInt32 := int32(200000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValInt32, cfg.RegInt32[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt32[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int32")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt32[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "2147483648") // Out of bounds for int32
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt32[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt64() {
	cfg := NewConfigImpl()
	key := Variable[int64]("ENV_INT64")
	fallback := int64(1000000000)
	envValStr := "2000000000"
	envValInt64 := int64(2000000000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValInt64, cfg.RegInt64[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt64[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int64")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt64[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "9223372036854775808") // Out of bounds for int64
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegInt64[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUint() {
	cfg := NewConfigImpl()
	key := Variable[uint]("ENV_UINT")
	fallback := uint(100)
	envValStr := "200"
	envValUint := uint(200)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUint, cfg.RegUint[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint[key])
	})
	s.Run("EnvVarSetNegative", func() {
		s.setEnvVar(string(key), "-1") // Negative, invalid for uint
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUint8() {
	cfg := NewConfigImpl()
	key := Variable[uint8]("ENV_UINT8")
	fallback := uint8(10)
	envValStr := "20"
	envValUint8 := uint8(20)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUint8, cfg.RegUint8[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint8[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint8")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint8[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "256") // Out of bounds for uint8
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint8[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUint16() {
	cfg := NewConfigImpl()
	key := Variable[uint16]("ENV_UINT16")
	fallback := uint16(1000)
	envValStr := "2000"
	envValUint16 := uint16(2000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUint16, cfg.RegUint16[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint16[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint16")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint16[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "65536") // Out of bounds for uint16
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint16[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUint32() {
	cfg := NewConfigImpl()
	key := Variable[uint32]("ENV_UINT32")
	fallback := uint32(100000)
	envValStr := "200000"
	envValUint32 := uint32(200000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUint32, cfg.RegUint32[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint32[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint32")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint32[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "4294967296") // Out of bounds for uint32
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint32[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUint64() {
	cfg := NewConfigImpl()
	key := Variable[uint64]("ENV_UINT64")
	fallback := uint64(1000000000)
	envValStr := "2000000000"
	envValUint64 := uint64(2000000000)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUint64, cfg.RegUint64[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint64[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uint64")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint64[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "18446744073709551616") // Out of bounds for uint64
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUint64[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadUintptr() {
	cfg := NewConfigImpl()
	key := Variable[uintptr]("ENV_UINTPTR")
	fallback := uintptr(0x1000)
	envValStr := "0x2000" // Using hex for variety
	var envValUintptr uintptr
	_, err := fmt.Sscan(envValStr, &envValUintptr) // Parse hex string to uintptr
	s.Require().NoError(err)

	s.Run("EnvVarSetValid", func() {
		// strconv.ParseUint expects decimal unless base is specified (e.g. 0 for auto-detect prefix, 16 for hex)
		// For simplicity, let's use decimal strings for env vars if uintptr helper uses ParseUint(str, 10, ...)
		// Or adjust the helper stub to handle "0x"
		decimalEnvValStr := strconv.FormatUint(uint64(envValUintptr), 10)
		s.setEnvVar(string(key), decimalEnvValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envValUintptr, cfg.RegUintptr[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUintptr[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-uintptr")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegUintptr[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadFloat32() {
	cfg := NewConfigImpl()
	key := Variable[float32]("ENV_FLOAT32")
	fallback := float32(1.23)
	envValStr := "4.56"
	envValFloat32 := float32(4.56)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.InDelta(s.T(), envValFloat32, cfg.RegFloat32[key], 0.0001)
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.RegFloat32[key], 0.0001)
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-float")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.RegFloat32[key], 0.0001)
	})
}

func (s *LoadEnvironmentSuite) TestLoadFloat64() {
	cfg := NewConfigImpl()
	key := Variable[float64]("ENV_FLOAT64")
	fallback := float64(1.23456)
	envValStr := "7.89012"
	envValFloat64 := float64(7.89012)

	s.Run("EnvVarSetValid", func() {
		s.setEnvVar(string(key), envValStr)
		LoadEnvironment(cfg, key, fallback)
		assert.InDelta(s.T(), envValFloat64, cfg.RegFloat64[key], 0.0000001)
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.RegFloat64[key], 0.0000001)
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-float64")
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.InDelta(s.T(), fallback, freshCfg.RegFloat64[key], 0.0000001)
	})
}

func (s *LoadEnvironmentSuite) TestLoadBool() {
	key := Variable[bool]("ENV_BOOL")

	testCases := []struct {
		name        string
		envValue    *string
		fallback    bool
		expectedReg bool
	}{
		{"EnvTrueFallbackF", func(s string) *string { return &s }("true"), false, true},
		{"EnvFalseFallbackT", func(s string) *string { return &s }("false"), true, false},
		{"Env1FallbackF", func(s string) *string { return &s }("1"), false, true},
		{"Env0FallbackT", func(s string) *string { return &s }("0"), true, false},
		{"EnvTFallbackF", func(s string) *string { return &s }("t"), false, true},
		{"EnvFFallbackT", func(s string) *string { return &s }("f"), true, false},
		{"EnvInvalidFallbackF", func(s string) *string { return &s }("invalid"), false, false},
		{"EnvInvalidFallbackT", func(s string) *string { return &s }("invalid"), true, true},
		{"EnvNotSetFallbackF", nil, false, false},
		{"EnvNotSetFallbackT", nil, true, true},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			currentCfg := NewConfigImpl()
			if tc.envValue != nil {
				s.setEnvVar(string(key), *tc.envValue)
			} else {
				s.unsetEnvVar(string(key))
			}
			LoadEnvironment(currentCfg, key, tc.fallback)
			assert.Equal(s.T(), tc.expectedReg, currentCfg.RegBool[key], "Mismatch in registered bool value")
		})
	}
}

func (s *LoadEnvironmentSuite) TestLoadBytes() {
	cfg := NewConfigImpl()
	key := Variable[[]byte]("ENV_BYTES")
	fallback := []byte("fb_bytes")
	envVal := []byte("env_bytes_val")

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envVal))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegBytes[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegBytes[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadRunes() {
	cfg := NewConfigImpl()
	key := Variable[[]rune]("ENV_RUNES")
	fallback := []rune("fb_runes")
	envValStr := "env_runes_😊"
	envVal := []rune(envValStr)

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envValStr)) // Store string in env
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegRunes[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		freshCfg := NewConfigImpl()
		LoadEnvironment(freshCfg, key, fallback)
		assert.Equal(s.T(), fallback, freshCfg.RegRunes[key])
	})
}

// --- Test Methods for FormatKeysSuite ---

func (s *FormatKeysSuite) TestFormatKeys() {
	testCases := []struct {
		name     string
		keys     []string
		expected string
	}{
		{"Empty", []string{}, "none"},
		{"Single", []string{"KEY1"}, "KEY1"},
		{"Multiple", []string{"KEY1", "KEY2", "KEY3"}, "KEY1, KEY2, KEY3"},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			assert.Equal(s.T(), tc.expected, formatKeys(tc.keys))
		})
	}
}

// --- Test Methods for CheckKeySuite ---
func (s *CheckKeySuite) TestCheckKey() {
	cfg := NewConfigImpl()

	strKey := Variable[string]("MY_STRING")
	intKey := Variable[int]("MY_INT")
	boolKey := Variable[bool]("MY_BOOL")
	float32Key := Variable[float32]("MY_FLOAT32")

	missingStrKey := Variable[string]("MISSING_STRING")
	missingIntKey := Variable[int]("MISSING_INT")
	uintptrKey := Variable[uintptr]("MY_UINTPTR_UNINIT_MAP_SCENARIO")

	cfg.RegString[strKey] = "value"
	cfg.RegInt[intKey] = 123
	cfg.RegBool[boolKey] = true
	cfg.RegFloat32[float32Key] = 3.14

	s.Run("ExistingKeys", func() {
		name, exists := cfg.checkKey(strKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(strKey), name)

		name, exists = cfg.checkKey(intKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(intKey), name)

		name, exists = cfg.checkKey(boolKey)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(boolKey), name)

		name, exists = cfg.checkKey(float32Key)
		assert.True(s.T(), exists)
		assert.Equal(s.T(), string(float32Key), name)
	})

	s.Run("MissingKeys", func() {
		name, exists := cfg.checkKey(missingStrKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(missingStrKey), name)

		name, exists = cfg.checkKey(missingIntKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(missingIntKey), name)

		name, exists = cfg.checkKey(uintptrKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(uintptrKey), name)
	})

	s.Run("DifferentKeyTypeSameName", func() {
		diffTypeSameNameKey := Variable[string]("MY_INT")
		name, exists := cfg.checkKey(diffTypeSameNameKey)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(diffTypeSameNameKey), name)

		diffTypeSameNameKey2 := Variable[int]("MY_STRING")
		name, exists = cfg.checkKey(diffTypeSameNameKey2)
		assert.False(s.T(), exists)
		assert.Equal(s.T(), string(diffTypeSameNameKey2), name)
	})
}

// --- Test Methods for ConfigurationKeysRegisteredSuite ---

func (s *ConfigurationKeysRegisteredSuite) TestConfigurationKeysRegistered() {
	cfg := NewConfigImpl()

	strKey1 := Variable[string]("STR_KEY_1")
	strKey2Missing := Variable[string]("STR_KEY_2_MISSING")
	intKey1 := Variable[int]("INT_KEY_1")
	floatKeyMissing := Variable[float32]("FLOAT_KEY_MISSING")

	cfg.RegString[strKey1] = "val1"
	cfg.RegInt[intKey1] = 100

	s.Run("AllCheckedKeysExist", func() {
		err := cfg.ConfigurationKeysRegistered(strKey1, intKey1)
		assert.NoError(s.T(), err)
	})

	s.Run("SomeKeysMissing", func() {
		err := cfg.ConfigurationKeysRegistered(strKey1, strKey2Missing, intKey1, floatKeyMissing)
		s.Require().Error(err)

		var missingErr missingVariableError
		s.Require().True(errors.As(err, &missingErr), "Error should be of type missingVariableError")

		s.Require().ErrorIs(err, ErrMissingVariable, "Error should unwrap to ErrMissingVariable")

		assert.ElementsMatch(s.T(), []string{string(strKey2Missing), string(floatKeyMissing)}, missingErr.Keys)
		assert.Contains(s.T(), err.Error(), "missing configuration variables:")
		assert.Contains(s.T(), err.Error(), string(strKey2Missing))
		assert.Contains(s.T(), err.Error(), string(floatKeyMissing))
	})

	s.Run("AllCheckedKeysMissing", func() {
		missingStr := Variable[string]("COMPLETELY_MISSING_S")
		missingInt := Variable[int]("COMPLETELY_MISSING_I")
		err := cfg.ConfigurationKeysRegistered(missingStr, missingInt)
		s.Require().Error(err)
		var missingErr missingVariableError
		s.Require().True(errors.As(err, &missingErr))
		assert.ElementsMatch(s.T(), []string{string(missingStr), string(missingInt)}, missingErr.Keys)
		s.Require().ErrorIs(err, ErrMissingVariable)
	})

	s.Run("NoKeysToCheck", func() {
		err := cfg.ConfigurationKeysRegistered()
		assert.NoError(s.T(), err)
	})

	s.Run("ErrorTypeAndUnwrap", func() {
		err := cfg.ConfigurationKeysRegistered(Variable[string]("ANY_MISSING_KEY"))
		s.Require().Error(err)

		_, ok := err.(missingVariableError)
		assert.True(s.T(), ok, "Error should be missingVariableError type")

		assert.ErrorIs(s.T(), err, ErrMissingVariable, "Error should unwrap to ErrMissingVariable via errors.Is")

		unwrappedErr := errors.Unwrap(err)
		assert.Equal(s.T(), ErrMissingVariable, unwrappedErr, "Unwrapped error should be exactly ErrMissingVariable")
	})
}

// --- Main Test Runner ---

func TestConfiguraSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
	suite.Run(t, new(LoadEnvironmentSuite))
	suite.Run(t, new(FormatKeysSuite))
	suite.Run(t, new(CheckKeySuite))
	suite.Run(t, new(ConfigurationKeysRegisteredSuite))
}
