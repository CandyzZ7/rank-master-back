package utils

import "strconv"

// Str2Int string to int
func Str2Int(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

// Str2IntE string to int with error
func Str2IntE(str string) (int, error) {
	return strconv.Atoi(str)
}

// Str2Uint32 string to uint32
func Str2Uint32(str string) uint32 {
	v, _ := strconv.ParseUint(str, 10, 64)
	return uint32(v)
}

// Str2Uint32E string to uint32 with error
func Str2Uint32E(str string) (uint32, error) {
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint32(v), nil
}

// Str2Int64 string to int64
func Str2Int64(str string) int64 {
	v, _ := strconv.ParseInt(str, 10, 64)
	return v
}

// Str2Int64E string to int64 with error
func Str2Int64E(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// Str2Uint64 string to uint64
func Str2Uint64(str string) uint64 {
	v, _ := strconv.ParseUint(str, 10, 64)
	return v
}

// Str2Uint64E string to uint64 with error
func Str2Uint64E(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

// Str2Float32 string to float32
func Str2Float32(str string) float32 {
	v, _ := strconv.ParseFloat(str, 32)
	return float32(v)
}

// Str2Float32E string to float32 with error
func Str2Float32E(str string) (float32, error) {
	v, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

// Str2Float64 string to float64
func Str2Float64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

// Str2Float64E string to float64 with error
func Str2Float64E(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Int2Str int to string
func Int2Str(v int) string {
	return strconv.Itoa(v)
}

// Uint642Str uint64 to string
func Uint642Str(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Int642Str int64 to string
func Int642Str(v int64) string {
	return strconv.FormatInt(v, 10)
}
