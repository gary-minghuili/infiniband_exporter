package global

var (
	HcaMlxMap = map[string]string{
		"HCA-1":  "mlx5_0",
		"HCA-5":  "mlx5_0",
		"HCA-2":  "mlx5_1",
		"HCA-6":  "mlx5_1",
		"HCA-11": "mlx5_8",
		"HCA-12": "mlx5_8",
		"HCA-3":  "mlx5_10",
		"HCA-4":  "mlx5_10",
	}
	LinkMap = make(map[string]any, 0)
)
