import { GET_BILL, GET_BILLS, CREATE_BILL, DELETE_BILL, UPDATE_BILL } from './types';
import axios from 'axios';
require('dotenv').config();
const api_base_url = `${process.env.API_BASE_URL}/`;

export const getBill = (id) => dispatch => {
  axios.get(api_base_url + id, {datatype:'json'})
    .then((response) => dispatch({
      type: GET_BILL,
      payload: response.data
    }))
    .catch((err) => {
      console.log(err);
    })
};

export const getBills = () => dispatch => {
  axios.get(api_base_url, {datatype:'json'})
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
    body: JSON.stringify(billData)
  })
    .then((response) => dispatch({
      type: CREATE_BILL,
      payload: response.data
    }))
    .catch((error) => {
      console.log(error)
    });
};

export const deleteBill = (id) => dispatch => {
  axios.delete(api_base_url + id, {datatype:'json'})
    .then((response) => dispatch({
      type: DELETE_BILL,
      payload: response.data
    }))
    .catch((error) => {
      console.log(error)
    })
};

export const updateBill = () => dispatch => {
  axios.put(api_base_url + id, {datatype:'json'})
    .then((response) => dispatch({
      type: UPDATE_BILL,
      payload: response.data
    }))
    .catch((error) => {
      console.log(error)
    })
};