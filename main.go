package main

import (
	"errors"
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// TfModel ...
type TfModel struct {
	Model *tf.SavedModel
}

// Init ...
func Init(path string, tags []string) *TfModel {
	model, err := tf.LoadSavedModel(path, tags, nil)
	if err != nil {
		fmt.Printf("Error loading Saved Model:%v", err.Error())
		return nil
	}
	return &TfModel{Model: model}
}

// Eval ...
func (m *TfModel) Eval(inputTensor []*tf.Tensor, inputName []string, outNames string, targets []*tf.Operation) (ret []*tf.Tensor, err error) {
	feeds := make(map[tf.Output]*tf.Tensor)
	if len(inputTensor) != len(inputName) {
		return ret, errors.New("input length not match")
	}
	for i, tensor := range inputTensor {
		feeds[m.Model.Graph.Operation(inputName[i]).Output(0)] = tensor
	}
	fetches := []tf.Output{
		m.Model.Graph.Operation(outNames).Output(0),
	}
	return m.Model.Session.Run(feeds, fetches, targets)
}

func tfModelLoadAndEval() {
	const MAXLEN int = 20
	// 将文本转换为id序列，为了实验方便直接使用转换好的ID序列即可
	inputData := [2][MAXLEN]float32{{1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0}, {0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0}}
	tensor, err := tf.NewTensor(inputData)
	dropOutTensor, _ := tf.NewTensor(false)
	if err != nil {
		fmt.Printf("Error NewTensor: err: %s", err.Error())
		return
	}
	model := Init("/Users/devops/workspace/gitlab/idmg/lieluobo/llblibrary/tfmodel/models/cnnModel", []string{"myTag"})
	result, err := model.Eval([]*tf.Tensor{tensor, dropOutTensor}, []string{"input_layer_input", "dropout_layer1/keras_learning_phase"}, "output_layer/Sigmoid", nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	value := result[0].Value().([][]float32)
	var predict = make([]float32, len(inputData))
	// 输出结果，interface{}格式
	fmt.Printf("Result value:%v", value)
	for i := range inputData {
		predict[i] = value[i][0]
	}
	fmt.Printf("predict value:%v", predict)
}

func main() {
	tfModelLoadAndEval()
}
