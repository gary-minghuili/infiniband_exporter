package global

var (
	HcaMlxMap = map[string]string{
		"HCA-1":  "mlx5_0",
		"HCA-2":  "mlx5_1",
		"HCA-3":  "mlx5_10",
		"HCA-4":  "mlx5_11",
		"HCA-5":  "mlx5_2",
		"HCA-6":  "mlx5_3",
		"HCA-11": "mlx5_8",
		"HCA-12": "mlx5_9",
	}
	MlxLeafMap = map[string]string{
		"mlx5_8":  "LEAF01",
		"mlx5_9":  "LEAF02",
		"mlx5_10": "LEAF03",
		"mlx5_11": "LEAF04",
		"mlx5_2":  "LEAF05",
		"mlx5_0":  "LEAF06",
		"mlx5_1":  "LEAF07",
		"mlx5_3":  "LEAF08",
	}
)
