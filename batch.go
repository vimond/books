package books

import "math"

// Batch can call the function f with batches of slice sl in sizes of batchSize.
// If failOnErr is true, it returns all responses and errors so far upon the first returned error from the function.
func Batch[D any, R any](req []D, f func([]D) (R, error), batchSize int, failOnErr bool) ([]R, []error) {
	numBatches := int(math.Ceil(float64(len(req)) / float64(batchSize)))
	resp := make([]R, 0, numBatches)
	errs := make([]error, 0, numBatches)

	for i := 0; i < len(req); i += batchSize {
		r, err := f(req[i:Min(i+batchSize, len(req))])
		resp = append(resp, r)
		errs = append(errs, err)

		if err != nil && failOnErr {
			return resp, errs
		}
	}

	return resp, errs
}
