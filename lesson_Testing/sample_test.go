package main

import (
	"testing"
)

func TestMultiple(t *testing.T) {
	t.Run("groupA", func(t *testing.T) {

		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	t.Run("groupB", func(t *testing.T) {

		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	t.Run("groupC", func(t *testing.T) {

		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	t.Run("groupD", func(t *testing.T) {

		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")

			var x, y, result = 222, 222, 49284

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	//var x, y, result = 2, 2, 4
	//
	//realResult := Multiple(x, y)
	//
	//if realResult != result {
	//	t.Errorf("%d != %d", result, realResult)
	//}
}
