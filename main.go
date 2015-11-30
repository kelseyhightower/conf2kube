// Copyright 2015 Google, Inc All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Secret holds a Kubernetes secret.
type Secret struct {
	APIVersion string                 `json:"apiVersion"`
	Data       map[string]string      `json:"data"`
	Kind       string                 `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	Type       string                 `json:"type"`
}

func main() {
	configFilePath := flag.String("f", "", "Path to configuration `file`.")
	name := flag.String("n", "", "The `name` to use for the Kubernetes secret. Defaults to basename of configuration file.")
	flag.Parse()

	log.SetFlags(0)

	// With no options set read a JSON formated Kubernetes secret from stdin and
	// print the contents to stdout. The secret data must match the secret name.
	if *configFilePath == "" && *name == "" {
		var s Secret
		decoder := json.NewDecoder(os.Stdin)
		if err := decoder.Decode(&s); err != nil {
			log.Fatal(err)
		}
		configName := s.Metadata["name"].(string)
		data, err := base64.StdEncoding.DecodeString(s.Data[configName])
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data)
		os.Exit(0)
	}

	// Format the contents of the given configuration file and print a JSON
	// formated Kubernetes secret to stdout.
	configFileData, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if *name == "" {
		*name = path.Base(*configFilePath)
	}
	data := map[string]string{*name: base64.StdEncoding.EncodeToString(configFileData)}
	metadata := map[string]interface{}{"name": name}
	out, err := json.Marshal(&Secret{"v1", data, "Secret", metadata, "Opaque"})
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(out)
}
