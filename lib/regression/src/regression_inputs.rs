use serde::Deserialize;
use nalgebra::DMatrix;

#[derive(Deserialize)]
pub struct RegressionInputs {
    x_values: Vec<Vec<f32>>,
    y_values: Vec<f32>
}


impl RegressionInputs {
    pub fn get_as_matrix(&self) -> (DMatrix<f32>, DMatrix<f32>) {
        let mut x_values_flattened = Vec::new();
        let y_matrix = DMatrix::from_row_vector(self.y_values.len(), 1, &self.y_values.clone());

        for coord in self.x_values.iter() {
            for val in coord {
                x_values_flattened.push(*val);
            }

            x_values_flattened.push(1.)
        }

        let x_matrix = DMatrix::from_row_vector(self.x_values.len(), self.x_values[0].len() + 1, &x_values_flattened);

        return (x_matrix, y_matrix)
    }
}