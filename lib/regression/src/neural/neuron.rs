use super::{ActivationFunction, GradientFunction};
use rand::{thread_rng, Rng};

#[derive(Clone)]
pub struct Neuron {
    weights: Vec<f32>,
    bias: f32,
    input_state: Vec<f32>,
    activation_function: fn(input: f32) -> f32,
    gradient_function: fn(weights: &Vec<f32>, inputs: &Vec<f32>, diff: f32) -> Vec<f32>,
}

impl Neuron {
    pub fn init_neuron(
        num_weights: usize,
        state: Vec<f32>,
        activation_function: ActivationFunction,
        gradient_function: GradientFunction,
    ) -> Self {
        let mut rng = thread_rng();
        let mut weights: Vec<f32> = Vec::new();
        for _ in 0..num_weights {
            weights.push(rng.gen::<f32>())
        }

        return Self {
            weights,
            bias: 0f32,
            activation_function,
            input_state: state,
            gradient_function,
        };
    }

    pub fn forward_pass(&mut self, input_state: Vec<f32>) -> f32 {
        self.input_state = input_state.clone();
        return (self.activation_function)(self.dot(&input_state) + self.bias);
    }

    pub fn backward_pass(&mut self, diff: f32) {
        self.weights = (self.gradient_function)(&self.weights, &self.input_state, diff);
        self.bias += diff
    }

    fn dot(&self, input_vec: &Vec<f32>) -> f32 {
        let mut ret_val = 0f32;
        for (idx, weight) in self.weights.iter().enumerate() {
            ret_val += weight * input_vec[idx]
        }

        return ret_val;
    }
}
