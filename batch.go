package books

import "math"

// Batch can call the function f with batches of slice sl in sizes of batchSize.
// If failOnErr is true, it returns all responses and errors so far upon the first returned error from the function.
func Batch[D any, R any](req []D, batchSize int, failOnErr bool, f func([]D) (R, error)) ([]R, []error) {
	numBatches := len(req) / batchSize
	ln := float64(len(req))
	resp := make([]R, 0, numBatches)
	errs := make([]error, 0, numBatches)

	for i := 0; i <= numBatches; i++ {
		currentBatch := req[i*batchSize : int(math.Min(ln, float64((i+1)*batchSize)))]
		if len(currentBatch) == 0 {
			continue
		}

		r, err := f(currentBatch)
		resp = append(resp, r)
		errs = append(errs, err)

		if err != nil && failOnErr {
			return resp, errs
		}
	}

	return resp, errs
}
