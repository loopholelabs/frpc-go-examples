/*
	Copyright 2022 Loophole Labs

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
package main

import (
	"context"
	"github.com/loopholelabs/frpc-go-examples/frpc/echo"
	"log"
	"os"
	"time"
)

func main() {
	req := new(echo.Request)
	req.Message = "Hello World"

	c, err := echo.NewClient(nil, nil)
	if err != nil {
		panic(err)
	}

	err = c.Connect(os.Args[1])
	if err != nil {
		panic(err)
	}

	for {
		log.Printf("Sending request: %s", req.Message)
		res, err := c.EchoService.Echo(context.Background(), req)
		if err != nil {
			panic(err)
		}
		if res.Message != req.Message {
			panic("invalid response")
		}
		log.Printf("Received response: %s", req.Message)

		time.Sleep(time.Millisecond * 100)
	}
}
