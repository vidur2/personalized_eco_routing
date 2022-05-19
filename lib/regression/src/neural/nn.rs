use super::{layer::Layer, ActivationFunction, GradientFunction, LossFunction};

pub struct NeuralNet {
    net: Vec<Layer>,
    loss_function: LossFunction,
}

impl NeuralNet {
    pub fn init_net(
        amt_layers: usize,
        amt_neurons_per_layer: Vec<usize>,
        amt_weights_per_neuron_by_layer: Vec<usize>,
        state: Vec<f32>,
        activation_function: Vec<ActivationFunction>,
        gradient_function: Vec<GradientFunction>,
        loss_function: LossFunction,
    ) -> Self {
        let mut layers = Vec::new();
        for layer_id in 0..amt_layers {
            layers.push(Layer::init_layer(
                amt_neurons_per_layer[layer_id],
                amt_weights_per_neuron_by_layer[layer_id],
                state.clone(),
                activation_function[layer_id],
                gradient_function[layer_id],
            ));
        }
        return Self {
            net: layers,
            loss_function,
        };
    }

    pub fn predict(&mut self, input_state: Vec<f32>) -> Vec<f32> {
        let mut prev_layer = input_state;
        for layer in self.net.iter_mut() {
            prev_layer = layer.forward_pass(prev_layer);
        }

        return prev_layer;
    }

    pub fn full_pass(&mut self, input_state: Vec<f32>, expected_output: Vec<f32>) {
        let output = self.predict(input_state);
        let diff = (self.loss_function)(&output, &expected_output);
        for layer in self.net.iter_mut() {
            layer.backward_pass(diff)
        }
    }
}
