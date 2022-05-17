use serde_closure::{traits::Fn};

pub(crate) mod layer;
mod neuron;
pub(crate) mod nn;
pub (crate) mod neural_inputs;


pub type ActivationFunction = fn(input: f32) -> f32;
pub type GradientFunction = fn(weights: &Vec<f32>, inputs: &Vec<f32>, diff: f32) -> Vec<f32>;
pub type LossFunction = fn(output: &Vec<f32>, expected_output: &Vec<f32>) -> f32;
