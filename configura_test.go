package configura

import (
	"errors"
	"os"
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

// --- Helper to initialize a new ConfigImpl for tests ---
func newTestConfigImpl() *ConfigImpl {
	return &ConfigImpl{
		RegString:  make(map[Variable[string]]string),
		RegInt:     make(map[Variable[int]]int),
		RegInt8:    make(map[Variable[int8]]int8),
		RegInt16:   make(map[Variable[int16]]int16),
		RegInt32:   make(map[Variable[int32]]int32),
		RegInt64:   make(map[Variable[int64]]int64),
		RegUint:    make(map[Variable[uint]]uint),
		RegUint8:   make(map[Variable[uint8]]uint8),
		RegUint16:  make(map[Variable[uint16]]uint16),
		RegUint32:  make(map[Variable[uint32]]uint32),
		RegUint64:  make(map[Variable[uint64]]uint64),
		RegUintptr: make(map[Variable[uintptr]]uintptr),
		RegBytes:   make(map[Variable[[]byte]][]byte),
		RegRunes:   make(map[Variable[[]rune]][]rune),
		RegFloat32: make(map[Variable[float32]]float32),
		RegFloat64: make(map[Variable[float64]]float64),
		RegBool:    make(map[Variable[bool]]bool),
	}
}

// --- Setup Methods ---

func (s *ConfigSuite) SetupTest() {
	s.config = newTestConfigImpl()
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
		// Use a slightly different key name to avoid conflicts if tests run in parallel or map is reused.
		falseKey := Variable[bool]("TEST_BOOL_FALSE_VAL")
		s.config.RegBool[falseKey] = false
		assert.False(s.T(), s.config.Bool(falseKey))
	})
}

// --- Test Methods for LoadEnvironmentSuite ---

// Helper for LoadEnvironment tests to set/unset environment variables
func (s *LoadEnvironmentSuite) setEnvVar(key string, value string) {
	err := os.Setenv(key, value)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		os.Unsetenv(key)
	})
}

func (s *LoadEnvironmentSuite) unsetEnvVar(key string) {
	err := os.Unsetenv(key)
	s.Require().NoError(err)
}

// Assuming helper functions like String, Int, etc. exist in the package for parsing.
// These would typically look like:
// func String(key Variable[string], fallback string) string { ... os.Getenv ... }
// func Int(key Variable[int], fallback int) int { ... strconv.Atoi ... }
// etc.
// The tests will verify LoadEnvironment uses them correctly.

func (s *LoadEnvironmentSuite) TestLoadString() {
	cfg := newTestConfigImpl()
	key := Variable[string]("ENV_STR")
	fallback := "fb_str"
	envVal := "env_str_val"

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), envVal)
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegString[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegString[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt() {
	cfg := newTestConfigImpl()
	key := Variable[int]("ENV_INT_VAL")
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
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegInt[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int")
		LoadEnvironment(cfg, key, fallback)
		// The Int helper (called by LoadEnvironment) should return fallback on parse error
		assert.Equal(s.T(), fallback, cfg.RegInt[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadInt8() {
	cfg := newTestConfigImpl()
	key := Variable[int8]("ENV_INT8_VAL")
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
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegInt8[key])
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-an-int8")
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegInt8[key])
	})
	s.Run("EnvVarSetOutOfBounds", func() {
		s.setEnvVar(string(key), "129") // Out of bounds for int8
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegInt8[key])
	})
}

// ... (Similar detailed tests for Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr) ...

func (s *LoadEnvironmentSuite) TestLoadFloat32() {
	cfg := newTestConfigImpl()
	key := Variable[float32]("ENV_FLOAT32_VAL")
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
		LoadEnvironment(cfg, key, fallback)
		assert.InDelta(s.T(), fallback, cfg.RegFloat32[key], 0.0001)
	})
	s.Run("EnvVarSetInvalid", func() {
		s.setEnvVar(string(key), "not-a-float")
		LoadEnvironment(cfg, key, fallback)
		assert.InDelta(s.T(), fallback, cfg.RegFloat32[key], 0.0001)
	})
}

// ... (Similar detailed tests for Float64) ...

