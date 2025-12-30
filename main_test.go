package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		validate func(t *testing.T, result []int)
	}{
		{
			name: "размер 0",
			size: 0,
			validate: func(t *testing.T, result []int) {
				assert.NotNil(t, result)
				assert.Empty(t, result)
				assert.Len(t, result, 0)
			},
		},
		{
			name: "отрицательный размер",
			size: -1,
			validate: func(t *testing.T, result []int) {
				assert.NotNil(t, result)
				assert.Empty(t, result)
				assert.Len(t, result, 0)
			},
		},
		{
			name: "размер 1",
			size: 1,
			validate: func(t *testing.T, result []int) {
				require.NotNil(t, result)
				assert.Len(t, result, 1)
				assert.GreaterOrEqual(t, result[0], 1)
				assert.LessOrEqual(t, result[0], 1)
			},
		},
		{
			name: "небольшой положительный размер",
			size: 100,
			validate: func(t *testing.T, result []int) {
				require.NotNil(t, result)
				assert.Len(t, result, 100)

				// Проверяем диапазон
				for _, num := range result {
					assert.GreaterOrEqual(t, num, 1)
					assert.LessOrEqual(t, num, 100)
				}

				// Проверяем, что есть хотя бы два разных числа
				unique := make(map[int]bool)
				for _, num := range result {
					unique[num] = true
					if len(unique) >= 2 {
						break
					}
				}
				assert.GreaterOrEqual(t, len(unique), 2, "должны быть разные числа")
			},
		},
		{
			name: "средний размер",
			size: 50_000,
			validate: func(t *testing.T, result []int) {
				require.NotNil(t, result)
				assert.Len(t, result, 50_000)

				// Проверяем диапазон
				for _, num := range result {
					assert.GreaterOrEqual(t, num, 1)
					assert.LessOrEqual(t, num, 50_000)
				}

				// Проверяем, что есть хотя бы два разных числа
				unique := make(map[int]bool)
				for _, num := range result {
					unique[num] = true
					if len(unique) >= 2 {
						break
					}
				}
				assert.GreaterOrEqual(t, len(unique), 2, "должны быть разные числа")
			},
		},
		{
			name: "максимальный размер",
			size: 100_000_000,
			validate: func(t *testing.T, result []int) {
				require.NotNil(t, result)
				assert.Len(t, result, 100_000_000)

				// Проверяем диапазон
				for _, num := range result {
					assert.GreaterOrEqual(t, num, 1)
					assert.LessOrEqual(t, num, 100_000_000)
				}

				// Проверяем, что есть хотя бы два разных числа
				unique := make(map[int]bool)
				for _, num := range result {
					unique[num] = true
					if len(unique) >= 2 {
						break
					}
				}
				assert.GreaterOrEqual(t, len(unique), 2, "должны быть разные числа")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := generateRandomElements(tc.size)
			tc.validate(t, result)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "слайс с одним элементом",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "несколько положительных чисел",
			input:    []int{1, 3, 5, 2, 4, 45, 1200},
			expected: 1200,
		},
		{
			name:     "одинаковые числа",
			input:    []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "максимум в начале",
			input:    []int{100, 1, 2, 3, 4},
			expected: 100,
		},
		{
			name:     "максимум в конце",
			input:    []int{1, 2, 3, 4, 100},
			expected: 100,
		},
		{
			name:     "максимум в середине",
			input:    []int{1, 2, 100, 3, 4},
			expected: 100,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maximum(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestMaximum_EdgeCases(t *testing.T) {
	t.Run("nil слайс считается пустым слайсом", func(t *testing.T) {
		result := maximum(nil)
		assert.Equal(t, 0, result)
	})

	t.Run("очень большой слайс", func(t *testing.T) {
		// Создаем слайс с 1 миллионом элементов
		size := 1_000_000
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = i
		}
		// Устанавливаем максимум в конце
		data[size-1] = size * 2

		result := maximum(data)
		assert.Equal(t, size*2, result)
	})
}
