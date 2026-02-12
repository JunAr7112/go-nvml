/**
# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/NVIDIA/go-nvml/pkg/dl"
)

func main() {
	libName, err := getLibName()
	if err != nil {
		log.Fatalf("error getting libname: %v", err)
	}
	lib := dl.New(libName, dl.RTLD_LAZY)

	if err := lib.Open(); err != nil {
		log.Fatalf("failed to open library: %v", err)
	}
	defer func() {
		if err := lib.Close(); err != nil {
			log.Fatalf("error closing library: %v", err)
		}
		log.Printf("library closed")
	}()

	libraryPath, err := lib.Path()
	if err != nil {
		log.Fatalf("error getting library Origin: %v", err)
	}

	log.Printf("libraryPath=%v", libraryPath)

	resolvedPath, err := filepath.EvalSymlinks(libraryPath)
	if err != nil {
		log.Fatalf("error evaluating symlink: %v", err)
	}

	log.Printf("resolvedPath=%v", resolvedPath)
}

func getLibName() (string, error) {
	var libName string
	if len(os.Args) > 1 {
		libName = os.Args[1]
	}
	if libName == "" {
		return "", fmt.Errorf("no libname specified")
	}
	return libName, nil
}
