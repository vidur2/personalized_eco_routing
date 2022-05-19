use crate::supervised_inputs::SupervisedInputs;
use super::nn::NeuralNet;

pub struct NeuralInputs {
    datapoints: SupervisedInputs,
    nn: NeuralNet
}