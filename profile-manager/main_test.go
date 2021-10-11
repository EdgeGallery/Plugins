/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"testing"
)

var ePanic = "Panic expected"
var panicProblem = "a problem"

const panicFormatString = "Panic: %v"

func TestLoadIoTConfig(t *testing.T) {
	t.Run("ReadConfigTest", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf(panicFormatString, r)
			}
		}()
		//out , _ := LoadIoTProfile()
		main()

		//assert.Equal(t, out.Config.Fledge.HostURL,  "127.0.0.10/fledge")
		//assert.Equal(t, out.Profile[2].Tdengine.STable,  "meters")
		//byte , _ := json.Marshal(out)
		//fmt.Println("Config is " + string(byte))
		//fmt.Println("Config is ", len(out.Profile))
	})
}
