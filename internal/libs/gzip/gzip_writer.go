package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
)

// Write は文字列 value を gzip 圧縮して書き込む
func Write(value string) (string, error) {

	// byte buffer に書き込む
	buf := bytes.Buffer{}
	// gzip writer の生成
	gw := gzip.NewWriter(&buf)
	// 生の値を書き込む（gzip 圧縮する）
	_, err := gw.Write([]byte(value))
	if err != nil {
		_ = gw.Close()
		return "", fmt.Errorf("failed to write. error=[%w]", err)
	}

	err = gw.Flush()
	if err != nil {
		_ = gw.Close()
		return "", fmt.Errorf("failed to flush. error=[%w]", err)
	}
	// compress 後の値を返したい場合は
	// defer ではなく return　前に閉じる
	_ = gw.Close()
	return buf.String(), nil
}

// Read は gzip 圧縮された文字列 value の読み込み
func Read(value string) (string, error) {
	// 読み込みたいgzip圧縮された値を byte buffer に変換
	buf := bytes.NewBuffer([]byte(value))
	// gzip reader の生成
	gr, err := gzip.NewReader(buf)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader. error=[%s]", err.Error())
	}
	defer gr.Close()

	// 解凍した値を保存する buffer
	result := bytes.Buffer{}
	// gzip 圧縮された文字の読み込み、解凍
	_, err = result.ReadFrom(gr)
	if err != nil {
		return "", fmt.Errorf("failed to read. error=[%w]", err)
	}

	return result.String(), nil
}

// HTTPWriter は HTTP Response の書き出し
func HTTPWriter(w http.ResponseWriter, statusCode int, headerMap map[string]string, body string) {
	// Header などの書き込み設定がしたければ
	for k, v := range headerMap {
		w.Header().Add(k, v)
	}

	// status code の書き込み設定
	w.WriteHeader(statusCode)
	gw := gzip.NewWriter(w)
	defer gw.Close()

	// body の書き込み、表示
	_, err := gw.Write([]byte(body))
	if err != nil {
		_ = fmt.Errorf("failed to write. error=[%w]", err)
	}
}
