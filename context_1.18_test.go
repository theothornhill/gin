// Copyright 2021 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !go1.19

ackage gin

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextFormFileFailed18(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	defer func(mw *multipart.Writer) {
		err := mw.Close()
		if err != nil {
			assert.Error(t, err)
		}
	}(mw)
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())
	c.engine.MaxMultipartMemory = 8 << 20
	assert.Panics(t, func() {
		f, err := c.FormFile("file")
		assert.Error(t, err)
		assert.Nil(t, f)
	})
}