func (s *LoadEnvironmentSuite) TestLoadBool() {
	cfg := newTestConfigImpl()
	key := Variable[bool]("ENV_BOOL_VAL")

	testCases := []struct {
		name         string
		envValue     *string // Pointer to distinguish between not set and empty string
		fallback     bool
		expectedReg  bool
		expectedHelp bool // Expected from Bool(key, fallback) direct call
	}{
		{"EnvTrueFallbackF", func(s string) *string { return &s }("true"), false, true, true},
		{"EnvFalseFallbackT", func(s string) *string { return &s }("false"), true, false, false},
		{"Env1FallbackF", func(s string) *string { return &s }("1"), false, true, true},
		{"Env0FallbackT", func(s string) *string { return &s }("0"), true, false, false},
		{"EnvTFallbackF", func(s string) *string { return &s }("t"), false, true, true},
		{"EnvFFallbackT", func(s string) *string { return &s }("f"), true, false, false},
		{"EnvInvalidFallbackF", func(s string) *string { return &s }("invalid"), false, false, false}, // strconv.ParseBool("invalid") is (false, err)
		{"EnvInvalidFallbackT", func(s string) *string { return &s }("invalid"), true, false, false},  // strconv.ParseBool("invalid") is (false, err)
		{"EnvNotSetFallbackF", nil, false, false, false},
		{"EnvNotSetFallbackT", nil, true, true, true},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cfg = newTestConfigImpl() // Reset config for each sub-test
			if tc.envValue != nil {
				s.setEnvVar(string(key), *tc.envValue)
			} else {
				s.unsetEnvVar(string(key))
			}
			LoadEnvironment(cfg, key, tc.fallback)
			assert.Equal(s.T(), tc.expectedReg, cfg.RegBool[key], "Mismatch in registered bool value")
			// This part assumes a Bool helper function like:
			// func Bool(key Variable[bool], fallback bool) bool {
			//    valStr, exists := os.LookupEnv(string(key))
			//    if !exists { return fallback }
			//    b, err := strconv.ParseBool(valStr)
			//    if err != nil { return false } // Or some other error handling for helper
			//    return b
			// }
			// For this test, let's simulate the direct helper call more accurately based on LoadEnvironment's behavior
			// LoadEnvironment uses: Bool(any(key).(Variable[bool]), any(fallback).(bool))
			// So the tc.fallback is what the internal Bool helper receives.
			// If env var exists & invalid, internal Bool helper will return false (from strconv.ParseBool error).
			// If env var does not exist, internal Bool helper returns the tc.fallback.
			// So tc.expectedReg should be the same as what direct Bool(key, tc.fallback) would yield.
			assert.Equal(s.T(), tc.expectedReg, Bool(key, tc.fallback), "Mismatch from direct Bool helper call simulation")
		})
	}
}

func (s *LoadEnvironmentSuite) TestLoadBytes() {
	cfg := newTestConfigImpl()
	key := Variable[[]byte]("ENV_BYTES_VAL")
	fallback := []byte("fb_bytes")
	envVal := []byte("env_bytes_val")

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envVal))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegBytes[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegBytes[key])
	})
}

