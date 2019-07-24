package main

import "github.com/bigchange/go-pro/tfmodel"

func main() {
	// tfmodel.TfModelLoadAndEval_imdb_model()
	// tfmodel.TfModelLoadAndEval()
	// tfmodel.TfModelLoadAndEval_work_tagger()
	// tfmodel.TfServing_imdb_model()
	// tfmodel.TfServing_work_tagger()
	tfmodel.TfServing_bert_model()
}
