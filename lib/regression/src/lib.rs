mod neural;
mod regression;
mod supervised_inputs;

use regression::Regression;
use std::ffi::{CStr, CString};
use supervised_inputs::SupervisedInputs;

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

#[no_mangle]
pub extern "C" fn train_neural_net(data_ptr: *const libc::c_char) {
    let data_buf = unsafe { CStr::from_ptr(data_ptr).to_bytes() };
    let data_str = String::from_utf8(data_buf.to_vec()).unwrap();
    let inputs: SupervisedInputs = serde_json::from_str(&data_str).unwrap();
}

#[no_mangle]
pub extern "C" fn print(name: *const libc::c_char) {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("{}", str_name);
}