func (s *LoadEnvironmentSuite) TestLoadRunes() {
	cfg := newTestConfigImpl()
	key := Variable[[]rune]("ENV_RUNES_VAL")
	fallback := []rune("fb_runes")
	envValStr := "env_runes_😊"
	envVal := []rune(envValStr)

	s.Run("EnvVarSet", func() {
		s.setEnvVar(string(key), string(envVal))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), envVal, cfg.RegRunes[key])
	})
	s.Run("EnvVarNotSet", func() {
		s.unsetEnvVar(string(key))
		LoadEnvironment(cfg, key, fallback)
		assert.Equal(s.T(), fallback, cfg.RegRunes[key])
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
	cfg := newTestConfigImpl()

	strKey := Variable[string]("MY_STRING")
	intKey := Variable[int]("MY_INT")
	boolKey := Variable[bool]("MY_BOOL")
	// Add a key for a type that will exist in the map
	float32Key := Variable[float32]("MY_FLOAT32")

	missingStrKey := Variable[string]("MISSING_STRING")
	missingIntKey := Variable[int]("MISSING_INT")
	// Key for a type that might not have an initialized map in a minimal config
	uintptrKey := Variable[uintptr]("MY_UINTPTR_UNINIT_MAP_SCENARIO")

	cfg.RegString[strKey] = "value"
	cfg.RegInt[intKey] = 123
	cfg.RegBool[boolKey] = true
	cfg.RegFloat32[float32Key] = 3.14

	s.Run("ExistingKeys", func() {
		keyName, ok := cfg.checkKey(strKey)
		assert.Equal(s.T(), string(strKey), keyName)
		assert.True(s.T(), ok)
		keyName, ok = cfg.checkKey(intKey)
		assert.Equal(s.T(), string(intKey), keyName)
		assert.True(s.T(), ok)
		keyName, ok = cfg.checkKey(boolKey)
		assert.Equal(s.T(), string(boolKey), keyName)
		assert.True(s.T(), ok)
		keyName, ok = cfg.checkKey(float32Key)
		assert.Equal(s.T(), string(float32Key), keyName)
		assert.True(s.T(), ok)
	})

	s.Run("MissingKeys", func() {
		keyName, ok := cfg.checkKey(missingStrKey)
		assert.Equal(s.T(), string(missingStrKey), keyName)
		assert.False(s.T(), ok, "Expected key to not exist in RegString map")
		keyName, ok = cfg.checkKey(missingIntKey)
		assert.Equal(s.T(), string(missingIntKey), keyName)
		assert.False(s.T(), ok, "Expected key to not exist in RegInt map")
		// uintptrKey would be checked against cfg.RegUintptr. Since newTestConfigImpl() initializes all maps,
		// this will correctly return false if key is not in map, not due to nil map.
		keyName, ok = cfg.checkKey(uintptrKey)
		assert.Equal(s.T(), string(uintptrKey), keyName)
		assert.False(s.T(), ok, "Expected key to not exist in RegUintptr map")
	})
}

// --- Test Methods for ConfigurationKeysRegisteredSuite ---

func (s *ConfigurationKeysRegisteredSuite) TestConfigurationKeysRegistered() {
	cfg := newTestConfigImpl()

	strKey1 := Variable[string]("STR_KEY_1")
	strKey2 := Variable[string]("STR_KEY_2_MISSING") // This key will be missing
	intKey1 := Variable[int]("INT_KEY_1")
	floatKeyMissing := Variable[float32]("FLOAT_KEY_MISSING")

	cfg.RegString[strKey1] = "val1"
	cfg.RegInt[intKey1] = 100
	// strKey2 and floatKeyMissing are not added to cfg

	s.Run("AllCheckedKeysExist", func() {
		err := cfg.ConfigurationKeysRegistered(strKey1, intKey1)
		assert.NoError(s.T(), err)
	})

	s.Run("SomeKeysMissing", func() {
		err := cfg.ConfigurationKeysRegistered(strKey1, strKey2, intKey1, floatKeyMissing)
		s.Require().Error(err)

		var missingErr missingVariableError
		s.Require().True(errors.As(err, &missingErr), "Error should be of type missingVariableError")

		// The order of keys in missingErr.Keys might not be guaranteed.
		assert.ElementsMatch(s.T(), []string{string(strKey2), string(floatKeyMissing)}, missingErr.Keys)
		assert.Contains(s.T(), err.Error(), "missing configuration variables:")
		assert.Contains(s.T(), err.Error(), string(strKey2))
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
	})

	s.Run("NoKeysToCheck", func() {
		err := cfg.ConfigurationKeysRegistered()
		assert.NoError(s.T(), err)
	})

	s.Run("ErrorTypeIsCorrect", func() {
		err := cfg.ConfigurationKeysRegistered(Variable[string]("ANY_MISSING_KEY"))
		s.Require().Error(err)
		s.ErrorIs(err, ErrMissingVariable, "The error should wrap or be ErrMissingVariable or a type that matches it")
		// More specifically, check if it's missingVariableError
		_, ok := err.(missingVariableError)
		assert.True(s.T(), ok, "Error should be missingVariableError type")
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
