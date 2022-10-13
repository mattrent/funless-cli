// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fn

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/funlessdev/fl-cli/pkg/client"
	"github.com/funlessdev/fl-cli/pkg/log"
	swagger "github.com/funlessdev/fl-client-sdk-go"
)

type (
	Fn struct {
		Invoke Invoke `cmd:"" help:"todo fn invoke help"`
		Create Create `cmd:"" help:"todo fn create help"`
		Delete Delete `cmd:"" help:"todo fn delete help"`
	}

	Create struct {
		Name      string `arg:"" name:"name" help:"name of the function to create"`
		Namespace string `name:"namespace" short:"n" help:"namespace of the function to create"`
		Source    string `name:"source" required:"" short:"s" type:"existingFile" help:"path of the source file"`
		Language  string `name:"language" short:"l" help:"programming language of the function"`
	}

	Invoke struct {
		Name      string            `arg:"" name:"name" help:"name of the function to invoke"`
		Namespace string            `name:"namespace" short:"n" help:"namespace of the function to invoke"`
		Args      map[string]string `name:"args" short:"a" help:"arguments of the function to invoke" xor:"args"`
		JsonArgs  string            `name:"json" short:"j" help:"json encoded arguments of the function to invoke; overrides args" xor:"args"`
	}

	Delete struct {
		Name      string `arg:"" name:"name" help:"name of the function to delete"`
		Namespace string `name:"namespace" short:"n" help:"namespace of the function to delete"`
	}
)

func (f *Invoke) Run(ctx context.Context, invoker client.FnHandler, logger log.FLogger) error {
	args := make(map[string]interface{}, len(f.Args))
	if f.Args != nil {
		for k, v := range f.Args {
			args[k] = v
		}
	} else if f.JsonArgs != "" {
		err := json.Unmarshal([]byte(f.JsonArgs), &args)
		if err != nil {
			return err
		}
	}
	res, err := invoker.Invoke(ctx, f.Name, f.Namespace, args)
	if err != nil {
		return extractError(err)
	}

	if res.Result != nil {
		decodedRes, err := json.Marshal(res.Result)
		if err != nil {
			return err
		}
		logger.Info(string(decodedRes))
	} else {
		return errors.New("received nil result")
	}

	return nil
}

func (f *Create) Run(ctx context.Context, invoker client.FnHandler, logger log.FLogger) error {
	code, err := os.Open(f.Source)
	if err != nil {
		return err
	}

	res, err := invoker.Create(ctx, f.Name, f.Namespace, code, f.Language)
	if err != nil {
		return extractError(err)
	}

	logger.Info(*res.Result)
	return nil
}

func (f *Delete) Run(ctx context.Context, invoker client.FnHandler, logger log.FLogger) error {
	res, err := invoker.Delete(ctx, f.Name, f.Namespace)
	if err != nil {
		return extractError(err)
	}

	logger.Info(*res.Result)
	return nil
}

func extractError(err error) error {
	openApiError, ok_sw := err.(swagger.GenericOpenAPIError)
	if ok_sw {
		switch openApiError.Model().(type) {
		case swagger.FunctionCreationError:
			specificError := openApiError.Model().(swagger.FunctionCreationError)
			return errors.New(*specificError.Error)
		case swagger.FunctionDeletionError:
			specificError := openApiError.Model().(swagger.FunctionDeletionError)
			return errors.New(*specificError.Error)
		case swagger.FunctionInvocationError:
			specificError := openApiError.Model().(swagger.FunctionInvocationError)
			return errors.New(*specificError.Error)
		}
	}
	return err
}
