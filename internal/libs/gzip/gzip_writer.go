package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

// gzipWrite は文字列 value を gzip 圧縮して書き込む
func gzipWrite(value string) (string, error) {

	// byte buffer に書き込む
	buf := bytes.NewBuffer([]byte(""))
	// gzip writer の生成
	gw := gzip.NewWriter(buf)
	// 生の値を書き込む（gzip 圧縮する）
	_, err := gw.Write([]byte(value))
	if err != nil {
		return "", fmt.Errorf("failed to write. error=[%s].", err.Error())
	}

	return buf.String(), nil
}

// gzipRead は gzip 圧縮された文字列 value の読み込み
func gzipRead(value string) (string, error) {
	// 読み込みたいgzip圧縮された値を byte buffer に変換
	buf := bytes.NewBuffer([]byte(value))
	// gzip reader の生成
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader. error=[%s].", err.Error())
	}
	defer gr.Close()

	// 解凍した値を保存する buffer
	result := bytes.Buffer{}
	// gzip 圧縮された文字の読み込み、解凍
	_, err = result.ReadFrom(gr)
	if err != nil {
		return "", fmt.Errorf("failed to read. error=[%s].", err.Error())
	}

	return result.String(), nil
}
