// mod neural;
mod regression;
mod supervised_inputs;

use regression::Regression;
use std::ffi::{CStr, CString};
use supervised_inputs::SupervisedInputs;
// use neural::neural_inputs::NeuralInputs;
// use neural::nn::NeuralNet;

#[no_mangle]
pub extern "C" fn init_lib() {
    env_logger::init();

    log::trace!("init_stuff() trace level message");
    log::debug!("init_stuff() debug level message");
    log::info!("init_stuff() info level message");
    log::warn!("init_stuff() warn level message");
    log::error!("init_stuff() error level message");
}

#[no_mangle]
pub extern "C" fn train_regression(regression_input: *const libc::c_char) -> *const libc::c_char {
    let buf = unsafe { CStr::from_ptr(regression_input).to_bytes() };

    let regression: SupervisedInputs =
        serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();
    let (a, b) = regression.get_as_matrix();
    let regr = Regression::train(a, b);

    return CString::new(serde_json::to_string(&regr).unwrap())
        .unwrap()
        .into_raw();
}

#[no_mangle]
pub extern "C" fn predict_regression(
    x_var: *const libc::c_char,
    regression_model: *const libc::c_char,
) -> *const libc::c_char {
    let mut buf = unsafe { CStr::from_ptr(x_var).to_bytes() };

    let mut ret_val: Vec<f32> = Vec::new();

    let x_val_parsed: Vec<Vec<f32>> =
        serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();

    buf = unsafe { CStr::from_ptr(regression_model).to_bytes() };

    let regression: Regression =
        serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();

    for x in x_val_parsed.iter() {
        ret_val.push(regression.predict(x).unwrap())
    }
    return CString::new(serde_json::to_string(&ret_val).unwrap())
        .unwrap()
        .into_raw();
}

// #[no_mangle]
// pub extern "C" fn train_neural_net(data_ptr: *const libc::c_char) -> *const libc::c_char {
//     let data_buf = unsafe { CStr::from_ptr(data_ptr).to_bytes() };
//     let data_str = String::from_utf8(data_buf.to_vec()).unwrap();
//     let mut inputs: NeuralInputs = serde_json::from_str(&data_str).unwrap();
//     for (idx, y_coord) in inputs.datapoints.y_values.iter_mut().enumerate() {
//         let x_coord = inputs.datapoints.x_values[idx];
//         inputs.nn.train(x_coord, vec![y_coord]);
//     }

//     return CString::new(serde_json::to_string(&inputs.nn).unwrap()).unwrap().into_raw()
// }

#[no_mangle]
pub extern "C" fn print(name: *const libc::c_char) {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("{}", str_name);
}

#[cfg(test)]
mod test {
    use crate::{Regression, train_regression, predict_regression};
    use crate::{CStr, CString};

    #[test]
    fn test_train_regression() {
        let test_string = String::from(r#"{
            "x_values": [[1,2], [2,3]],
            "y_values": [2,5]
        }"#);
        let mut cstr = CString::new(test_string).unwrap().into_raw();
        let regr_model = train_regression(cstr);
        cstr = CString::new("[[3,1]]").unwrap().into_raw();
        predict_regression(cstr, regr_model);
    }
}
