//go:build sspl && !cli

package compression

// zstandardSupportStatus returns Zstandard compression support status.
func zstandardSupportStatus() AlgorithmSupportStatus {
	return AlgorithmSupportStatusSupported
}
