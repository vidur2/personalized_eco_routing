use super::{neuron::Neuron, ActivationFunction, GradientFunction};


#[derive(Clone)]
pub struct Layer<'a>(Vec<Neuron<'a>>);

impl<'a> Layer<'a> {
    pub fn init_layer(
        amt_neurons: usize,
        amt_weights: usize,
        state: &'a Vec<f32>,
        activation_function: ActivationFunction,
        gradient_function: GradientFunction,
    ) -> Self {
        let mut layer = Vec::new();
        for _ in 0..amt_neurons {
            layer.push(Neuron::init_neuron(
                amt_weights,
                state,
                activation_function,
                gradient_function,
            ))
        }

        return Self(layer);
    }

    pub fn forward_pass(&mut self, input_state: &'a Vec<f32>) -> Vec<f32> {
        let mut outputs = Vec::new();
        for neuron in self.0.iter_mut() {
            outputs.push(neuron.forward_pass(input_state))
        }

        return outputs;
    }

    pub fn backward_pass(&mut self, diff: f32) {
        for neuron in self.0.iter_mut() {
            neuron.backward_pass(diff)
        }
    }
}
