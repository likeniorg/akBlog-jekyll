package util

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

// 未保存时间戳，与tar czvf 压缩出来的文件Hash值不同
func CreateTarGz(archiveName string, filesToCompress []string) error {
	// 内存缓冲区保存tar文件内容
	var tarBuffer bytes.Buffer

	// 创建tar.Writer，将文件写入tar文件内存缓冲区
	tarWriter := tar.NewWriter(&tarBuffer)

	// 添加文件到内存缓冲区
	for _, fileToCompress := range filesToCompress {
		file, err := os.Open(fileToCompress)
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		header := &tar.Header{
			Name: fileInfo.Name(),
			Mode: int64(fileInfo.Mode()),
			Size: fileInfo.Size(),
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}
	}

	tarWriter.Close()

	// 创建tar.gz文件
	tarGzFile, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer tarGzFile.Close()

	// 创建一个gzip.Writer，用于将tar文件内容压缩到.tar.gz文件
	gzipWriter := gzip.NewWriter(tarGzFile)
	defer gzipWriter.Close()

	// 将tar文件内容从内存缓冲区复制到gzip.Writer，从而创建.tar.gz文件
	_, err = io.Copy(gzipWriter, &tarBuffer)
	if err != nil {
		return err
	}

	return nil
}

func CreateTar(archiveName string, filesToCompress []string) error {
	// 创建tar文件
	tarFile, err := os.Create("all.tar")
	if err != nil {
		return err
	}
	defer tarFile.Close()

	// 创建一个tar.Writer，用于将数据写入tar文件
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	// 添加文件到tar文件
	for _, fileToCompress := range filesToCompress {
		file, err := os.Open(fileToCompress)
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		header := &tar.Header{
			Name: fileInfo.Name(),
			Mode: int64(fileInfo.Mode()),
			Size: fileInfo.Size(),
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}
	}
	return nil

}
