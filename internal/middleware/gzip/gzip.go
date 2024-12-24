package gzip

import (
	"compress/gzip"
	"github.com/Painkiller675/url_shortener_6750/internal/middleware/logger"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

// compressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
// сжимать передаваемые данные и выставлять правильные HTTP-заголовки
type compressWriter struct {
	w  http.ResponseWriter
	zw *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	return c.zw.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *compressWriter) Close() error {
	return c.zw.Close()
}

// compressReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
// декомпрессировать получаемые от клиента данные
type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}

func GzipMW(h http.Handler) http.Handler {
	gzipFunc := func(res http.ResponseWriter, req *http.Request) {
		// проверяем, что клиент отправил серверу сжатые данные в формате gzip
		// TODO go up - serveHTTP
		or := req
		contentEncoding := req.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			// оборачиваем тело запроса в io.Reader с поддержкой декомпрессии
			cr, err := newCompressReader(req.Body)
			if err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				return
			}
			// меняем тело запроса на новое
			req.Body = cr
			defer cr.Close()
		}
		//h.ServeHTTP(res, req)

		// проверяем, что клиент умеет получать от сервера сжатые данные в формате gzip
		acceptEncoding := req.Header.Get("Accept-Encoding") // это выставляет клиенот
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if !supportsGzip {
			return
		} // заретёрнить

		// TODO проверить умеет ли клинент получать NO - RETURN
		// выставил ли хэндлер это всё? Мой хендлерю Если длина ответа меньше 200 байт => не гзипуем ТУТ ПРОВЕРЯМ RES?
		if !(strings.Contains(req.Header.Get("Content-Type"), "application/json") || strings.Contains(req.Header.Get("Content-Type"), "text/html")) {
			logger.Log.Info("[INFO]", zap.String("[INFO]", "gzip IS NOT supported by the client!"), zap.String("method", req.Method), zap.String("url", req.URL.Path))
			//  continue without gzip
			h.ServeHTTP(res, or)
			return
		}
		// по умолчанию устанавливаем оригинальный http.ResponseWriter как тот,
		// который будем передавать следующей функции
		//ow := res МОЖЕТ И НЕ НАДО ЭТО
		// TODO move here
		// оборачиваем оригинальный http.ResponseWriter новым с поддержкой сжатия
		cw := newCompressWriter(res)
		// меняем оригинальный http.ResponseWriter на новый
		//ow = cw
		// не забываем отправить клиенту все сжатые данные после завершения middleware
		defer cw.Close()

		h.ServeHTTP(cw, req)

	}
	return http.HandlerFunc(gzipFunc)
}
