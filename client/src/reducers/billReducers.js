import { GET_BILL, GET_BILLS, CREATE_BILL, DELETE_BILL, UPDATE_BILL } from '../actions/types';

const initialState = {
  items: [],
  item: {}
};

export default function (state = initialState, action) {
  switch (action.type) {
    case GET_BILL:
      console.log('GET_BILL');
      return {
        ...state,
        items: action.payload
      };
    case GET_BILLS:
      console.log('GET_BILLS');
      return {
        ...state,
        items: action.payload
      };
    case CREATE_BILL:
      console.log('CREATE_BILL');
      return {
        ...state,
        items: action.payload
      };
    case DELETE_BILL:
      console.log('DELETE_BILL');
      return {
        ...state,
        items: action.payload
      };
    case UPDATE_BILL:
      console.log('UPDATE_BILL');
      return {
        ...state,
        items: action.payload
      };
    default:
      console.log('default case');
      return state;
  }
}