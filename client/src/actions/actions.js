import { GET_BILL, GET_BILLS, CREATE_BILL, DELETE_BILL, UPDATE_BILL } from './types';
import axios from 'axios';

const api_base_url = 'http://127.0.0.1:8080/api/v1/bills/';

export const getBill = (id) => dispatch => {
  axios.get(api_base_url + id, {dataType:'json'})
    .then((response) => dispatch({
      type: GET_BILL,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err);
    })
};

export const getBills = () => dispatch => {
  axios.get(api_base_url, {dataType:'json'})
    .then((response) => dispatch({
      type: GET_BILLS,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err);
    })

};

export const createBill = (billData) => dispatch => {
  axios.post(api_base_url, {
    body: billData
  })
    .then((response) => dispatch({
      type: CREATE_BILL,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err)
    });
};

export const deleteBill = (id) => dispatch => {
  axios.delete(api_base_url + id, {datatype:'json'})
    .then((response) => dispatch({
      type: DELETE_BILL,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err)
    })
};

export const updateBill = (id) => dispatch => {
  axios.put(api_base_url + id, {datatype:'json'})
    .then((response) => dispatch({
      type: UPDATE_BILL,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err)
    })
};