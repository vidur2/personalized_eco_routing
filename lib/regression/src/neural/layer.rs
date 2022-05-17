use super::{neuron::Neuron, ActivationFunction, GradientFunction};

#[derive(Clone)]
pub struct Layer(Vec<Neuron>);

impl Layer {
    pub fn init_layer(
        amt_neurons: usize,
        amt_weights: usize,
        state: Vec<f32>,
        activation_function: ActivationFunction,
        gradient_function: GradientFunction,
    ) -> Self {
        let mut layer = Vec::new();
        for _ in 0..amt_neurons {
            layer.push(Neuron::init_neuron(
                amt_weights,
                state.clone(),
                activation_function,
                gradient_function,
            ))
        }

        return Self(layer);
    }

    pub fn forward_pass(&mut self, input_state: Vec<f32>) -> Vec<f32> {
        let mut outputs = Vec::new();
        for neuron in self.0.iter_mut() {
            outputs.push(neuron.forward_pass(input_state.clone()))
        }

        return outputs;
    }

    pub fn backward_pass(&mut self, diff: f32) {
        for neuron in self.0.iter_mut() {
            neuron.backward_pass(diff)
        }
    }
}
