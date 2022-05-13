mod regression_inputs;
mod regression;

use std::ffi::{CStr, CString};
use nalgebra::DMatrix;
use regression_inputs::RegressionInputs;
use regression::Regression;


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
pub extern "C" fn train_regression(regression_input: *const libc::c_char) -> *const libc::c_char{
    let buf = unsafe {
        CStr::from_ptr(regression_input).to_bytes()
    };

    let regression: RegressionInputs = serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();
    let (a, b) = regression.get_as_matrix();
    let regr = Regression::train(a, b);

    return CString::new(serde_json::to_string(&regr).unwrap()).unwrap().into_raw()
}

#[no_mangle]
pub extern "C" fn predict_regression(x_var: *const libc::c_char, regression_model: *const libc::c_char) -> *const libc::c_char {
    let mut buf = unsafe {
        CStr::from_ptr(x_var).to_bytes()
    };

    let mut ret_val: Vec<f32> = Vec::new();

    let x_val_parsed: Vec<Vec<f32>> = serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();

    buf = unsafe {
        CStr::from_ptr(regression_model).to_bytes()
    };

    let regression: Regression = serde_json::from_str(&String::from_utf8(buf.to_vec()).unwrap()).unwrap();

    for x in x_val_parsed.iter() {
        ret_val.push(regression.predict(x).unwrap())
    }
    return CString::new(serde_json::to_string(&ret_val).unwrap()).unwrap().into_raw()
}