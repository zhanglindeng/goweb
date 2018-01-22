package util

import "fmt"

func SizeFormat(size int) string {

	size2 := float64(size)

	if size2 < 1024 {
		return fmt.Sprintf("%d B", size)
	}

	if size2 < 1024*1024 {
		return fmt.Sprintf("%0.2f KB", size2/1024)
	}

	if size2 < 1024*1024*1024 {
		return fmt.Sprintf("%0.2f MB", size2/(1024*1024))
	}

	if size2 < 1024*1024*1024*1024 {
		return fmt.Sprintf("%0.2f GB", size2/(1024*1024*1024))
	}

	return fmt.Sprintf("%0.2f TB", size2/(1024*1024*1024*1024))
}

