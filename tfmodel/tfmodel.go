package tfmodel

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	tensorflow_proto "github.com/bigchange/go-pro/tensorflow"
	tensorflow_serving_proto "github.com/bigchange/go-pro/tensorflow_serving"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// TfModel ...
type TfModel struct {
	Model *tf.SavedModel
}

// Init model with a dir including model file and variables
func Init(path string, tags []string) *TfModel {
	model, err := tf.LoadSavedModel(path, tags, nil)
	if err != nil {
		fmt.Printf("Error loading Saved Model:%v", err.Error())
		return nil
	}
	operations := model.Graph.Operations()
	for _, op := range operations {
		name := op.Name()
		if strings.HasPrefix(name, "sentence") || strings.HasPrefix(name, "dropout") || strings.HasPrefix(name, "inference") {
			log.Println("op name:", op.Name(), "NumOutputs:", op.NumOutputs())
		}
	}
	log.Println("op loading finished")
	return &TfModel{Model: model}
}

// Init model with a already freezed graph
func InitModel(modelDir string, modelFile string) *TfModel {
	modelpath := filepath.Join(modelDir, modelFile)
	model, err := ioutil.ReadFile(modelpath)
	if err != nil {
		log.Fatal(err)
	}

	// Construct an in-memory graph from the serialized form.
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal(err)
	}
	operations := graph.Operations()
	for _, op := range operations {
		log.Println("op name:", op.Name(), "NumOutputs:", op.NumOutputs())
	}
	// Create a session for inference over graph.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &TfModel{Model: &tf.SavedModel{Graph: graph, Session: session}}
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

// llb WorkTagger Model
// Test_Status: passed
func TfModelLoadAndEval_work_tagger() {
	const max_sentence_len_ int = 199
	var input = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	input_data := make([][]int32, 2)
	input_tokens := make([]int32, max_sentence_len_)
	var lengthInput = len(input)
	if lengthInput < max_sentence_len_ {
		for i := 0; i < lengthInput; i++ {
			input_tokens[i] = input[i]
		}
		for i := lengthInput; i < max_sentence_len_; i++ {
			input_tokens[i] = 0
		}
	}
	input_data[0] = input_tokens
	input_data[1] = input_tokens

	dropOutTensor, _ := tf.NewTensor(float32(0.0))
	input_tensors, _ := tf.NewTensor(input_data)
	var ab_length = make([]int32, 2)
	ab_length[0] = int32(max_sentence_len_)
	ab_length[1] = int32(max_sentence_len_)
	ab_length_tensor, _ := tf.NewTensor(ab_length)
	// passed
	/**
	model := InitModel("idmg/matching/models", "work.pbtxt")
	result, err := model.Eval([]*tf.Tensor{input_tensors}, []string{"sentence"}, "inference_final", nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	*/
	// passed
	model := Init("/Users/devops/Downloads/github/models/tf_serving/work_tagger/1", []string{"serve"})
	result, err := model.Eval([]*tf.Tensor{input_tensors, dropOutTensor, ab_length_tensor}, []string{"sentence", "dropout", "ab_length"}, "inference_final", nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	value := result[0].Value().([][]float32)
	fmt.Printf("Result value length:%v\n", len(value[0]))
}

// Test_Status: passed
func TfModelLoadAndEval_imdb_model() {
	const MAXLEN int = 20
	// 将文本转换为id序列，为了实验方便直接使用转换好的ID序列即可
	inputData := [2][MAXLEN]float32{{1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0}, {0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0}}
	tensor, err := tf.NewTensor(inputData)
	dropOutTensor, _ := tf.NewTensor(false)
	if err != nil {
		fmt.Printf("Error NewTensor: err: %s", err.Error())
		return
	}
	model := Init("/Users/devops/workspace/gitlab/idmg/lieluobo/llblibrary/models/cnnModel", []string{"myTag"})
	// model := InitModel("/Users/devops/workspace/gitlab/idmg/lieluobo/llblibrary/models/cnnModel", "saved_model.pb")
	result, err := model.Eval([]*tf.Tensor{tensor, dropOutTensor}, []string{"input_layer_input", "dropout_layer1/keras_learning_phase"}, "output_layer/Sigmoid", nil)
	if err != nil {
		fmt.Printf("Error running the session with input, err: %s  ", err.Error())
		return
	}
	value := result[0].Value().([][]float32)
	var predict = make([]float32, len(inputData))
	// 输出结果，interface{}格式
	fmt.Printf("Result value:%v\n", value)
	for i := range inputData {
		predict[i] = value[i][0]
	}
	fmt.Printf("predict value:%v", predict)
}

func TfApiTesting() {
	const MAXLEN int = 20
	inputData := []float32{1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 208.0, 659.0, 180.0, 408.0, 42.0, 547.0, 829.0, 285.0, 334.0, 42.0, 642.0, 81.0, 800.0}
	tensor, _ := tf.NewTensor(inputData)
	// dropOutTensor, _ := tf.NewTensor(false)
	tensorByte, ok := tensor.Value()
	if !ok {
		fmt.Println("Cannot type assert tensor value to ")
		return
	}
	request := tensorflow_serving_proto.PredictRequest{
		ModelSpec: &tensorflow_serving_proto.ModelSpec{
			Name:          "test",
			SignatureName: "test_sign",
		},
		Inputs: map[string]*tensorflow_proto.TensorProto{
			"input_layer_input": &tensorflow_proto.TensorProto{
				Dtype: tensorflow_proto.DataType_DT_INT32,
				TensorShape: &tensorflow_proto.TensorShapeProto{
					Dim: []*tensorflow_proto.TensorShapeProto_Dim{
						&tensorflow_proto.TensorShapeProto_Dim{
							Size: 2,
						},
						&tensorflow_proto.TensorShapeProto_Dim{
							Size: 20,
						},
					},
				},
				FloatVal: inputData,
			},
			"dropout_layer1/keras_learning_phase": &tensorflow_proto.TensorProto{
				Dtype: tensorflow_proto.DataType_DT_BOOL,
				TensorShape: &tensorflow_proto.TensorShapeProto{
					Dim: []*tensorflow_proto.TensorShapeProto_Dim{
						&tensorflow_proto.TensorShapeProto_Dim{
							Size: 1,
						},
					},
				},
				BoolVal: []bool{false},
			},
		},
		OutputFilter: []string{"output_layer/Sigmoid"},
	}
	fmt.Println("convert to tensorProto:", request)

}
