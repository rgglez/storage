/*
Copyright 2024 Rodolfo Gonzalez Gonzalez.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/kr/pretty"
	_ "github.com/rgglez/go-storage/services/oss/v3"
	services "github.com/rgglez/go-storage/v5/services"
	"github.com/rgglez/go-storage/v5/types"
	tracerr "github.com/ztrue/tracerr"
)

//-----------------------------------------------------------------------------

type Storage struct {
	cnn   string
	store types.Storager
}

//-----------------------------------------------------------------------------

func NewStorage(cnn string) *Storage {
	// Here we replace env vars in case they are defined in the connection
	// string (cnn) by using the $ prefix ($VAR for instance).
	// Define the regular expression to match the substrings that start with $
	re := regexp.MustCompile(`\$(\w+)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(cnn, -1)

	// Create the replacements map using the extracted keys
	replacements := make(map[string]string)
	for _, match := range matches {
		if len(match) > 1 {
			key := match[1]
			replacements[key] = os.Getenv(key)
		}
	}

	// Replace the matches with the corresponding values
	result := re.ReplaceAllStringFunc(cnn, func(m string) string {
		// Extract the key (remove the leading $)
		key := strings.TrimPrefix(m, "$")
		// Return the replacement value if it exists, otherwise return the original match
		if val, ok := replacements[key]; ok {
			return val
		}
		return m
	})

	store, err := services.NewStoragerFromString(result)
	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	return &Storage{
		cnn:   result,
		store: store,
	}
}

//-----------------------------------------------------------------------------

func (s *Storage) Read(objectName string, filePath string) (err error) {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println("===============================================================")
			pretty.Println(err)
			fmt.Println("===============================================================")
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			fmt.Println("===============================================================")
			os.Exit(1)
		}
	}()

	ctx := context.Background()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	sizeB, err := s.store.ReadWithContext(ctx, objectName, file)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	fmt.Printf("Successfully downloaded %d bytes to %s\n", sizeB, filePath)
	return nil
}

//-----------------------------------------------------------------------------

func (s *Storage) Write(filePath string, objectName string) (err error) {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println("===============================================================")
			pretty.Println(err)
			fmt.Println("===============================================================")
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			fmt.Println("===============================================================")
			os.Exit(1)
		}
	}()

	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get the file info (size, etc.)
	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	_, err = s.store.WriteWithContext(ctx, objectName, file, fileStat.Size())
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filePath, objectName)
	return nil
}

//-----------------------------------------------------------------------------

func (s *Storage) ReadWithSignedURL(path string, expire time.Duration) (req *http.Request, err error) {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println("===============================================================")
			pretty.Println(err)
			fmt.Println("===============================================================")
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			fmt.Println("===============================================================")
			os.Exit(1)
		}
	}()

	// QuerySignHTTPRead needs at least two arguments.
	// `path` is the path of object.
	// `expire` provides the time period, with type time.Duration, for which the generated req.URL is valid.
	//
	// QuerySignHTTPRead will return two values.
	// `req` is the generated `*http.Request`, `req.URL` specifies the URL to access with signature in the query string. And `req.Header` specifies the HTTP headers included in the signature.
	// `err` is the error during this operation.
	req, err = s.store.QuerySignHTTPRead(path, expire)
	if err != nil {
		return nil, fmt.Errorf("read %v: %v", path, err)
	}	

	return req, nil
}

//-----------------------------------------------------------------------------

func (s *Storage) WriteWithSignedURL(path string, expire time.Duration, size int64)  (req *http.Request, err error) {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println("===============================================================")
			pretty.Println(err)
			fmt.Println("===============================================================")
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			fmt.Println("===============================================================")
			os.Exit(1)
		}
	}()

	// QuerySignHTTPWrite needs at least three arguments.
	// `path` is the path of object.
	// `size` is the length, in bytes, of the data for uploading.
	// `expire` provides the time period, with type time.Duration, for which the generated req.URL is valid.
	//
	// QuerySignHTTPWrite will return two values.
	//
	// `req` is the generated `*http.Request`:
	// `req.URL` specifies the URL to access with signature in the query string.
	// `req.Header` specifies the HTTP headers included in the signature.
	// `req.ContentLength` records the length of the associated content, the value equals to `size`.
	//
	// `err` is the error during this operation.
	req, err = s.store.QuerySignHTTPWrite(path, size, expire)
	if err != nil {
		return nil, fmt.Errorf("write %v: %v", path, err)
	}

	return req, nil
}