use nalgebra::DMatrix;
use nalgebra::{Inverse, Transpose};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Regression {
    coeffs: Vec<f32>,
    constant: f32,
}

impl Regression {
    pub fn predict(&self, input_values: &Vec<f32>) -> Option<f32> {
        if input_values.len() != self.coeffs.len() {
            return None;
        } else {
            let mut output = 0f32;

            for (idx, value) in input_values.iter().enumerate() {
                output += self.coeffs[idx] * value
            }

            return Some(output + self.constant);
        }
    }

    pub fn train(a: DMatrix<f32>, b: DMatrix<f32>) -> Self {
        let at = a.transpose();
        let ls = at.clone() * a;
        let rs = at * b;
        let outputs = ls.inverse().unwrap() * rs;
        let mut coeffs = outputs.into_vector();

        return Self {
            constant: coeffs.pop().unwrap(),
            coeffs,
        };
    }
}
