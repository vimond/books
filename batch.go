package main

import "math"

// Batch can call the function f with batches of slice sl in sizes of batchSize
func Batch[D any, R any](req []D, batchSize int, f func([]D) ([]R, error)) ([]R, error) {
	numBatches := len(req) / batchSize
	ln := float64(len(req))
	resp := make([]R, 0)

	for i := 0; i <= numBatches; i++ {
		currentBatch := req[i*batchSize : int(math.Min(ln, float64((i+1)*batchSize)))]
		if len(currentBatch) == 0 {
			continue
		}

		r, err := f(currentBatch)
		if err != nil {
			return nil, err
		}

		resp = append(resp, r...)
	}

	return resp, nil
}
