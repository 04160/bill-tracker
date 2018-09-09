import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createBill } from '../actions/actions';

class BillForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      total: '',
      description: ''
    };

    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onChange(e) {
    this.setState({ [e.target.name]: e.target.value });
  }

  onSubmit(e) {
    e.preventDefault();
    const bill = {
      total: this.state.total,
      description: this.state.description
    };

    this.props.createBill(bill);
  }

  render() {
    return (
      <div>
        <h1>Add Bill</h1>
        <form onSubmit={this.onSubmit()}>
          <div>
            <label>Total: </label><br/>
            <input type="text" name="total" onChange={this.onChange} value={this.state.total}/>
          </div>
          <br/>
          <div>
            <label>Description: </label>
            <input type="text" name="description" onChange={this.onChange} value={this.state.description}/>
          </div>
          <br/>
          <button type="submit">Submit</button>
        </form>
      </div>
    )
  }
}

BillForm.propTypes = {
  createBill: PropTypes.func.isRequired
};

export default connect(null, { createBill })(BillForm);